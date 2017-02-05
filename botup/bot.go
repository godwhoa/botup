package botup

import "errors"

type Bot struct {
	BID     int    `json:"bid"`
	UID     string `json:"-"`
	Nick    string `json:"nick"`
	Addr    string `json:"server"`
	Channel string `json:"channel"`
	Alive   bool   `json:"alive"`
}

type Plugin struct {
	BID    int    `json:"bid"`
	UID    string `json:"-"`
	Plugin string `json:"plugin"`
}

var BotAlreadyExists = errors.New("Bot for server already exists")
var PluginAlreadyExists = errors.New("Plugin for bot already exists")
var BotDoesntExists = errors.New("Bot for server doesn't exists")
var PluginDoesntExists = errors.New("Plugin for bot doesn't exists")

type BotService interface {
	AddBot(bot Bot) (int, error)
	AddPlugin(plugin Plugin) error
	GetBot(UID string, BID int) (Bot, error)
	GetBots(UID string) ([]Bot, error)
	GetAllPlugins(UID string) ([]Plugin, error)
	GetPlugins(UID string, BID int) ([]string, error)
	RemoveBot(UID string, BID int) error
	RemovePlugin(plugin Plugin) error
}
