Go2HTML5
========

[![Go Walker](https://img.shields.io/badge/Go%20Walker-API%20Documentation-green.svg?style=flat)](https://gowalker.org/github.com/Bluek404/go2html5)
[![GoDoc](https://img.shields.io/badge/GoDoc-API%20Documentation-blue.svg?style=flat)](http://godoc.org/github.com/Bluek404/go2html5)

脑洞产物

用go写html

除了普通的标签外

还可以直接写控件

所有函数均返回转换为HTML后的`string`

HTML5的所有标签正在逐步添加中

如果想实现自定义标签

可以自己写函数

例如：

```go
func CustomTag(attributes Attr, html []string) string {
    var tag = "custom-tag"
	var s bytes.Buffer
	s.WriteString("<" + tag)
	for k, v := range attributes {
		s.WriteString(" " + k + "=\"" + v + "\"")
	}
	s.WriteString(">")
	for _, v := range html {
		s.WriteString(v)
	}
	s.WriteString("</" + tag + ">")
	return s.String()
}
```

实现自定义控件：

```go
// 自定义按钮
func CustomButton(href, name string, color int) string {
	var c string
	switch color {
	case 0:
        c = "btn-red"
	case 1:
        c = "btn-yellow"
	case 2:
        c = "btn-green"
	case 3:
        c = "btn-blue"
    default:
        c = "btn-red"
	}

	return A(Attr{"class":"btn "+c, "href":href}, name)
}
```

Example
-------

```go
Html(Attr{"lang": "cn"},
	Head(nil,
		Title(nil, "Go2HTML5 Example"),
	),
	Body(nil,
		P(nil,
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
		Hr(nil),
		P(nil,
			"Written in ", A(Attr{"href": "http://golang.org"}, "Go!"),
		),
	),
)
```

更多请查看`example`文件夹
