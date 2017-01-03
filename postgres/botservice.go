package postgres

import (
	"database/sql"
	"github.com/godwhoa/random-shit/botup.me/botup"
	_ "github.com/lib/pq"
)

type BotService struct {
	DB *sql.DB
}

var addbot_stmt = "INSERT INTO BOTS (UID,ALIVE,SERVER,CHANNEL) VALUES($1,$2,$3,$4) RETURNING BID"

func (b BotService) AddBot(bot botup.Bot) (int, error) {
	var bid int
	err := b.DB.QueryRow(addbot_stmt, bot.UID, bot.Alive, bot.Addr, bot.Channel).Scan(&bid)
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

var removebot_stmt = "DELETE FROM BOTS WHERE UID = $1 AND BID = $2"

func (b BotService) RemoveBot(bot botup.Bot) error {
	stmt, err := b.DB.Prepare(removebot_stmt)
	if err != nil {
		return err
	}
	ret, err := stmt.Exec(bot.UID, bot.BID)
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
