package session

type ServerMethod struct {
	Version byte
	Auth    byte
	Method  byte
}

type ClientMethod struct {
	Version byte
	Method  byte
}
