package G

import (
	"github.com/kmou424/runabot/app/config"
	"github.com/pocketbase/pocketbase"
	"gopkg.in/telebot.v3"
)

var (
	AppName = "runabot"
	Debug   = false

	AppConfig *config.AppConfig

	Server *pocketbase.PocketBase
	Bot    *telebot.Bot
)
