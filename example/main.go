package main

import (
	"fmt"
	"time"

	"bytes"
	. "github.com/Bluek404/go2html5"
)

func main() {
	html := Html(Attr{"lang": "en"},
		Head(nil,
			Title(nil, "Go2HTML5 Example"),
		),
		Body(nil,
			P(nil,
				"现在是：",
				func() string {
					var s bytes.Buffer
					time := time.Now().Hour()
					switch {
					case time >= 0 && time <= 4:
						s.WriteString("凌晨")
					case time >= 5 && time <= 7:
						s.WriteString("早上")
					case time >= 8 && time < 10:
						s.WriteString("上午")
					case time >= 11 && time <= 13:
						s.WriteString("中午")
					case time >= 14 && time <= 19:
						s.WriteString("下午")
					case time >= 20 && time <= 22:
						s.WriteString("晚上")
					case time >= 23 && time <= 24:
						s.WriteString("深夜")
					}
					return s.String()
				}(),
			),
			Hr(nil),
			P(nil,
				"Written in ", A(Attr{"href": "http://golang.org"}, "Go!"),
			),
		),
	)
	fmt.Println(html)
}
