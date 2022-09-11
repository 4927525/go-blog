package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type TomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}

type SystemConfig struct {
	Ip             string
	Port           string
	AppName        string
	Version        float32
	CurrentDir     string
	CdnURL         string
	QiniuAccessKey string
	QiniuSecretKey string
	Valine         bool
	ValineAppid    string
	ValineAppkey   string
	ValineServerURL string
}

var Cfg *TomlConfig

func init() {
	Cfg = new(TomlConfig)
	Cfg.System.AppName = "go-blog"
	Cfg.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir

	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}

}
