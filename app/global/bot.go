package G

var isExit bool

func StartBot() {
	go func() {
		for {
			Logger.Debug("starting bot")
			Bot.Start()
			// stopped
			OnStop()
		}
	}()
}

func StopBot(exit bool) {
	Logger.Debug("stopping bot")
	isExit = exit
	Bot.Stop()
}

func OnError(err error) {
	Logger.Error("an BotError occurred", "err", err)
	switch AppConfig.Bot.OnError {
	case "stop":
		StopBot(false)
	}
}

func OnStop() {
	if isExit {
		Logger.Debug("exiting bot")
		close(updatesHooker)
		Exit()
	}
	switch AppConfig.Bot.OnStop {
	case "restart":
		Logger.Error("bot stopped, restarting")
		return
	case "exit":
		Exit()
	}
}
