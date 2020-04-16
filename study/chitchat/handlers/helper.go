package handlers

import (
	"errors"
	"fmt"
	"github.com/wuqinqiang/chitchat/models"
	"html/template"
	"net/http"
)

//通过Cookie 判断用户是否已登录
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("invalid session")
		}
	}
	return
}

//解析 HTML 模板 (应对需要传入多个模板文件的情况，避免重复编写模板)
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files [] string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

//生成响应 HTML
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	templates:=template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w,"layout",data)
}


//返回版本号
func Version() string  {
	return "0.1"
}