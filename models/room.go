package models

// Representa uma Sala
type Room struct {
	Name string 	`json:"name"`		// nome da sala
	MaxPeers int 	`json:"maxPeers"`	// número máximo de Peers que podem se conectar
	Peers Peers		`json:"peers"`		// peers conectado a sala
	Owner string 	`json:"owner"`
	Addr  string 	`json:"addr"`
}

// Cria uma nova sala dado seu nome
func NewRoom(name, owner, addr string) *Room {
	return &Room{
		Name: name,
		MaxPeers: 5,
		Owner: owner,
		Peers: make(Peers),
		Addr: addr,
	}
}

// Retorna todos os peers conectados a esta sala
func (r *Room) GetPeers() Peers {
	return r.Peers
}

// Altera todos os peers da sala
func (r *Room) SetPeers(peers Peers) {
	r.Peers = peers
}

// Adiciona um novo peer a sala
func (r *Room) Add(peer *Peer)  {
	r.Peers[peer.ID] = peer
}

// Remove um Peer, passando seu ID
func (r *Room) Remove(peerID string)  {
	delete(r.Peers, peerID)
}