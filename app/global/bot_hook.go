package G

import (
	"gopkg.in/telebot.v3"
)

const MaxUpdates = 100

var updatesHooker = make(chan telebot.Update, MaxUpdates)

type HookPoller struct {
	telebot.LongPoller
}

func (poller HookPoller) Poll(b *telebot.Bot, _ chan telebot.Update, stop chan struct{}) {
	poller.LongPoller.Poll(b, updatesHooker, stop)
}

func HookBotUpdates() {
	go func() {
		Logger.Debug("updates hooker started")
		for {
			select {
			case update, ok := <-updatesHooker:
				if !ok {
					Logger.Info("updates hooker closed")
					return
				}
				// context := G.Bot.NewContext(update)
				Bot.Updates <- update
			}
		}
	}()
}
