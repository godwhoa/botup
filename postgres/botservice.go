package postgres

import (
	"database/sql"
	"github.com/godwhoa/random-shit/botup.me/botup"
	_ "github.com/lib/pq"
)

type BotService struct {
	DB *sql.DB
}

var addbot_stmt = "INSERT INTO BOTS (UID,ALIVE,NICK,SERVER,CHANNEL) VALUES($1,$2,$3,$4,$5) RETURNING BID"

func (b BotService) AddBot(bot botup.Bot) (int, error) {
	var bid int
	err := b.DB.QueryRow(addbot_stmt, bot.UID, bot.Alive, bot.Nick, bot.Addr, bot.Channel).Scan(&bid)
	if err != nil {
		return bid, botup.BotAlreadyExists
	}
	return bid, nil
}

var addplugin_stmt = "INSERT INTO PLUGINS (BID,UID,PLUGIN) VALUES($1,$2,$3)"

func (b BotService) AddPlugin(plugin botup.Plugin) error {
	stmt, err := b.DB.Prepare(addplugin_stmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(plugin.BID, plugin.UID, plugin.Plugin)
	if err != nil {
		return botup.PluginAlreadyExists
	}
	return nil
}

var get_bots_stmt = "SELECT BID,NICK,ALIVE,SERVER,CHANNEL FROM BOTS WHERE UID = $1"

func (b BotService) GetBots(UID string) ([]botup.Bot, error) {
	bots := []botup.Bot{}
	rows, err := b.DB.Query(get_bots_stmt, UID)
	if err == sql.ErrNoRows {
		return bots, botup.BotDoesntExists
	}

	for rows.Next() {
		bot := botup.Bot{}
		err = rows.Scan(&bot.BID, &bot.Nick, &bot.Alive, &bot.Addr, &bot.Channel)
		if err != nil {
			return bots, nil
		}
		bots = append(bots, bot)
	}
	return bots, nil
}

var get_bot_stmt = "SELECT BID,NICK,ALIVE,SERVER,CHANNEL FROM BOTS WHERE UID = $1 AND BID = $2"

func (b BotService) GetBot(UID string, BID int) (botup.Bot, error) {
	bot := botup.Bot{}
	err := b.DB.QueryRow(get_bot_stmt, UID, BID).Scan(&bot.BID, &bot.Nick, &bot.Alive, &bot.Addr, &bot.Channel)
	if err == sql.ErrNoRows {
		return botup.Bot{}, botup.BotDoesntExists
	}
	return bot, nil
}

var get_allplugins_stmt = "SELECT BID,PLUGIN FROM PLUGINS WHERE UID = $1"

func (b BotService) GetAllPlugins(UID string) ([]botup.Plugin, error) {
	plugins := []botup.Plugin{}
	rows, err := b.DB.Query(get_allplugins_stmt, UID)
	if err == sql.ErrNoRows {
		return plugins, botup.PluginDoesntExists
	}

	for rows.Next() {
		plugin := botup.Plugin{}
		err = rows.Scan(&plugin.BID, &plugin.Plugin)
		if err != nil {
			return plugins, nil
		}
		plugins = append(plugins, plugin)
	}
	return plugins, nil
}

var get_plugins_stmt = "SELECT PLUGIN FROM PLUGINS WHERE UID = $1 AND BID = $2"

func (b BotService) GetPlugins(UID string, BID int) ([]string, error) {
	plugins := []string{}
	rows, err := b.DB.Query(get_plugins_stmt, UID, BID)
	if err != nil {
		return plugins, botup.PluginDoesntExists
	}

	for rows.Next() {
		var plugin string
		err = rows.Scan(&plugin)
		if err != nil {
			return plugins, nil
		}
		plugins = append(plugins, plugin)
	}
	return plugins, nil
}

var removebot_stmt = "DELETE FROM BOTS WHERE UID = $1 AND BID = $2"

func (b BotService) RemoveBot(UID string, BID int) error {
	stmt, err := b.DB.Prepare(removebot_stmt)
	if err != nil {
		return err
	}
	ret, err := stmt.Exec(UID, BID)
	affected, _ := ret.RowsAffected()
	if err != nil || affected < 1 {
		return botup.BotDoesntExists
	}
	return nil
}

var removeplugin_stmt = "DELETE FROM PLUGINS WHERE BID = $1 AND UID = $2 AND PLUGIN = $3"

func (b BotService) RemovePlugin(plugin botup.Plugin) error {
	stmt, err := b.DB.Prepare(removeplugin_stmt)
	if err != nil {
		return err
	}
	ret, err := stmt.Exec(plugin.BID, plugin.UID, plugin.Plugin)
	affected, _ := ret.RowsAffected()
	if err != nil || affected < 1 {
		return botup.PluginDoesntExists
	}
	return nil
}
