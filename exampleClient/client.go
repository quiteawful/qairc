package main

import (
	"crypto/tls"
	"log"

	"qairc"
)

func main() {
	client := qairc.QAIrc("test1", "test2")
	client.Address = "irc.quiteawful.net:6697"
	client.UseTLS = true
	client.TLSCfg = &tls.Config{InsecureSkipVerify: true}
	err := client.Run()
	log.Println(qairc.Numerics["005"])
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("go!")
	for {
		m, status := <-client.Out
		if !status {
			log.Println("Out closed, exiting")
			client.Reconnect()
		}

		if m.Type == "001" {
			client.Join("#g0")
		}
		if m.Type == "PRIVMSG" {
			l := len(m.Args)
			log.Println(m.Sender)
			log.Println(m.Timestamp)
			log.Println(m.Args[0 : l-1])
			log.Println(m.Args[l-1])
		}
	}
}
