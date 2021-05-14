package main

import (
	"github.com/thoj/go-ircevent"
	"strings"
)

func (i *IRCCat) inChannel(channel string) bool {
	return i.channels.Contains(channel)
}

func (i *IRCCat) authorisedUser(nick string) bool {
	_, exists := i.auth_users[nick]
	return exists
}

func (i *IRCCat) handleJoin(e *irc.Event) {
	if e.Arguments[0] == i.auth_channel {
		if i.metrics != nil {
			i.metrics.KnownUsers.Inc()
		}
		i.auth_users[e.Nick] = true
	}
}

func (i *IRCCat) handlePart(e *irc.Event) {
	if e.Arguments[0] == i.auth_channel {
		if i.metrics != nil {
			i.metrics.KnownUsers.Dec()
		}
		delete(i.auth_users, e.Nick)
	}
}

func (i *IRCCat) handleQuit(e *irc.Event) {
	if i.metrics != nil {
		i.metrics.KnownUsers.Dec()
	}
	delete(i.auth_users, e.Nick)
}

func (i *IRCCat) handleNames(e *irc.Event) {
	if e.Arguments[2] == i.auth_channel {
		nicks := strings.Split(e.Arguments[3], " ")
		for _, nick := range nicks {
			// TODO: this is probably not an optimal way of trimming the mode characters.
			nick = strings.TrimLeft(nick, "@%+")
			i.auth_users[nick] = true
		}
		if i.metrics != nil {
			i.metrics.KnownUsers.Add(float64(len(nicks)))
		}
	}
}

func (i *IRCCat) handleNick(e *irc.Event) {
	if i.auth_users[e.Nick] {
		delete(i.auth_users, e.Nick)
		i.auth_users[e.Arguments[0]] = true
	}
}
