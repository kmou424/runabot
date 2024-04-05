package initializer

import (
	"errors"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
	"github.com/kmou424/runabot/app/config"
	G "github.com/kmou424/runabot/app/global"
	"github.com/pelletier/go-toml/v2"
	"strings"
)

var AppConfigPaths = []string{
	"./runabot.toml",
	sysutil.UserConfigDir("runabot/runabot.toml"),
}

func initAppConfig() {
	appConfig := new(config.AppConfig)
	defer func() {
		G.AppConfig = appConfig
		afterLoad()
	}()
	for _, path := range AppConfigPaths {
		if strings.Trim(path, " ") == "" {
			continue
		}
		if !fsutil.IsFile(path) {
			continue
		}

		if data, err := fsutil.ReadOrErr(path); err == nil {
			err := toml.Unmarshal(data, appConfig)
			if err != nil {
				G.Logger.Error("parse config file failed", "path", path, "err", err)
				continue
			}
			G.Logger.Info("config file loaded", "path", path)
			return
		} else {
			G.Logger.Error("read config file failed", "path", path, "err", err)
		}
	}

	G.ErrExit("failed to initialize app", errors.New("no available config file found"))
}

func afterLoad() {
	G.Debug = G.AppConfig.App.Debug
}
