package Protocol

const (
	TCP_TYPE = "tcp"
	UDP_TYPE = "udp"
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
