package qairc

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

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
	//Messages come in two flavours:
	//Either	":<nick>!<user>@<host> <type(arg0)> <arg1> <arg2>[ :<payload>]"
	//or		"<type(arg0) :<payload>"
	//distinguishable by the colon in front of "long" messages.
	islong := s[0] == ':'

	//The payload part is optional, so we need to check for the first
	//occurence of " :"
	p := strings.Index(s, " :")
	if p > -1 {
		//Lets drop the leading colon if it's a long message
		//Either way we handle payload as an additional element of Message.Args
		if islong {
			args = append(strings.Split(s[1:p], " "), s[p+2:])
		} else {
			args = append(strings.Split(s[0:p], " "), s[p+2:])
		}
	} else {
		args = strings.Split(s[1:], " ")
	}

	//If it's a long message we need to parse the message source too
	if islong {
		//CTCP: PRIVMSG ((target):^A ((command>) ((value)^A
		fmt.Println("args[3] is", len(args))
		if args[1] == "PRIVMSG" && len(args) > 2 && strings.HasPrefix(args[3], "\x01") {
			args[1] = "CTCP"
		}
		return Message{args[1], args[2:], s, ParseIdentity(args[0])}
	}
	return Message{args[0], args, s, Identity{"", "", "", ""}}
}

func newEngine() (c *Engine) {
	return &Engine{make(chan Message),
		make(chan string),
		make(chan struct{}),
		nil,
		nil,
		false,
		1 * time.Minute,
		"",
		Misc{"qairc", "qairc", "qairc", "qairc", "Qairc Golang Package"},
	}
}

func (c *Engine) SendRawf(format string, i ...interface{}) {
	c.In <- fmt.Sprintf(format, i...) + "\r\n"
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
			c.Out <- Parse(s)
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
	initMap()
	if err := c.checkSanity(); err != nil {
		return err
	}
	if c.UseTLS == true {
		dialer := &net.Dialer{Timeout: c.Timeout}
		c.Socket, err = tls.DialWithDialer(dialer, "tcp", c.Address, c.TLSCfg)
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

func (c *Engine) checkSanity() error {
	if len(c.Misc.Nick) == 0 {
		return errors.New("missing nick")
	}
	if len(c.Misc.RealName) == 0 {
		return errors.New("missing realname")
	}
	if len(c.Address) == 0 {
		return errors.New("missing serveraddress")
	}
	if strings.Count(c.Address, ":") != 1 || strings.Index(c.Address, ":") == len(c.Address)-1 {
		return errors.New("missing port")
	}
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
