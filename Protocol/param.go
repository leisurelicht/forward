package Protocol

const (
	TCPType = "tcp"
	UDPType = "udp"
)

type Param struct {
	Protocol    string
	ListenIP    *string
	ListenPort  *int
	ForwardIP   *string
	ForwardPort *int
}

func NewParam() *Param {
	return &Param{}
}
