package api

import (
	"net/http"

	"github.com/godwhoa/random-shit/botup.me/botup"
	"github.com/gorilla/sessions"
)

type Bot struct {
	Service botup.BotService
	Store   *sessions.CookieStore
}

func (b *Bot) AddBot(w http.ResponseWriter, r *http.Request) {
	bot, err := validate_addbot_form(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		return
	}
	_, err = b.Service.AddBot(bot)
	switch err {
	case botup.BotAlreadyExists:
		w.Write(botup.ERR_BOT_ALREADY_EXISTS)
	case nil:
		w.Write(botup.OK_BOT_ADDED)
	default:
		w.Write(botup.ERR_INTERNAL)
	}
}

func (b *Bot) AddPlugin(w http.ResponseWriter, r *http.Request) {
	plugin, err := validate_addplugin_form(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		return
	}

	err = b.Service.AddPlugin(plugin)
	switch err {
	case botup.PluginAlreadyExists:
		w.Write(botup.ERR_PLUGIN_ALREADY_EXISTS)
	case nil:
		w.Write(botup.OK_PLUGIN_ADDED)
	default:
		w.Write(botup.ERR_INTERNAL)
	}
}

func (b *Bot) RemoveBot(w http.ResponseWriter, r *http.Request) {
	bot, err := validate_removebot_form(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		return
	}

	err = b.Service.RemoveBot(bot)
	switch err {
	case botup.BotDoesntExists:
		w.Write(botup.ERR_BOT_DOESNT_EXISTS)
	case nil:
		w.Write(botup.OK_BOT_REMOVED)
	default:
		w.Write(botup.ERR_INTERNAL)
	}
}

func (b *Bot) RemovePlugin(w http.ResponseWriter, r *http.Request) {
	plugin, err := validate_removeplugin_form(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		return
	}

	err = b.Service.RemovePlugin(plugin)
	switch err {
	case botup.PluginDoesntExists:
		w.Write(botup.ERR_PLUGIN_DOESNT_EXISTS)
	case nil:
		w.Write(botup.OK_PLUGIN_REMOVED)
	default:
		w.Write(botup.ERR_INTERNAL)
	}
}
