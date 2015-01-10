package qairc

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"strings"
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
	Out     chan string   //From server
	In      chan string   //To server
	control chan struct{} //to be extended
	Socket  net.Conn
	TLSCfg  *tls.Config
	Timeout time.Duration
	Address string
	Misc    Misc
}

func ParseIdentity(idstr string) Identity {
	if strings.Contains(idstr, "!") && strings.Contains(idstr, "@") {
		nick := strings.Split(idstr, "!")[0]
		remainder := strings.Split(idstr, "!")[1]
		user := strings.Split(remainder, "@")[0]
		addr := strings.Split(remainder, "@")[1]
		return Identity{nick, user, addr, idstr}
	}
	return Identity{"", "", idstr, idstr}
}

func Parse(s string) Message {
	var args []string
	islong := s[0] == ':'

	p := strings.Index(s, " :")
	if p > -1 {
		if islong {
			args = append(strings.Split(s[1:p], " "), s[p+2:])
		} else {
			args = append(strings.Split(s[0:p], " "), s[p+2:])
		}
	} else {
		args = strings.Split(s[1:], " ")
	}

	if islong {
		return Message{args[1], args[2:], s, ParseIdentity(args[0])}
	}
	return Message{args[0], args, s, Identity{"", "", "", ""}}
}

func newEngine() (c *Engine) {
	return &Engine{make(chan string),
		make(chan string),
		make(chan struct{}),
		nil,
		nil,
		1 * time.Minute,
		"",
		Misc{"qairc", "qairc", "qairc", "Qairc Golang Package"},
	}
}

func (c *Engine) readloop() {
	br := bufio.NewReaderSize(c.Socket, 1024)
	for true {
		select {
		case <-c.control:
			close(c.Out)
			return
		default:
			s, err := br.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			if len(s) > 4 && s[0:4] == "PING" {
				c.In <- "PONG" + s[4:]
			}

			c.Out <- s
		}
	}
}

func (c *Engine) writeloop() {
	bw := bufio.NewWriterSize(c.Socket, 1024)
	for true {
		select {
		case msgin := <-c.In:
			bw.WriteString(msgin)
			bw.Flush()
		case <-c.control:
			close(c.In)
			return
		}
	}
}

func (c *Engine) Stop() {
	close(c.control)
}

func (c *Engine) Run() error {
	var err error

	if c.TLSCfg != nil {
		if c.Socket, err = net.DialTimeout("tcp", c.Address, c.Timeout); err == nil {
			c.Socket = tls.Client(c.Socket, c.TLSCfg)
		} else {
			return err
		}
	} else {
		c.Socket, err = net.DialTimeout("tcp", c.Address, c.Timeout)
	}

	if err != nil {
		return err
	}

	go c.readloop()
	go c.writeloop()

	c.In <- "NICK " + c.Misc.Nick + "\r\n"
	c.In <- "USER " + c.Misc.RealName + " 0.0.0.0 0.0.0.0 :" + c.Misc.Description + "\r\n"

	return nil
}

func QAIrc(nick, user string) *Engine {
	if len(nick) == 0 || len(user) == 0 {
		return nil
	}
	q := newEngine()
	q.Misc = Misc{Nick: nick, AltNick: nick + "_", RealName: nick}

	return q
}
