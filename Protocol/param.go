package protocol

const (
	// TCPType TCP数据类型名
	TCPType = "tcp"
	// UDPType UDP数据类型名
	UDPType = "udp"
)

// Param 参数结构体
type Param struct {
	Protocol    string
	ListenIP    *string
	ListenPort  *int
	ForwardIP   *string
	ForwardPort *int
}

// NewParam 创建参数结构体
func NewParam() *Param {
	return &Param{}
}
