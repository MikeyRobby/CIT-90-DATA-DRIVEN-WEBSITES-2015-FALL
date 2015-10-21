package main

import (
	"github.com/gorilla/securecookie"
	"io"
	"net/http"
	"fmt"
)

var hashKey = securecookie.GenerateRandomKey(32)
var blockKey = securecookie.GenerateRandomKey(32)
var s = securecookie.New(hashKey, blockKey)

func main() {
	http.HandleFunc("/", SetCookieHandler)
	http.ListenAndServe(":8080", nil)

}

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		"foo": "bar",
	}
	if encoded, err := s.Encode("E-mail", value); err == nil {
		cookie := &http.Cookie{
			Name: "E-mail",
			Value: encoded,
			Path: "/",
		}

		http.SetCookie(w, cookie)

		io.WriteString(w, `<!DOCTYPE html>
		<html>
			<body>
				<form method="POST">
				<p>Enter Your Email:</p>
					<input type="email" name="email">
					<input type="submit"
				</form>
			</body>
		</html>`)
	}
}
	func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
		if cookie, err := r.Cookie("E-mail"); err == nil {
			value := make(map[string]string)

			if err = s.Decode("E-mail", cookie.Value, &value); err == nil {
				fmt.Fprintf(w, "The value of foo is %q", value["foo"])

			}
		}
	}








