package server

import (
	"boggle/models"
	"net"
	"log"
	"github.com/pkg/errors"
	"encoding/json"
)

type TCPServer struct {
	Peers  models.Peers `json:"peers"`
	Rooms  models.Rooms `json:"rooms"`
	Socket net.Listener
	NewConnections chan net.Conn
	DeadConnections chan net.Conn
	RoomAddr net.IP
}

func NewTCPServer(addr string) (*TCPServer, error) {
	c, err := net.Listen("tcp", addr)

	if err != nil {
		return nil, err
	}

	ip := net.IP{224, 0, 0, 0}
	ip = ip.To4()

	if ip == nil {
		log.Fatal("non ipv4 address")
		return nil, err
	}

	ip = ip.Mask(ip.DefaultMask())

	return &TCPServer{
		Peers: make(models.Peers),
		Rooms: make(models.Rooms),
		Socket: c,
		NewConnections: make(chan net.Conn),
		DeadConnections: make(chan net.Conn),
		RoomAddr: ip,
	}, nil
}

func (this *TCPServer) Accept() {
	for {
		client, err := this.Socket.Accept()

		if err != nil {
			log.Fatalln(err)
		}

		go this.connect(client)
	}
}

func (this *TCPServer) CreateRoom(name, owner string) error {
	_, exists := this.Rooms[name]

	if exists {
		return errors.New(ERR_EXISTS_ROOM)
	}

	room := models.NewRoom(name, owner, this.nextRoomAddr())
	this.Rooms[room.Name] = room
	this.Peers[owner].Initial = true

	return nil
}

func (this *TCPServer) AddPeer(peer *models.Peer) {
	this.Peers[peer.ID] = peer
}

// func (s *TCPServer) AddPeerToRoom(peerID string, roomID String)

func (this *TCPServer) HandleRequest(msgReceive models.Message, conn net.Conn) *models.Message {

	response := models.NewMessage()

	payload := msgReceive.Payload

	switch msgReceive.Type {
	case "CONNECT":

		err := checkFieldsInPayload([]string{"username"}, payload)

		if err != nil {
			response.Type = "ERROR"
			response.Payload["error"] = err.Error()
			return response
		}

		peer := models.NewPeer(payload["username"].(string))
		peer.SetSocket(conn)
		this.AddPeer(peer)

		rooms := []string{}

		for _, r := range this.Rooms {
			rooms = append(rooms, r.Name)
		}

		response.Type = "CONNECTED"

		response.Payload["id"] = peer.ID
		response.Payload["username"] = peer.Username
		response.Payload["rooms"] = rooms
	case "DISCONNECT":
		strongParams := []string{"id", "roomName"}

		err := checkFieldsInPayload(strongParams, payload)

		if err != nil {
			response.Type = "ERROR"
			response.Payload["error"] = err.Error()
			return response
		}

		roomID := payload["roomName"].(string)

		room := this.Rooms[roomID]
		peerID := payload["id"].(string)

		log.Println(room.Owner)
		log.Println()

		// Se o Peer é dono de uma sala e a sala é a que ele criou
		if this.Peers[peerID].Initial && room.Owner == peerID {
			// Remove todos os peers dessa sala, do servidor
			// 	e remove a sala
			for peerID, peer := range room.Peers {
				delete(this.Peers, peerID)

				m := models.NewMessage()
				m.Type = "DISCONNECT"

				json.NewEncoder(peer.Socket).Encode(m)

				peer.Socket.Close()
			}

			delete(this.Rooms, roomID)
		} else {
			response.Type = "ERROR"
			response.Payload["error"] = ERR_UNAUTHORIZED
			return response
		}

	case "CREATE_ROOM":
		strongParams := []string{"id", "roomName"}

		err := checkFieldsInPayload(strongParams, payload)

		if err != nil {
			response.Type = "ERROR"
			response.Payload["error"] = err.Error()
			return response
		}

		roomID := payload["roomName"].(string)
		peerID := payload["id"].(string)

		err = this.CreateRoom(roomID, peerID)

		if err != nil {
			response.Type = "ERROR"
			response.Payload["error"] = err.Error()
			return response
		}

		response.Type = "CREATED_ROOM"
		response.Payload["initial"] = true

		for _, peer := range this.Peers {
			if peer.ID != peerID {
				message := models.NewMessage()
				message.Type = "NEW_ROOM"
				message.Payload["room"] = roomID

				json.NewEncoder(peer.Socket).Encode(message)
			}

		}

	case "JOIN_ROOM":
		strongParams := []string{"id", "roomName"}

		err := checkFieldsInPayload(strongParams, payload)

		if err != nil {
			response.Type = "ERROR"
			response.Payload["error"] = err.Error()
			return response
		}

		roomID := payload["roomName"].(string)
		peerID := payload["id"].(string)

		// Se a sala existe
		if _, ok := this.Rooms[roomID]; ok {
			// recupera sala
			room := this.Rooms[roomID]

			// Checa se a sala está cheia
			// Se não estiver, adiciona o cliente na sala
			if len(room.Peers) < room.MaxPeers {
				room.Add(this.Peers[peerID])
				response.Type = "JOINED_ROOM"
				response.Payload["roomAddr"] = room.Addr
				response.Payload["peers"] = room.Peers
				response.Payload["initial"] = room.Owner == peerID
				for _, peer := range room.Peers {
					if peer.ID != peerID {
						message := models.NewMessage()
						message.Type = "NEW_PEER"
						message.Payload["peer"] = room.Peers[peerID]

						json.NewEncoder(peer.Socket).Encode(message)
					}

				}
			} else {
				// Retorna um erro
				response.Type = "ERROR"
				response.Payload["error"] = ERR_FULL_ROOM
				return response
			}

		} else {
			response.Type = "ERROR"
			response.Payload["error"] = ERR_NOT_FOUND_ROOM
			return response
		}
	}

	return response
}

func (this *TCPServer) connect(client net.Conn) {
	this.NewConnections <- client
}

func checkFieldsInPayload(fields[] string, payload map[string]interface{}) error {
	for _, field := range fields {
		_, ok := payload[field]
		if !ok {
			return errors.New(field + "not found")
		}

	}

	return nil
}

func (this *TCPServer) nextRoomAddr() string {
	this.RoomAddr[3]++
	return this.RoomAddr.String()
}