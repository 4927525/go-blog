package views

import (
	"errors"
	"go-blog/common"
	"go-blog/context"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	category := common.Template.Category

	//页面上涉及到的所有的数据，必须有定义

	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		category.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}

	path := r.URL.Path
	cidStr := strings.TrimPrefix(path, "/c/")
	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		category.WriteError(w, errors.New("路径不匹配"))
	}

	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	hr, err := service.GetPostsByCategoryID(cid, page, pageSize)

	if err != nil {
		log.Println("Index 获取数据出错", err)
		category.WriteError(w, errors.New("系统错误，请联系管理员"))
	}

	category.WriteData(w, hr)
}

func (*HTMLApi) CategoryNew(ctx *context.MsContext)  {
	categoryTemplate := common.Template.Category
	//http://localhost:8080/c/1  1参数 分类的id
	cIdStr := ctx.GetPathVariable("id")
	cId,_ := strconv.Atoi(cIdStr)
	pageStr,_ := ctx.GetForm("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page,_ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryResponse,err := service.GetPostsByCategoryID(cId,page,pageSize);
	if err != nil {
		categoryTemplate.WriteError(ctx.W,err)
		return
	}
	categoryTemplate.WriteData(ctx.W,categoryResponse)
}