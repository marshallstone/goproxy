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

type Reply struct {
	Version byte
	Rep     byte
	RSV     byte
	Atyp    byte
	BndAddr []byte
	BndPort []byte
}
