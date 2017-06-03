package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var privateHtml = `
こんにちわ %s さん<br />
<a href="/login?logout=true">ログアウト</a>
`

var publicHtml = `
<a href="/login">ログイン</a>
`

func main() {
	secret := "deadbeef"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html; charset=utf8")
		if c, err := r.Cookie("jwt_token"); err == nil {
			token, err := jwt.Parse(c.Value, func(*jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Fprintf(w, privateHtml, claims["sub"])
				return
			}
		}
		fmt.Fprintln(w, publicHtml)
	})
	http.ListenAndServe(":8888", nil)
}
