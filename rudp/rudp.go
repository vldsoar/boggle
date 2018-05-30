package rudp

import (
	"boggle/models"
	"net"
	"log"
	"encoding/json"
	"time"
)

// Essa estrutura representa um Protocolo de confiabilidade entrega
// do protocolo UDP. Abaixo segue descrição dos campos:

// 	Recipients: os destinatários esperados
// 	Attempts: tabela de tentativas
// 	PendingMessages: tabela de mensagens pendentes
// 	ServerAddr: endereço
// 	MaxDatagramSize: o tamanho máximo do buffer
// 	NewMessage: canal de novas mensagens
// 	IDMsgSent: id atual para mensagens
// 	Listen: socket da conexão UDP Multicast
// 	ReceivedMsgs: mensagens recebidas

type RUDP struct {
	ID              string 					 `json:"id"`
	Recipients      map[string]bool          `json:"recipients"`
	Attempts        map[string]int           `json:"attempts"`
	PendingMessages map[uint]map[string]bool `json:"pendingMessages"`
	ServerAddr      string                   `json:"serverAddr"`
	MaxDatagramSize int                      `json:"maxDatagramSize"`
	NewMessage      chan models.Message
	IDMsgSent       *counter
	Listen          *net.UDPConn
	Conn            *net.UDPConn
	ReceivedMsgs    map[string]map[uint]bool
}

// const MaxDatagramSize

const (
	N_ATTEMPTS = 4
	DATAGRAM_SIZE = 2048
)

// Cria uma nova 'instancia' do Protocolo
func create(clients []string, id, serverAddr string, maxDatagramSize int) RUDP {

	recipients := make(map[string]bool)

	for _, value := range clients {
		recipients[value] = true
	}

	count := counter(1)

	return RUDP{
		ID:              id,
		Recipients:      recipients,
		Attempts:        make(map[string]int),
		PendingMessages: make(map[uint]map[string]bool),
		ReceivedMsgs:    make(map[string]map[uint]bool),
		ServerAddr:      serverAddr,
		MaxDatagramSize: maxDatagramSize,
		IDMsgSent:       &count,
		NewMessage:      make(chan models.Message),
	}
}

// Função para criar conexão UDP
func ListenMulticastRUDP(recipients []string, id, serverAddr string) (RUDP, error) {
	this := create(recipients, id, serverAddr, DATAGRAM_SIZE)

	addr, err := net.ResolveUDPAddr("udp", this.ServerAddr)

	if err != nil {
		log.Fatalln(err)
		return this, err
	}
	// Cria uma conexão
	this.Listen, err = net.ListenMulticastUDP("udp", nil, addr)

	if err != nil {
		log.Fatalln(err)
		return this, err
	}

	this.Conn, err = net.DialUDP("udp", nil, addr)

	if err != nil {
		log.Fatalln(err)
		return this, err
	}

	//defer this.Listen.Close()

	// Limita o tamanho do buffer
	this.Listen.SetReadBuffer(this.MaxDatagramSize)

	return this, nil
}

// Função para envio de mensagem para endereço destino
func (this *RUDP) Send(msg *models.Message) {

	log.Println("# SEND MESSAGE #")
	log.Println("RECIPIENTES: ", this.Recipients)

	// Se a mensagem do remetente é de confirmação de recebimento,
	// 	não espera confirmação para esse tipo de mensagem
	notExpectConfirmation := msg.From == this.ID && msg.Type == "RECEIVED"

	if notExpectConfirmation {
		// Atribui um ID
		msg.SetID(uint(this.IDMsgSent.GetAndIncrement()))
		// Codifica mensagem e envia
		this.encodeMessageAndSend(msg)

		log.Println("Sending confirmation message -> ", msg.ID)

		return
	}


	// Checa se é uma mensagem existente, se sim
	//  altera o timeout, caso contrário, atribui um ID a nova mensagem
	if clients, ok := this.PendingMessages[msg.ID]; ok {
		log.Println("Resending message -> ", msg.ID)
		msg.Timeout *= 2

		// Se não tiver destinatários
		//	deleta mensagem e sai da função
		if len(clients) == 0 {
			delete(this.PendingMessages, msg.ID)
			return
		}
	} else {
		// add timeout
		msg.Timeout = 3

		msg.SetID(uint(this.IDMsgSent.GetAndIncrement()))

		log.Println("Sending message -> ", msg.ID)

		// Adiciona na tabela de pendentes
		this.addPendingMessage(msg)
	}

	this.encodeMessageAndSend(msg)

	// 	Inicializa Timeout
	go this.initializeTimeout(msg)

}

// Função para receber uma nova mensagem
func (this *RUDP) Receive() models.Message {
	select {
	case msg := <-this.NewMessage:
		log.Println("Receive Message -> ", msg.ID)
		return msg
	}
}

// Função para ler mensagens do UDP Multicast
func (this *RUDP) Read() {
	// Inicializa o buffer
	b := make([]byte, this.MaxDatagramSize)

	var msg models.Message

	// ler mensagem do UDP
	n, _, err := this.Listen.ReadFromUDP(b)

	if err != nil {
		log.Println(err)
		return
	}

	// Decodifica mensagem
	err = json.Unmarshal(b[:n], &msg)

	if err != nil {
		log.Println(err)
		return
	}

	// Se foi se o destinatário é o mesmo que o remetente
	// sai da função
	if msg.From == this.ID {
		return
	} else {
		log.Println("Read -> ", msg)
	}

	// Trata mensagem em uma goroutine
	go this.handleMessage(msg)

}

// Fechar conexão
func (this *RUDP) Close() {
	this.Listen.Close()
}

func (this *RUDP) SetRecipients(clients []string) {
	recipients := make(map[string]bool)

	for _, value := range clients {
		recipients[value] = true
	}

	this.Recipients = recipients
}

// Função para adicionar uma mensagem pendente
func (this *RUDP) addPendingMessage(msg *models.Message) {
	rec := make(map[string]bool)

	for k, v := range this.Recipients {
		rec[k] = v
	}

	this.PendingMessages[msg.ID] = rec
	log.Println("Add Pending Message: ", this.PendingMessages[msg.ID])
}

// Função para remover um Cliente de uma mensagem pendente
func (this *RUDP) removeFromPendingMessages(msgID uint, clientID string) {
	log.Println("Remove from Pending Messages")
	log.Println("Before deleting from the Pending Messages: ", len(this.PendingMessages[msgID]))
	delete(this.PendingMessages[msgID], clientID)
	log.Println("After deleting from the Pending Messages: ", len(this.PendingMessages[msgID]))
}

// Função para incrementar o número de tentativas de um Cliente
func (this *RUDP) incrementAttempts(id string) int {
	this.Attempts[id] += 1
	return this.Attempts[id]
}

// Retorna o número de mensagens pendentes de um determinado cliente
func (this *RUDP) pendingMessagesFor(ClientID string) int {
	count := 0

	for _, v := range this.PendingMessages {
		if _, ok := v[ClientID]; ok {
			count++
		}
	}

	return count
}

// Função para tratar as mensagens recebidas
func (this *RUDP) handleMessage(msg models.Message)  {
	// Se mensagem vazia descarta
	if msg.IsEmpty() {
		return
	}

	// Se a mensagem ja foi recebida, descarta
	if _, ok := this.ReceivedMsgs[msg.From][msg.ID]; ok {
		log.Println("Discard Message: ", msg)
		return
	}

	// Adiciona mensagem nas mensagens recebidas
	this.ReceivedMsgs[msg.From] = make(map[uint]bool)
	this.ReceivedMsgs[msg.From][msg.ID] = true

	log.Println("Add in Received Messages: ", this.ReceivedMsgs)

	// Se a mensagem é uma confirmação de recebimento
	// remove o cliente que confirmou da tabela de mensagens pendentes
	if msg.Type == "RECEIVED" {
		msgID := uint(msg.Payload["msgID"].(float64))
		this.removeFromPendingMessages(msgID, msg.From)

	} else {
		// Cria uma mensagem de Confirmação de Recebimento
		m := models.NewMessage().SetFrom(this.ID).SetType("RECEIVED")
		m.Payload["msgID"] = msg.ID
		go this.Send(m) // Envia mensagem

		this.NewMessage <- msg // Envia para o canal

	}
}

func (this *RUDP) initializeTimeout(msg *models.Message) {
	select {
	// Caso atinja o timeout
	case <-time.After(time.Duration(msg.Timeout) * time.Second):
		//mutex := sync.Mutex{}
		log.Println("Message timeout expired: ", msg.ID)

		// Se existe destinatários sem receber a mensagem
		if len(this.PendingMessages[msg.ID]) > 0 {
			log.Println("Pending Recipients: ", len(this.PendingMessages[msg.ID]))
			//mutex.Lock()
			// Itera sobre os destinatários da mensagem
			for clientID, _ := range this.PendingMessages[msg.ID] {
				_, ok := this.Attempts[clientID]

				// Se destinatário não está na tabela de tentativas
				// adiciona à tabela
				if !ok {
					this.Attempts[clientID] = 2
				} else {
					// Incrementa contador de tentativas do destinatários
					att := this.incrementAttempts(clientID)

					// Se atingiu o número máximo de tentativas
					// 	remove o destinatário da lista de destinatários esperados
					//	e de todas as mensagens pendentes

					if (att / this.pendingMessagesFor(clientID)) >= N_ATTEMPTS {
						log.Println("Client disconnected", clientID)
						delete(this.Recipients, clientID)
						for k, _ := range this.PendingMessages {
							this.removeFromPendingMessages(k, clientID)
						}
						msg := models.NewMessage()
						msg.Type = "CLIENT_DISCONNECTED"
						msg.Payload["clientID"] = clientID
						this.NewMessage <- *msg
					}
				}
			}
			//mutex.Unlock()
			// Reenvia a mensagem
			go this.Send(msg)
		} else {
			// Deleta mensagem da tabela de mensagens pendentes
			delete(this.PendingMessages, msg.ID)
			log.Println("Message removed from Pending Table", msg)
		}

	}

}

func (this *RUDP) encodeMessageAndSend(msg *models.Message) {
	// Codifica mensagem
	encodeMsg, err := json.Marshal(msg)

	if err != nil {
		log.Println(err)
		return
	}

	// Envia mensagem
	this.Conn.Write(encodeMsg)
}