package middleware

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
	"just.for.test/rbacdemo/dao"
	"just.for.test/rbacdemo/rule"
)

// Authorizor 鉴权中间件
func Authorizor() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			name := r.PostFormValue("name")

			// TODO:实际使用中，从请求的token中解析出角色，然后进行判断
			user, err := dao.FindUserByName(name)
			if err != nil {
				return
			}
			role := user.Role
			if role == "" {
				role = "anonymous"
			}

			// casbin enforce
			logrus.Infoln(role, r.URL.Path, r.Method)
			res, err := rule.AuthEnforce.Enforce(role, r.URL.Path, r.Method)
			logrus.Infoln(res)
			if err != nil {
				writeError(http.StatusInternalServerError, "内部错误", w, err)
				return
			}
			if res {
				next.ServeHTTP(w, r)
			} else {
				writeError(http.StatusForbidden, "未授权页面", w, errors.New("未授权"))
				return
			}
		}

		return http.HandlerFunc(fn)
	}
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	logrus.Print("错误: ", err.Error())
	w.WriteHeader(status)
	_, _ = w.Write([]byte(message))
}
