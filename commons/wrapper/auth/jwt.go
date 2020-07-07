package auth

import (
	"github.com/micro/micro/plugin"
	"log"
	"net/http"
	"strings"
)

// JWTAuthWrapper JWT鉴权Wrapper
func JWTAuthWrapper() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("auth plugin received: " + r.URL.Path)
			// TODO 从配置中心动态获取白名单URL
			if r.URL.Path == "/user/login" || r.URL.Path == "/user/register" || r.URL.Path == "/user/test" {
				h.ServeHTTP(w, r)
				return
			}

			// 简单验证
			if strings.Contains(r.URL.String(), "token") {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
