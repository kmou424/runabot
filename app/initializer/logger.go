package initializer

import (
	"github.com/charmbracelet/log"
	G "github.com/kmou424/runabot/app/global"
)

func initLogger() {
	level := func() log.Level {
		if G.Debug {
			return log.DebugLevel
		}
		return log.InfoLevel
	}()

	G.Logger.SetLevel(level)
	G.Logger.SetReportCaller(G.Debug)

	if G.Debug {
		G.Logger.Info("app is running under debug mode!")
	}
}
