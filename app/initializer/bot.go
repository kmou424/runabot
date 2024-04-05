package initializer

import (
	"fmt"
	G "github.com/kmou424/runabot/app/global"
	"github.com/kmou424/runabot/bot"
	"gopkg.in/telebot.v3"
)

func initBot() {
	botConfig := G.AppConfig.Bot
	var err error
	G.Bot, err = telebot.NewBot(telebot.Settings{
		Token:   botConfig.Token,
		OnError: handleBotError,
		Updates: G.MaxUpdates,
		Poller:  G.HookPoller{},
	})
	if err != nil {
		G.ErrExit("can't create bot", err)
	}
	// hook updates from bot
	G.HookBotUpdates()

	// register handlers for bot
	bot.Uses(G.Bot)
	bot.Routes(G.Bot)

	if user := G.Bot.Me; user != nil {
		G.Logger.Info("bot create succeed",
			"id", user.ID,
			"name", fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			"username", user.Username,
		)
	}
}

func handleBotError(err error, _ telebot.Context) {
	G.OnError(err)
}
