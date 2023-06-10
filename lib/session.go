package session

type InitPacket struct {
	Version    byte
	Auth       byte
	AuthMethod byte
}
