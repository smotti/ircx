Note
====

Add some functions and a Commands map attribute to the Bot. This way one can
attach specific bot commands in the form of regexp and can add callbacks for
them to the bot.

Example:

```
package main

import (
	"flag"
	"log"

	"github.com/nickvanw/ircx"
	"github.com/sorcix/irc"
)

var (
	name     = flag.String("name", "ircx", "Nick to use in IRC")
	server   = flag.String("server", "chat.freenode.org:6667", "Host:Port to connect to")
	channels = flag.String("chan", "#test", "Channels to join")
  cmds     = map[string]string{
    "CMD_HELLO": "^!hello$",
  }
)

func init() {
	flag.Parse()
}

func main() {
	bot := ircx.Classic(*server, *name)
  bot.Commands = cmds
	if err := bot.Connect(); err != nil {
		log.Panicln("Unable to dial IRC Server ", err)
	}
	RegisterHandlers(bot)
	bot.CallbackLoop()
	log.Println("Exiting..")
}

func RegisterHandlers(bot *ircx.Bot) {
	bot.AddCallback(irc.RPL_WELCOME, ircx.Callback{Handler: ircx.HandlerFunc(RegisterConnect)})
	bot.AddCallback(irc.PING, ircx.Callback{Handler: ircx.HandlerFunc(PingHandler)})
  bot.AddCallback("CMD_HELLO", ircx.Callback{Handler: ircx.HandlerFunc(Hello)})
}

func RegisterConnect(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command: irc.JOIN,
		Params:  []string{*channels},
	})
}

func PingHandler(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command:  irc.PONG,
		Params:   m.Params,
		Trailing: m.Trailing,
	})
}

func Hello(s ircx.Sender, m *irc.Message) {
  s.Send(&irc.Message{
    Command: irc.PRIVMSG,
    Params: []string{m.Prefix.Name},
    Trailing: "hello world",
  })
}

```


ircx
====

ircx is a very basic IRC bot written on top of the wonderfully small [sorcix/irc](https://github.com/sorcix/irc) library. It's designed to be a small building block, a small example of one way to use the library.

Using it is very simple:

```
package main

import (
	"flag"
	"log"

	"github.com/nickvanw/ircx"
	"github.com/sorcix/irc"
)

var (
	name     = flag.String("name", "ircx", "Nick to use in IRC")
	server   = flag.String("server", "chat.freenode.org:6667", "Host:Port to connect to")
	channels = flag.String("chan", "#test", "Channels to join")
)

func init() {
	flag.Parse()
}

func main() {
	bot := ircx.Classic(*server, *name)
	if err := bot.Connect(); err != nil {
		log.Panicln("Unable to dial IRC Server ", err)
	}
	RegisterHandlers(bot)
	bot.CallbackLoop()
	log.Println("Exiting..")
}

func RegisterHandlers(bot *ircx.Bot) {
	bot.AddCallback(irc.RPL_WELCOME, ircx.Callback{Handler: ircx.HandlerFunc(RegisterConnect)})
	bot.AddCallback(irc.PING, ircx.Callback{Handler: ircx.HandlerFunc(PingHandler)})
}

func RegisterConnect(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command: irc.JOIN,
		Params:  []string{*channels},
	})
}

func PingHandler(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command:  irc.PONG,
		Params:   m.Params,
		Trailing: m.Trailing,
	})
}
```


This example doesn't do anything other than connect to specified channels and idle, but it's trivial to add additional handlers for any IRC event you want.

Context can be passed around by creating custom Handlers and Senders and using them, versus the default sender created, and an empty handler struct.
