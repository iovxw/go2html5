package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"

	. "github.com/Bluek404/go2html5"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().UnixNano()
		io.WriteString(w, index())
		log.Println(r.RemoteAddr, r.Method, r.URL, time.Now().UnixNano()-t, "ns")
	})

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index() string {
	html := Html_(Attr{"lang": "cn"},
		Head_(nil,
			Title_(nil, "Go2HTML5 Example"),
		),
		Body_(nil,
			P_(nil,
				"现在是：",
				func() string {
					var s bytes.Buffer
					t := time.Now().Hour()
					switch {
					case t >= 0 && t <= 4:
						s.WriteString("凌晨")
					case t >= 5 && t <= 7:
						s.WriteString("早上")
					case t >= 8 && t < 10:
						s.WriteString("上午")
					case t >= 11 && t <= 13:
						s.WriteString("中午")
					case t >= 14 && t <= 19:
						s.WriteString("下午")
					case t >= 20 && t <= 22:
						s.WriteString("晚上")
					case t >= 23 && t <= 24:
						s.WriteString("深夜")
					}
					s.WriteString(time.Now().Format("3点4分5秒"))
					return s.String()
				}(),
			),
			Hr_(nil),
			P_(nil,
				"Written in ", A_(Attr{"href": "http://golang.org"}, "Go!"),
			),
		),
	)
	return html
}
