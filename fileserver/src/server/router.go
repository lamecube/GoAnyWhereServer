package server

import (
	"log"
	"net/http"
	"strings"
	// "fileserver/template/status"
)

func Home(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, r)
	log.Print("Yo")
	w.Write([]byte("OK"))
}

func ApiRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			SetHeaders(w, r)
			// w.Header().Set("content-type","text/html; charset=utf-8")
			// status.S_404(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func FileServer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			SetHeaders(w, r)
			// w.Header().Set("content-type","text/html; charset=utf-8")
			// status.S_404(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
