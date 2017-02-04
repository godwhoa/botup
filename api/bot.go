package api

import (
	"net/http"

	"github.com/godwhoa/random-shit/botup.me/api/validate"
	"github.com/godwhoa/random-shit/botup.me/botup"
	"github.com/gorilla/sessions"
	"github.com/uber-go/zap"
)

type Bot struct {
	Service botup.BotService
	Store   *sessions.CookieStore
	Log     zap.Logger
}

func (b *Bot) AddBot(w http.ResponseWriter, r *http.Request) {
	bot, err := validate.AddBot(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		b.Log.Error("api.bot.addbot",
			zap.Error(err),
			zap.String("info", "addbot form fields missing"),
			zap.String("ip", r.RemoteAddr),
		)
		return
	}
	_, err = b.Service.AddBot(bot)
	switch err {
	case botup.BotAlreadyExists:
		w.Write(botup.ERR_BOT_ALREADY_EXISTS)
		b.Log.Error("api.bot.addbot",
			zap.Error(err),
			zap.String("info", "bot already exists"),
			zap.String("ip", r.RemoteAddr),
		)
	case nil:
		w.Write(botup.OK_BOT_ADDED)
		b.Log.Info("api.bot.addbot",
			zap.String("info", "bot added"),
			zap.String("uid", bot.UID),
			zap.String("uid", bot.Nick),
			zap.String("addr", bot.Addr),
			zap.String("channel", bot.Channel),
			zap.String("ip", r.RemoteAddr),
		)
	default:
		w.Write(botup.ERR_INTERNAL)
		b.Log.Error("api.user.addbot",
			zap.Error(err),
			zap.String("info", "error adding bot"),
			zap.String("ip", r.RemoteAddr),
		)
	}
}

func (b *Bot) AddPlugin(w http.ResponseWriter, r *http.Request) {
	plugin, err := validate.AddPlugin(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		b.Log.Error("api.bot.addplugin",
			zap.Error(err),
			zap.String("info", "addplugin form fields missing"),
			zap.String("ip", r.RemoteAddr),
		)
		return
	}

	err = b.Service.AddPlugin(plugin)
	switch err {
	case botup.PluginAlreadyExists:
		w.Write(botup.ERR_PLUGIN_ALREADY_EXISTS)
		b.Log.Error("api.bot.addplugin",
			zap.Error(err),
			zap.String("info", "plugin already exists"),
			zap.String("ip", r.RemoteAddr),
		)
	case nil:
		w.Write(botup.OK_PLUGIN_ADDED)
		b.Log.Info("api.bot.addplugin",
			zap.String("info", "plugin added"),
			zap.String("uid", plugin.UID),
			zap.Int("bid", plugin.BID),
			zap.String("plugin", plugin.Plugin),
			zap.String("ip", r.RemoteAddr),
		)
	default:
		w.Write(botup.ERR_INTERNAL)
		b.Log.Error("api.user.addplugin",
			zap.Error(err),
			zap.String("info", "error adding plugin"),
			zap.String("ip", r.RemoteAddr),
		)
	}
}

func (b *Bot) RemoveBot(w http.ResponseWriter, r *http.Request) {
	bot, err := validate.RemoveBot(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		b.Log.Error("api.bot.removebot",
			zap.Error(err),
			zap.String("info", "removebot form fields missing"),
			zap.String("ip", r.RemoteAddr),
		)
		return
	}

	err = b.Service.RemoveBot(bot.UID, bot.BID)
	switch err {
	case botup.BotDoesntExists:
		w.Write(botup.ERR_BOT_DOESNT_EXISTS)
		b.Log.Error("api.bot.removebot",
			zap.Error(err),
			zap.String("info", "bot doesn't exists"),
			zap.String("ip", r.RemoteAddr),
		)
	case nil:
		w.Write(botup.OK_BOT_REMOVED)
		b.Log.Info("api.bot.removebot",
			zap.String("info", "bot removed"),
			zap.String("uid", bot.UID),
			zap.Int("bid", bot.BID),
			zap.String("ip", r.RemoteAddr),
		)
	default:
		w.Write(botup.ERR_INTERNAL)
		b.Log.Error("api.bot.removebot",
			zap.Error(err),
			zap.String("info", "error removing bot"),
			zap.String("ip", r.RemoteAddr),
		)
	}
}

// TODO
func (b *Bot) GetBot(w http.ResponseWriter, r *http.Request)    {}
func (b *Bot) GetPlugin(w http.ResponseWriter, r *http.Request) {}

func (b *Bot) RemovePlugin(w http.ResponseWriter, r *http.Request) {
	plugin, err := validate.RemovePlugin(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		b.Log.Error("api.bot.removeplugin",
			zap.Error(err),
			zap.String("info", "removeplugin form fields missing"),
			zap.String("ip", r.RemoteAddr),
		)
		return
	}

	err = b.Service.RemovePlugin(plugin)
	switch err {
	case botup.PluginDoesntExists:
		w.Write(botup.ERR_PLUGIN_DOESNT_EXISTS)
		b.Log.Error("api.bot.removeplugin",
			zap.Error(err),
			zap.String("info", "plugin doesn't exists"),
			zap.String("ip", r.RemoteAddr),
		)
	case nil:
		w.Write(botup.OK_PLUGIN_REMOVED)
		b.Log.Info("api.bot.removeplugin",
			zap.String("info", "plugin removed"),
			zap.String("uid", plugin.UID),
			zap.Int("bid", plugin.BID),
			zap.String("ip", r.RemoteAddr),
		)
	default:
		w.Write(botup.ERR_INTERNAL)
		b.Log.Error("api.bot.removeplugin",
			zap.Error(err),
			zap.String("info", "error removing plugin"),
			zap.String("ip", r.RemoteAddr),
		)
	}
}
