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
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("go!")
	for {
		m := qairc.Parse(<-client.Out)
		log.Println(m.Raw)
	}
}
