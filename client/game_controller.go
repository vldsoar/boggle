package main

import (
	"boggle/models"
	"github.com/therecipe/qt/core"
	"boggle/rudp"
	"net"
	"encoding/json"
	"log"
	"github.com/fatih/structs"
	"fmt"
	"os"
	"math/rand"
	"time"
)

// Essa estrutura é a responsável pelo controler do jogo
// fazendo conexão entre a interface gráfica e o Golang
type GameController struct {
	// signal envia para QML
	// slot envia para GO
	core.QObject
	Peers 	map[string]string
	Rooms 	[]string
	Info	map[string]interface{}

	GameWords map[string][]string

	connRUDP 	*rudp.RUDP
	connTCP  	net.Conn

	resServer chan models.Message
	runRUDP   chan bool

	_ func(data string) 		`slot:"login"`
	_ func(data string) 		`slot:"joinRoom"`
	_ func(data string) bool 	`slot:"sendWord"`
	_ func(data string) 		`slot:"startGame"`
	_ func(data string) 		`slot:"createRoom"`
	_ func()					`slot:"finishGame"`

	_ func(reply string)  `signal:"appendUser"`
	_ func(reply string)  `signal:"appendRoom"`
	_ func(reply string)  `signal:"deleteUser"`
	_ func(reply string)  `signal:"updateRooms"`
	_ func(reply string)  `signal:"sessionAuthenticated"`
	_ func() 			  `signal:"sessionAuthenticationError"`
	_ func(reply string)  `signal:"createRoomError"`
	_ func(reply string)  `signal:"joinedRoom"`
	_ func(reply string)  `signal:"joinRoomError"`
	_ func() 			  `signal:"preparingGame"`
	_ func(reply string)  `signal:"readyGame"`
	_ func(reply string)  `signal:"finishedGame"`
}

// Função para inicializar a estrutura
func (this *GameController) init(addr string) {
	c, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalln(err)
	}

	this.connTCP = c

	this.Info 		= make(map[string]interface{})
	this.Peers 		= make(map[string]string)
	this.Rooms 		= []string{}
	this.resServer 	= make(chan models.Message)
	this.runRUDP 	= make(chan bool)

	this.GameWords 	= make(map[string][]string)

	this.ConnectLogin(this.GCLogin)
	this.ConnectJoinRoom(this.GCJoinRoom)
	this.ConnectCreateRoom(this.GCCreateRoom)
	this.ConnectSendWord(this.GCSendWord)
	this.ConnectStartGame(this.GCStartGame)
	this.ConnectFinishGame(this.GCFinishGame)

	// trata requisições do servidor em uma go routine
	go handleMessagesTCP(this)

	// inicializa o listener multicast em uma go routine
	go this.StartRUDP()

}

// Função para efetuar login
func (this *GameController) GCLogin(data string)  {
	// Cria uma mensagem
	msg := models.NewMessage()
	msg.Type = "CONNECT"
	msg.Payload["username"] = data

	// Codifica e envia para o servidor
	json.NewEncoder(this.connTCP).Encode(msg)

	// espera uma resposta do servidor, através de um channel
	res := <- this.resServer

	// Se não houver erro
	if res.Type != "ERROR" {

		// Atribui o payload contendo as salas, no atributo
		//	do controlador
		r, _ := json.Marshal(res.Payload["rooms"])
		json.Unmarshal(r, &this.Rooms)

		// Atribui as informações de id e username
		// 	aos atributos do controlador
		this.Info["id"] = res.Payload["id"].(string)
		this.Info["username"] = res.Payload["username"].(string)

		//
		b, _ := json.Marshal(res.Payload)
		json.Unmarshal(b, this.Info)

		log.Println(this.Info)
		// Envia para interface as informações do usuário
		this.SessionAuthenticated(string(b))
		// Envia para interface as salas disponíveis
		this.UpdateRooms(string(r))
	} else {
		this.SessionAuthenticationError()
	}

}

func (this *GameController) GCJoinRoom(data string) {
	// Cria uma nova mensagem
	msg := models.NewMessage()
	msg.Type = "JOIN_ROOM"
	msg.Payload["id"] = this.Info["id"].(string)
	msg.Payload["roomName"] = data

	// decodifica e envia para servidor
	json.NewEncoder(this.connTCP).Encode(msg)
	// Espera resposta do servidor através de um channel
	res := <- this.resServer

	if res.Type != "ERROR" {
		var peers models.Peers

		p := res.Payload["peers"]

		binaryPeers, _ := json.Marshal(p)

		json.Unmarshal(binaryPeers, &peers)

		users := []string{}

		for _, p := range peers {
			if p.ID != this.Info["id"] {
				this.Peers[p.ID] = p.Username
			}
			users = append(users, p.Username)
		}

		this.Info["joinedRoom"] = struct {
			Name string 	`json:"name"`
			Addr string 	`json:"addr"`
			Initial bool 	`json:"initial"`
			Users []string  `json:"users"`
		}{
			data,
			res.Payload["roomAddr"].(string),
			res.Payload["initial"].(bool),
			users,
		}

		log.Println("**** SHOW INFO ****")
		log.Println(this.Info)

		b, _ := json.Marshal(this.Info)

		this.JoinedRoom(string(b))
	} else {
		this.JoinRoomError(res.Payload["error"].(string))
	}
}

func (this *GameController) GCCreateRoom(data string)  {
	msg := models.NewMessage()
	msg.Type = "CREATE_ROOM"
	msg.Payload["id"] = this.Info["id"].(string)
	msg.Payload["roomName"] = data

	json.NewEncoder(this.connTCP).Encode(msg)

	res := <- this.resServer

	if res.Type != "ERROR" {
		//this.Info["initial"] = res.Payload["initial"].(bool)

		log.Println(this.Info)
		this.Rooms = append(this.Rooms, data)
		this.GCJoinRoom(data)
	} else {
		this.CreateRoomError(res.Payload["error"].(string))
	}

}

func (this *GameController) GCSendWord(data string) bool {
	myUsername := this.Info["username"].(string)

	// Checa se ja tem mensagem, se sim retorna falso
	for _, word := range this.GameWords[myUsername] {
		if word == data {
			return false
		}
	}

	// Checa se existe no dicionário
	if _, ok := dic[data]; ok {
		// Atribui ao controlador
		this.GameWords[myUsername] = append(this.GameWords[myUsername], data)
		// Cria uma nova mensagem
		msg := models.NewMessage()
		msg.From = this.Info["id"].(string)
		msg.Type = "NEW_WORD"
		msg.Payload["word"] = data
		msg.Payload["username"] = myUsername

		log.Println("************")
		log.Println(this.connRUDP.Recipients)
		log.Println("************")

		// Envia para multicast
		go this.connRUDP.Send(msg)

		return true
	}

	return false

}

func (this *GameController) GCStartGame(data string)  {
	msg := models.NewMessage()
	msg.Type = "DISCONNECT"
	msg.Payload["id"] = this.Info["id"].(string)
	msg.Payload["roomName"] = data

	json.NewEncoder(this.connTCP).Encode(msg)

}

func (this *GameController) GCFinishGame() {
	myID := this.Info["id"].(string)
	myUsername := this.Info["username"].(string)

	var myWords []string

	// Coloca as palavras em um array
	for _, word := range this.GameWords[myUsername] {
		myWords = append(myWords, word)
	}

	// Cria uma nova mensagem
	msg := models.NewMessage()
	msg.From = myID
	msg.Type = "FINISH_GAME"
	msg.Payload["words"] = myWords
	msg.Payload["username"] = myUsername

	// Envia pra o multicast
	this.connRUDP.Send(msg)

	// Aguarda 10 segundos
	go func() {
		<-time.After(time.Second * 10)

		// Decodifica todas as mensagens
		b, _ := json.Marshal(this.GameWords)

		// Avisa interface que o jogo acabou
		this.FinishedGame(string(b))
	}()

}

func (this *GameController) StartRUDP() {
	select {
	case <-this.runRUDP:
		fmt.Println("Start RUDP...")
		// Muda o estado do jogo
		this.PreparingGame()

		var recipients []string

		// Adiciona todos os peers em um array
		for id, _ := range this.Peers {
			recipients = append(recipients, id)
		}

		addr := structs.Map(this.Info["joinedRoom"])["Addr"].(string)
		myID := this.Info["id"].(string)

		c, err := rudp.ListenMulticastRUDP(recipients, myID, addr + ":9999")

		if err != nil {
			os.Exit(1)
		}

		this.connRUDP = &c

		// Fecha conexão quando finalizar aplicação
		defer this.connRUDP.Close()

		// Goroutine para sempre escutar do multicast
		go func(gc *GameController) {
			for {
				gc.connRUDP.Read()
			}
		}(this)

		// Se for um "dono" de uma sala, envia comando para começar o jogo
		if structs.Map(this.Info["joinedRoom"])["Initial"].(bool) {
			this.PrepareGame()
		}

		// Loop para escutar e tratar mensagens que alteram o estado
		// 	do jogo
		for {
			m := this.connRUDP.Receive()

			switch m.Type {
			case "READY_GAME":
				// Informa a interface que o jogo começou
				this.Info["currentGame"] = m.Payload["currentGame"]
				b, _ := json.Marshal(this.Info["currentGame"])

				this.ReadyGame(string(b))

			case "NEW_WORD":
				// Caso uma nova palavra seja enviada por algum peer
				// armazena no controller
				username := m.Payload["username"].(string)
				arr := this.GameWords[username]

				this.GameWords[username] = append(arr, m.Payload["word"].(string))
			case "FINISH_GAME":
				// Quando encerra o jogo, recebe todas as palavras de cada peer
				var words []string
				username := m.Payload["username"].(string)

				r, _ := json.Marshal(m.Payload["words"])
				json.Unmarshal(r, &words)

				this.GameWords[username] = words
			case "CLIENT_DISCONNECTED":
				// Quando o Peer cai, atualiza informação no controller
				//	e interface
				peerID := m.Payload["clientID"].(string)
				username := this.Peers[peerID]

				delete(this.Peers, peerID)

				this.DeleteUser(username)
			}
		}


	}

}

// Função que prepara a partida
func (this *GameController) PrepareGame()  {
	// gera os dado
	board := generateDices()

	// atribui as informações ao controllador
	this.Info["currentGame"] = struct {
		Score int 			`json:"score"`
		Board [16]string 	`json:"board"`
	}{0, board}

	// Cria uma nova mensagem
	msg := models.NewMessage()
	msg.Type = "READY_GAME"
	msg.From = this.Info["id"].(string)
	msg.Payload["currentGame"] = this.Info["currentGame"]

	// Envia mensagem para todos os peers da sala
	this.connRUDP.Send(msg)

	// Codifica mensagem e envia informações para interface gráfica
	b, _ := json.Marshal(this.Info["currentGame"])
	this.ReadyGame(string(b))

}

// Função que trata as mensagens enviadas pelo Servidor
func handleMessagesTCP(gc *GameController) {
	// Loop
	for {
		var msg models.Message
		// decodifica mensagem do servidor e armazena na variável msg
		err := json.NewDecoder(gc.connTCP).Decode(&msg)

		// Se ocorrer algum erro sai do loop
		if err != nil {
			log.Println(err)
			break
		}

		log.Println(msg)
		// Diversas ações de acordo o Type da mensagem
		switch msg.Type {
		case "NEW_PEER": // Caso um novo peer se conecte a sala
			peer := msg.Payload["peer"].(map[string]interface{})
			// atualiza o controller
			gc.Peers[peer["id"].(string)] = peer["username"].(string)
			// envia para interface
			gc.AppendUser(peer["username"].(string))
		case "REMOVE_PEER": {
			// o mesmo ocorre que em NEW_PEER, exceto que aqui remove
			// 	um peer desconectado
			peerID := msg.Payload["peerID"].(string)
			username := gc.Peers[peerID]

			delete(gc.Peers, peerID)

			gc.DeleteUser(username)
		}
		case "NEW_ROOM":
			// Quando uma nova sala é criada, repete o processo acima
			room := msg.Payload["room"].(string)
			gc.Rooms = append(gc.Rooms, room)
			b, _ := json.Marshal(gc.Rooms)
			gc.UpdateRooms(string(b))
		case "DELETE_ROOM":
			// Quando uma sala é deletada, repete o processo de codificação e envio
			//	para interface
			b, _ := json.Marshal(msg.Payload["rooms"])

			json.Unmarshal(b, &gc.Rooms)

			gc.UpdateRooms(string(b))
		case "DISCONNECT":
			// Quando recebe do servidor a mensagem de desfazer conexão
			// 	ou seja, o jogo irá começar
			// 	avisa através de um channel para go routine que irá efetuar o
			// 	processo de inicialização da partida
			gc.runRUDP <- true

		default:
			// caso não seja nenhumas das opções anteriores
			// 	foi uma solicitação que espera resposta em mesmo bloco de código
			// 	nesse caso, envia mensagem através de um channel
			gc.resServer <- msg
		}
	}
}


// Função para gerar os Dados
func generateDices() [16]string {

	dices := []string{
		"AAEEGN", "ABBJOO", "ACHOPS", "AFFKPS",
		"AOOTTW", "CIMOTU", "DEILRX", "DELRVY",
		"DISTTY", "EEGHNW", "EEINSU", "EHRTVW",
		"EIOSST", "ELRTTY", "HIMNUQ", "HLNNRZ",
	}

	board := [16]string{}

	for i := 0; i < 16; i++  {
		rand.Seed(time.Now().UnixNano())
		c := string(dices[i][rand.Intn(6)])
		if c == "Q" {
			board[i] = "QU"
		} else {
			board[i] = c
		}
	}


	return board

}