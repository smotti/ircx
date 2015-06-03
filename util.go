package ircx

import (
    "regexp"
    "log"

    "github.com/sorcix/irc"
)

// connectMessages is a list of IRC messages to send when attempting to
// connect to the IRC server.
func (b *Bot) connectMessages() []*irc.Message {
	messages := []*irc.Message{}
	messages = append(messages, &irc.Message{
		Command:  irc.USER,
		Params:   []string{b.User, "0", "*"},
		Trailing: b.User,
	})
	messages = append(messages, &irc.Message{
		Command: irc.NICK,
		Params:  []string{b.OriginalName},
	})
	if b.Password != "" {
		messages = append(messages, &irc.Message{
			Command: irc.PASS,
			Params:  []string{b.Password},
		})
	}
	return messages
}

// isQuery checks if the PRIVMSG is a private query or from a channel.
// True if it's a private query.
func isQuery(m *irc.Message) bool {
    if m.Command == irc.PRIVMSG {
        for _, v := range m.Params {
            matched, err := regexp.MatchString("^#.*", v)
            if err != nil {
                log.Println("Error:", err)
                return false
            }
            if ! matched {
                return true
            }
        }
    }
    return false
}
