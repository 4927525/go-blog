package api

import (
	"errors"
	"go-blog/common"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/service"
	"go-blog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cid := params["categoryId"].(string)
		cids, _ := strconv.Atoi(cid)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		title := params["title"].(string)
		slug := params["slug"].(string)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			cids,
			uid,
			0,
			0,
			time.Now(),
			time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)

	case http.MethodPut:
		params := common.GetRequestJsonParam(r)
		cid := params["categoryId"].(string)
		cids, _ := strconv.Atoi(cid)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		title := params["title"].(string)
		slug := params["slug"].(string)
		postType := params["type"].(float64)
		pidFloat := params["pid"].(float64)
		pid := int(pidFloat)
		pType := int(postType)
		post := &models.Post{
			pid,
			title,
			slug,
			content,
			markdown,
			cids,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}
	//params := common.GetRequestJsonParam(r)
}

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	post, err := dao.GetPostByID(pid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)
	common.Success(w, searchResp)
}
