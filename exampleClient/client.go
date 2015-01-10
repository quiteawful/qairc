package main

import (
	"log"

	"github.com/quiteawful/qairc"
)

func main() {
	client := qairc.QAIrc("test1", "test2")
	client.Address = "irc.quiteawful.net:6667"

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
