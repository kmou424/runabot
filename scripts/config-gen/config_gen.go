package main

import (
	"github.com/charmbracelet/log"
	"github.com/gookit/goutil/fsutil"
	"github.com/kmou424/runabot/app/config"
	"github.com/pelletier/go-toml/v2"
	"os"
	"time"
)

var defaultConfig = &config.AppConfig{
	App: config.AppTable{
		Debug: false,
	},
	Bot: config.BotTable{
		Token:   "",
		OnError: "stop",
		OnStop:  "restart",
	},
	Server: config.ServerTable{
		Host: "127.0.0.1",
		Port: 6060,
	},
}

var logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportCaller:    false,
	ReportTimestamp: false,
	TimeFormat:      time.DateTime,
	Prefix:          "config_generator",
	Level:           log.DebugLevel,
})

func main() {
	data, err := toml.Marshal(defaultConfig)
	if err != nil {
		logger.Error(err)
		return
	}
	err = fsutil.WriteFile("./runabot.default.toml", data, os.ModePerm)
	if err != nil {
		logger.Error(err)
		return
	}
}
