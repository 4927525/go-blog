package views

import (
	"go-blog/common"
	"go-blog/config"
	"go-blog/context"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter,r *http.Request)  {
	login := common.Template.Login
	login.WriteData(w,config.Cfg.Viewer)
}

func (*HTMLApi) LoginNew(ctx *context.MsContext)  {
	login := common.Template.Login

	login.WriteData(ctx.W,config.Cfg.Viewer)
}