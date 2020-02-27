package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/go-session/session"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"

	"auth/auth"
	"auth/utils"
)

var (
	srv *server.Server // oauth2.v3/server
)

func main() {
	initAuthServer()

	initHttpServer()
}

func initAuthServer() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	// TODO: TokenStorage切换为REDIS存储
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// TODO: ClientStore切换为REDIS存储
	clientStore := store.NewClientStore()
	clientStore.Set("SomeAppId", &models.Client{
		ID:     "SomeAppId",
		Secret: "SomeSecret",
		Domain: "http://docs.qq.com",
	})
	manager.MapClientStorage(clientStore)

	srv = server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetAllowedGrantType(oauth2.AuthorizationCode, oauth2.Refreshing)
	srv.SetAllowedResponseType(oauth2.Code, oauth2.Token)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	srv.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		store, err := session.Start(nil, w, r)
		if err != nil {
			return
		}

		// 用户还没登录则跳转到登录页
		code, ok := store.Get("code")
		if !ok || code == "" {
			if r.Form == nil {
				r.ParseForm()
			}
			store.Set("form", r.Form)
			store.Save()

			w.Header().Set("Location", "/login")
			w.WriteHeader(http.StatusFound)
			return
		}

		// 用户登录成功则跳转回第三方应用
		openid, err := auth.Authorization(code.(string))
		if err != nil {
			return
		}

		userID = openid

		store.Delete("code")
		store.Save()
		return
	})

	srv.SetClientInfoHandler(func(r *http.Request) (clientID, clientSecret string, err error) {
		clientID = r.FormValue("client_id")
		clientSecret = r.FormValue("client_secret")
		return clientID, clientSecret, nil
	})
}

func initHttpServer() {
	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		store, err := session.Start(nil, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var form url.Values
		if v, ok := store.Get("form"); ok {
			form = v.(url.Values)
		}
		r.Form = form

		store.Delete("form")
		store.Save()

		err = srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		utils.OutputHTML(w, r, "static/login.html")
	})

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		store, err := session.Start(nil, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		code := r.FormValue("code")
		store.Set("code", code)
		store.Save()

		utils.OutputHTML(w, r, "static/auth.html")
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Server is running.")
	log.Fatal(http.ListenAndServe(":80", nil))
}
