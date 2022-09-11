package main

import (
	"go-blog/common"
	"go-blog/config"
	"go-blog/server"
)

func init() {
	common.LoadTemplate()
}

func main() {
	server.App.Start(config.Cfg.System.Ip, config.Cfg.System.Port)
}
