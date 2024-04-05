package main

import (
	G "github.com/kmou424/runabot/app/global"
	"github.com/kmou424/runabot/app/initializer"
)

func init() {
	initializer.Apply()
}

func main() {
	err := G.Server.Start()
	if err != nil {
		G.ErrExit("server terminated", err)
	}
}
