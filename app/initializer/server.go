package initializer

import (
	"errors"
	"fmt"
	"github.com/kmou424/runabot/app/dao"
	G "github.com/kmou424/runabot/app/global"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"os"
)

func initServer() {
	G.Server = pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDev:      G.Debug,
		HideStartBanner: true,
	})
	handleBotLifecycle()
	handleDaoInit()
	handleServerError()
	hookCmdline()
}

func handleBotLifecycle() {
	G.Server.OnBeforeServe().PreAdd(func(_ *core.ServeEvent) error {
		G.StartBot()
		return nil
	})
	G.Server.OnTerminate().Add(func(_ *core.TerminateEvent) error {
		G.Cleanup()
		return nil
	})
}

func handleDaoInit() {
	G.Server.OnBeforeServe().Add(func(_ *core.ServeEvent) error {
		dao.Init()
		return nil
	})
}

func handleServerError() {
	G.Server.OnAfterApiError().Add(func(e *core.ApiErrorEvent) error {
		G.Logger.Error("an ApiError occurred", "err", e.Error)
		return nil
	})
}

func hookCmdline() {
	// hook serve
	if G.AppConfig.Server.Port < 1024 {
		G.ErrExit("cmdline error", errors.New("only non-privileged port is allowed"))
	}

	serverAddr := fmt.Sprintf("%s:%d", G.AppConfig.Server.Host, G.AppConfig.Server.Port)
	os.Args = []string{
		os.Args[0], "serve", fmt.Sprintf("--http=%s", serverAddr),
	}
	G.Server.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		G.Logger.Infof("REST API: http://%s/api/", serverAddr)
		G.Logger.Infof("Admin UI: http://%s/_/", serverAddr)
		return nil
	})
}
