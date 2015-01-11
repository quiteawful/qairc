package qairc

func (e *Engine) SendRaw(message string) {
	e.In <- message + "\r\n"
}

func (e *Engine) Quit() {
	e.SendRaw("QUIT")
}

func (e *Engine) Join(channel string) {
	e.SendRaw("JOIN " + channel)
}

func (e *Engine) Part(channel string) {
	e.SendRaw("PART " + channel)
}

func (e *Engine) Notice(target, message string) {
	e.SendRaw("NOTICE " + target + " :" + message)
}

func (e *Engine) Action(target, message string) {
	e.SendRaw("PRIVMSG " + target + " :\001ACTION " + message + "\001")
}

func (e *Engine) Privmsg(target, message string) {
	e.SendRaw("PRIVMSG " + target + " :" + message)
}

func (e *Engine) Whois(nick string) {
	e.SendRaw("WHOIS " + nick)
}

func (e *Engine) Who(nick string) {
	e.SendRaw("WHO " + nick)
}
