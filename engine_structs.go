package qairc

import (
	"crypto/tls"
	"net"
	"time"
)

type Identity struct {
	Nick string
	User string
	Host string
	Full string
}

type Message struct {
	Type   string
	Args   []string
	Raw    string
	Sender Identity
}

type Misc struct {
	Nick        string
	AltNick     string
	ActualNick  string
	RealName    string
	Description string
}

type Engine struct {
	Out     chan Message  //From server
	In      chan string   //To server
	control chan struct{} //to be extended
	Socket  net.Conn
	TLSCfg  *tls.Config
	UseTLS  bool
	Timeout time.Duration
	Address string
	Misc    Misc
}
