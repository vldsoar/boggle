package main

import (
	"os"
	"fmt"
	"boggle/server"
	"boggle/models"
	"net"
	"encoding/json"
	"flag"
	"boggle/util"
	"log"
)

func main() {
	// Instancia um servidor
	ip, _ := util.ExternalIP()

	portTCP := flag.String("portTCP", "3000", "8000")

	log.Println("IP: ", ip)

	s, err := server.NewTCPServer(fmt.Sprintf(":%s", *portTCP))

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	fmt.Println("Start Server...")

	// Uma nova go routine ouve novas conexões
	go s.Accept()

	// Close the listener when the application closes.
	defer s.Socket.Close()

	for {
		select {
		// Quando uma nova conexão chega, o 'socket' do cliente é
		// enviado para uma função que irá tratar suas requisições
		case conn := <-s.NewConnections:
			fmt.Println("Accept new connection" + conn.RemoteAddr().String())
			go handleClient(s, conn)
			// Quando o cliente desconecta do servidor
		case conn := <-s.DeadConnections:
			fmt.Println("Client disconnected " + conn.RemoteAddr().String())
			go handleDisconnectClient(s, conn)
			fmt.Println("Count: ", len(s.Peers))
		}

	}

}

func handleClient(server *server.TCPServer, conn net.Conn)  {
	// Loop enquanto o client estiver conectado
	for {
		var msg models.Message
		// decodifica mensagem em json enviada pelo cliente
		// e salva os dados na variavel msg
		err := json.NewDecoder(conn).Decode(&msg)

		if err != nil {
			break
		}
		// Envia a mensagem para o servidor tratar
		// 	que retorna uma resposta
		res := server.HandleRequest(msg, conn)

		fmt.Println(res)
		// Envia resposta para o client
		json.NewEncoder(conn).Encode(res)
	}
	// Notifica que o cliente desconectou
	server.DeadConnections <- conn
}

// FALTA ENVIAR PARA TODOS Q ELE DESCONECTOU

func handleDisconnectClient(server *server.TCPServer, conn net.Conn) {
	var findPeer *models.Peer
	// Verifica se o Peer está na lista do servidor
	for _, peer := range server.Peers {
		if peer.Socket == conn {
			findPeer = peer
		}
	}
	// Se não estiver retorna
	if findPeer == nil {
		return
	}

	var findRoom *models.Room

	Loop:
		for _, room := range server.Rooms {
			for _, p := range room.Peers {
				if p.ID == findPeer.ID {
					findRoom = room
					break Loop
				}
			}
		}

	// Deleta o Peer
	delete(server.Peers, findPeer.ID)

	// Se não encontrar Sala
	if findRoom == nil {
		return
	}

	// delete o peer da sala
	findRoom.Remove(findPeer.ID)

	// Checa se a sala tem outros peers
	if len(findRoom.Peers) > 0 {

		msg := models.NewMessage()
		msg.Type = "REMOVE_PEER"
		msg.Payload["peerID"] = findPeer.ID

		for _, peer := range findRoom.Peers {
			json.NewEncoder(peer.Socket).Encode(msg)
		}

		return

	} else { // caso contrário, deleta a sala
		msg := models.NewMessage()
		msg.Type = "REMOVE_ROOM"
		msg.Payload["roomName"] = findRoom.Name

		for _, peer := range findRoom.Peers {
			json.NewEncoder(peer.Socket).Encode(msg)
		}

		delete(server.Rooms, findRoom.Name)
	}


}
