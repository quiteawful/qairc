package qairc

import "testing"

func TestQAIrc(t *testing.T) {
	q := QAIrc("", "")
	if q != nil {
		t.Fatal("Should got an <nil> error")
	}

	nick := "testnick"
	altnick := nick + "_"
	realname := "testreal"
	q = QAIrc(nick, realname)
	if q == nil {
		t.Fatal("Return value should not be nil.")
	}
	if q.Misc.Nick != nick {
		t.Fatal("Nick should match, doesnt.")
	}
	if q.Misc.AltNick != altnick {
		t.Fatal("AltNick should match, doesnt.")
	}
	if q.Misc.RealName != realname {
		t.Fatal("Realname should match, doesnt.")
	}
}

func TestRun(t *testing.T) {

}
