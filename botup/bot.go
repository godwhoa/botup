package botup

import "errors"

type Bot struct {
	BID     int
	UID     string
	Nick    string
	Addr    string
	Channel string
	Alive   bool
}

type Plugin struct {
	BID    int
	UID    string
	Plugin string
}

var BotAlreadyExists = errors.New("Bot for server already exists")
var PluginAlreadyExists = errors.New("Plugin for bot already exists")
var BotDoesntExists = errors.New("Bot for server doesn't exists")
var PluginDoesntExists = errors.New("Plugin for bot doesn't exists")

type BotService interface {
	AddBot(bot Bot) (int, error)
	AddPlugin(plugin Plugin) error
	RemoveBot(bot Bot) error
	RemovePlugin(plugin Plugin) error
}
