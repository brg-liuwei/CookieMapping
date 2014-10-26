package main

import (
	"fmt"
	"net/http"
	"runtime"
	//"time"
)

var cookieAddr string = "/cm"
var domain string = "http://localhost:9999"
var redirectAddr string = "/redir"

func CookieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Cookie invoke")

	cookies := r.Cookies()
	if len(cookies) == 0 {

		fmt.Println("len cookies = 0, set cookie")

		var cookie http.Cookie

		cookie.Name = "testCookie"
		cookie.Value = "hallo"

		/* 半个月失效 */
		cookie.MaxAge = 3600 * 24 * 15

		http.SetCookie(w, &cookie)

		cookie.Name = "testCookie2"
		cookie.Value = "hello"
		http.SetCookie(w, &cookie)

	} else {
		for i, c := range cookies {
			fmt.Printf("%d: cookie: %#v\n", i, *c)
		}
	}

	// redirect
	http.Redirect(w, r, domain+redirectAddr, 302)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("redirected")
	w.Write([]byte("Hello 302\n"))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc(cookieAddr, CookieHandler)
	http.HandleFunc(redirectAddr, RedirectHandler)
	panic(http.ListenAndServe(":9999", nil))
}
