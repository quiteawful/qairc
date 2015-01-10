package main

import (
	"log"

	"github.com/quiteawful/qairc"
)

func main() {
	irc := qairc.New()
	irc.Address = "irc.quiteawful.net:6667"

	err := irc.Run()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("go!")
	for {
		m := qairc.Parse(<-irc.Out)
		log.Println(m.Raw)
	}
}
