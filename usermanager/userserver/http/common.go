package http

import (
	"fmt"
	"shopee/usermanager/userserver/g"
	"shopee/usermanager/userserver/rpc"
	"net/http"
	"strings"
)

func configCommonRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%s\n", g.VersionMsg())))
	})

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, g.Config())
	})

	http.HandleFunc("/config/reload", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.RemoteAddr, "127.0.0.1") {
			g.ParseConfig(g.ConfigFile)
			RenderDataJson(w, "ok")
		} else {
			RenderDataJson(w, "no privilege")
		}
	})

	http.HandleFunc("/testadduser", func(w http.ResponseWriter, r *http.Request) {
		rpc.TestAddUser(10000)
		RenderDataJson(w, "ok")
	})
}
