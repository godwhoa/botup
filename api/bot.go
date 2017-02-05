package api

import (
	"encoding/json"
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

func (b *Bot) GetBot(w http.ResponseWriter, r *http.Request) {
	uid, bid, err := validate.GetBot(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		b.Log.Error("api.bot.getbot",
			zap.Error(err),
			zap.String("info", "getbot form fields missing"),
			zap.String("ip", r.RemoteAddr),
		)
		return
	}

	bot, err := b.Service.GetBot(uid, bid)
	switch err {
	case botup.BotDoesntExists:
		w.Write(botup.ERR_BOT_DOESNT_EXISTS)
		b.Log.Error("api.bot.getbot",
			zap.Error(err),
			zap.String("info", "bot doesn't exists"),
			zap.String("ip", r.RemoteAddr),
		)
	case nil:
		encoded, err := json.Marshal(bot)
		if err != nil {
			w.Write(botup.ERR_INTERNAL)
			b.Log.Error("api.bot.getbot",
				zap.Error(err),
				zap.String("info", "error encoding bot to json"),
				zap.String("ip", r.RemoteAddr),
			)
			return
		}
		w.Write(encoded)
		b.Log.Info("api.bot.getbot",
			zap.String("info", "bot fetched"),
			zap.String("uid", uid),
			zap.Int("bid", bid),
			zap.String("ip", r.RemoteAddr),
		)
	default:
		w.Write(botup.ERR_INTERNAL)
		b.Log.Error("api.bot.getbot",
			zap.Error(err),
			zap.String("info", "error getting bot"),
			zap.String("ip", r.RemoteAddr),
		)
	}
}

func (b *Bot) GetBots(w http.ResponseWriter, r *http.Request) {
	uid, err := validate.GetBots(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		b.Log.Error("api.bot.getbots",
			zap.Error(err),
			zap.String("info", "getbot form fields missing"),
			zap.String("ip", r.RemoteAddr),
		)
		return
	}

	bots, err := b.Service.GetBots(uid)
	switch err {
	case botup.BotDoesntExists:
		w.Write(botup.ERR_BOT_DOESNT_EXISTS)
		b.Log.Error("api.bot.getbots",
			zap.Error(err),
			zap.String("info", "bot doesn't exists"),
			zap.String("ip", r.RemoteAddr),
		)
	case nil:
		encoded, err := json.Marshal(bots)
		if err != nil {
			w.Write(botup.ERR_INTERNAL)
			b.Log.Error("api.bot.getbot",
				zap.Error(err),
				zap.String("info", "error encoding bot to json"),
				zap.String("ip", r.RemoteAddr),
			)
			return
		}
		w.Write(encoded)
		b.Log.Info("api.bot.getbots",
			zap.String("info", "bots fetched"),
			zap.String("uid", uid),
			zap.String("ip", r.RemoteAddr),
		)
	default:
		w.Write(botup.ERR_INTERNAL)
		b.Log.Error("api.bot.getbots",
			zap.Error(err),
			zap.String("info", "error getting bots"),
			zap.String("ip", r.RemoteAddr),
		)
	}
}

// TODO: this and GetPlugins
func (b *Bot) GetAllPlugins(w http.ResponseWriter, r *http.Request) {
	uid, err := validate.GetAllPlugins(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		b.Log.Error("api.bot.getallplugins",
			zap.Error(err),
			zap.String("info", "getallplugins form fields missing"),
			zap.String("ip", r.RemoteAddr),
		)
		return
	}

	plugins, err := b.Service.GetAllPlugins(uid)
	switch err {
	case botup.PluginDoesntExists:
		w.Write(botup.ERR_PLUGIN_DOESNT_EXISTS)
		b.Log.Error("api.bot.getallplugins",
			zap.Error(err),
			zap.String("info", "no plugins exists"),
			zap.String("ip", r.RemoteAddr),
		)
	case nil:
		encoded, err := json.Marshal(plugins)
		if err != nil {
			w.Write(botup.ERR_INTERNAL)
			b.Log.Error("api.bot.getallplugins",
				zap.Error(err),
				zap.String("info", "error encoding all plugins to json"),
				zap.String("ip", r.RemoteAddr),
			)
			return
		}
		w.Write(encoded)
		b.Log.Info("api.bot.getbots",
			zap.String("info", "all plugins fetched"),
			zap.String("uid", uid),
			zap.String("ip", r.RemoteAddr),
		)
	default:
		w.Write(botup.ERR_INTERNAL)
		b.Log.Error("api.bot.getallplugins",
			zap.Error(err),
			zap.String("info", "error getting all plugins"),
			zap.String("ip", r.RemoteAddr),
		)
	}
}

func (b *Bot) GetPlugins(w http.ResponseWriter, r *http.Request) {
	uid, bid, err := validate.GetPlugins(r, b.Store)
	if err != nil {
		w.Write(botup.ERR_FIELDS_MISSING)
		b.Log.Error("api.bot.getplugins",
			zap.Error(err),
			zap.String("info", "getplugins form fields missing"),
			zap.String("ip", r.RemoteAddr),
		)
		return
	}

	plugins, err := b.Service.GetPlugins(uid, bid)
	switch err {
	case botup.PluginDoesntExists:
		w.Write(botup.ERR_PLUGIN_DOESNT_EXISTS)
		b.Log.Error("api.bot.getplugins",
			zap.Error(err),
			zap.String("info", "no plugins exists"),
			zap.String("ip", r.RemoteAddr),
		)
	case nil:
		encoded, err := json.Marshal(plugins)
		if err != nil {
			w.Write(botup.ERR_INTERNAL)
			b.Log.Error("api.bot.getplugins",
				zap.Error(err),
				zap.String("info", "error encoding all plugins to json"),
				zap.String("ip", r.RemoteAddr),
			)
			return
		}
		w.Write(encoded)
		b.Log.Info("api.bot.getbots",
			zap.String("info", "plugins fetched"),
			zap.String("uid", uid),
			zap.Int("bid", bid),
			zap.String("ip", r.RemoteAddr),
		)
	default:
		w.Write(botup.ERR_INTERNAL)
		b.Log.Error("api.bot.getplugins",
			zap.Error(err),
			zap.String("info", "error getting plugins"),
			zap.String("ip", r.RemoteAddr),
		)
	}
}

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
