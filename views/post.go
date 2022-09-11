package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	path := r.URL.Path
	pidStr := strings.TrimPrefix(path, "/p/")
	pidStr = strings.TrimSuffix(pidStr, ".html")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		detail.WriteError(w, errors.New("路径不匹配"))
	}
	postRes, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteData(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)
}
