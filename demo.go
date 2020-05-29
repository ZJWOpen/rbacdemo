package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"just.for.test/rbacdemo/common"
	"just.for.test/rbacdemo/dao"
	"just.for.test/rbacdemo/middleware"
	"just.for.test/rbacdemo/model"
	"just.for.test/rbacdemo/rule"
)

// UserRespone 首页的请求处理
type UserRespone struct{}

func (*UserRespone) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome...")
}

func main() {
	model.InitDb()
	rule.Init()
	mux := http.NewServeMux()
	mux.Handle("/", &UserRespone{})
	mux.HandleFunc("/login", loginHandler())
	mux.HandleFunc("/logout", logoutHandler())
	mux.HandleFunc("/member/current", currentMemberHandler())
	http.ListenAndServe(":8080", (middleware.Authorizor())(mux))
}

type loginResp struct {
	Code    int64      `json:"code"`
	Message string     `json:"message"`
	Data    *tokenResp `json:"data"`
}

type tokenResp struct {
	Token string `json:"token"`
}

func loginHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")

		user, err := dao.FindUserByName(name)
		if err != nil {
			return
		}
		// 创建 token 值
		token := common.Token{
			ID:       user.ID,
			UserName: user.UserName,
			Role:     user.Role,
		}
		tokenStr := common.TokenToString(&token)
		resp := new(loginResp)
		tokenRsp := new(tokenResp)
		tokenRsp.Token = tokenStr
		resp.Code = http.StatusOK
		resp.Data = tokenRsp
		w.WriteHeader(http.StatusOK)
		respbyt, err := json.Marshal(resp)
		if err != nil {
			return
		}

		_, _ = w.Write(respbyt)
	})
}

func logoutHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")

		user, err := dao.FindUserByName(name)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(user.UserName + " logout"))
	})
}

func currentMemberHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")

		user, err := dao.FindUserByName(name)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(user.UserName + " Hello world"))
	})
}
