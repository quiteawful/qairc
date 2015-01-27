package main

import (
	"crypto/tls"
	"log"

	"github.com/quiteawful/qairc"
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
		m := <-client.Out
		log.Println(m.Raw)

		if m.Type == "001" {
			client.Join("#g0")
		}

		if m.IsCTCP() {
			log.Println("CTCP received")
		}

		if m.IsPrivMsg() {
			log.Println(m.GetPrivmsg())
		}
	}
}
