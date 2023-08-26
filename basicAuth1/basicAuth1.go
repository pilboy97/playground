package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//basic authenication 구현

		id, pw, ok := r.BasicAuth()
		//basic 인증이 잘못 되었다면
		if !ok {
			w.Header().Set("WWW-Authenticate", "Basic realm=\"main page\"")
			w.WriteHeader(401)
			return
		}

		//basic 인증을 시도
		log.Println(id, pw)
		//id와 password를 확인
		if !(id == "pilki" && pw == "pw") {
			w.Header().Set("WWW-Authenticate", "Basic realm=\"main page\"")
			w.WriteHeader(401)
			return
		}
		//맞다면 메인 페이지 출력
		w.WriteHeader(200)
		w.Write([]byte{'h', 'i', '!'})
	})
	http.ListenAndServe(":8080", nil)
}
