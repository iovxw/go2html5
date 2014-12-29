/*
 Copyright 2014 Bluek404

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

// HTML5中的标签
// 标签注释文档来自 www.w3school.com.cn
package go2html5

import (
	"bytes"
)

// 定义锚
//
// <a> 标签定义超链接，用于从一张页面链接到另一张页面。
func A(attributes Attr, html ...string) string {
	return d("a", attributes, html)
}

// 定义缩写
//
// <abbr> 标签指示简称或缩写，比如 "WWW" 或 "NATO"。
// 通过对缩写进行标记，您能够为浏览器、拼写检查和搜索引擎提供有用的信息。
// <abbr> 标签最初是在 HTML 4.0 中引入的，表示它所包含的文本是一个更长的单词或短语的缩写形式。
func Abbr(attributes Attr, html ...string) string {
	return d("abbr", attributes, html)
}

// 定义文档作者或拥有者的联系信息
//
// <address> 标签定义文档或文章的作者/拥有者的联系信息。
// 如果 <address> 元素位于 <body> 元素内，则它表示文档联系信息。
// 如果 <address> 元素位于 <article> 元素内，则它表示文章的联系信息。
// <address> 元素中的文本通常呈现为斜体。大多数浏览器会在 address 元素前后添加折行。
func Address(attributes Attr, html ...string) string {
	return d("address", attributes, html)
}

// 定义图像映射内部的区域
//
// <area> 标签定义图像映射中的区域（注：图像映射指得是带有可点击区域的图像）。
// area 元素总是嵌套在 <map> 标签中。
func Area(attributes Attr, html ...string) string {
	return d("area", attributes, html)
}

// 定义文章
//
// <article> 标签规定独立的自包含内容。
// 一篇文章应有其自身的意义，应该有可能独立于站点的其余部分对其进行分发。
// <article> 元素的潜在来源：
//   - 论坛帖子
//   - 报纸文章
//   - 博客条目
//   - 用户评论
func Article(attributes Attr, html ...string) string {
	return d("article", attributes, html)
}

// 定义声音内容
//
// <<audio> 标签定义声音，比如音乐或其他音频流。
func Audio(attributes Attr, html ...string) string {
	return d("audio", attributes, html)
}

// 定义粗体字
//
// <b> 标签定义粗体的文本。
// <b> 标签定义了文本中的部分比其余的部分更重要，并呈现为粗体。
func B(attributes Attr, html ...string) string {
	return d("b", attributes, html)
}

// 定义页面中所有链接的默认地址或默认目标
//
// <base> 标签为页面上的所有链接规定默认地址或默认目标。
// 通常情况下，浏览器会从当前文档的 URL 中提取相应的元素来填写相对 URL 中的空白。
// 使用 <base> 标签可以改变这一点。浏览器随后将不再使用当前文档的 URL，而使用指定的基本 URL 来解析所有的相对 URL。这其中包括 <a>、<img>、<link>、<form> 标签中的 URL。
func Base(attributes Attr, html ...string) string {
	return d("base", attributes, html)
}

// 定义文本的文本方向，使其脱离其周围文本的方向设置
//
// bdi 指的是 bidi 隔离。
// <bdi> 标签允许您设置一段文本，使其脱离其父元素的文本方向设置。
// 在发布用户评论或其他您无法完全控制的内容时，该标签很有用。
func Bdi(attributes Attr, html ...string) string {
	return d("bdi", attributes, html)
}

// 定义文字方向
//
// bdo 元素可覆盖默认的文本方向。
func Bdo(attributes Attr, html ...string) string {
	return d("bdo", attributes, html)
}

// 定义长的引用
//
// <blockquote> 标签定义块引用。
// <blockquote> 与 </blockquote> 之间的所有文本都会从常规文本中分离出来，经常会在左、右两边进行缩进（增加外边距），而且有时会使用斜体。
// 也就是说，块引用拥有它们自己的空间。
func Blockquote(attributes Attr, html ...string) string {
	return d("blockquote", attributes, html)
}

// 定义文档的主体
//
// <body> 元素定义文档的主体。
// <body> 元素包含文档的所有内容，比如文本、超链接、图像、表格、列表等等。
func Body(attributes Attr, html ...string) string {
	return d("body", attributes, html)
}

// 定义简单的折行
//
// <br/> 标签插入简单的换行符。
// <br/> 标签是一个空标签，意味着它没有结束标签。所以这样是错误的：<br></br>
func Br(attributes Attr) string {
	var s bytes.Buffer
	s.WriteString("<br")
	for k, v := range attributes {
		s.WriteString(" " + k + "=\"" + v + "\"")
	}
	s.WriteString("/>")
	return s.String()
}

// 定义按钮
//
// <button> 标签定义按钮。
// 您可以在 button 元素放置内容，比如文本或图像。这是该元素与通过 input 元素创建的按钮的不同之处。
// 请始终为按钮规定 type 属性。不同的浏览器根据 type 属性使用不同的默认值。
func Button(attributes Attr, html ...string) string {
	return d("button", attributes, html)
}

// 定义图形
//
// <canvas> 标签定义图形，比如图表和其他图像。
// <canvas> 标签只是图形容器，您必须使用脚本来绘制图形。
func Canvas(attributes Attr, html ...string) string {
	return d("canvas", attributes, html)
}

// 定义引用
//
// <cite> 标签定义作品（比如书籍、歌曲、电影、电视节目、绘画、雕塑等等）的标题。
func Cite(attributes Attr, html ...string) string {
	return d("cite", attributes, html)
}

// 定义计算机代码文本
//
// <code> 标签用于表示计算机源代码或者其他机器可以阅读的文本内容。
// 软件代码的编写者已经习惯了编写源代码时文本表示的特殊样式。<code> 标签就是为他们设计的。
// 包含在该标签内的文本将用等宽、类似电传打字机样式的字体（Courier）显示出来，对于大多数程序员来说，这应该是十分熟悉的。
// 只应该在表示计算机程序源代码或者其他机器可以阅读的文本内容上使用 <code> 标签。虽然 <code> 标签通常只是把文本变成等宽字体，但它暗示着这段文本是源程序代码。
// 将来的浏览器有可能会加入其他显示效果。例如，程序员的浏览器可能会寻找 <code> 片段，并执行某些额外的文本格式化处理，如循环和条件判断语句的特殊缩进等。
func Code(attributes Attr, html ...string) string {
	return d("code", attributes, html)
}

// 定义表格中一个或多个列的属性值
//
// <col> 标签为表格中的一个或多个列定义属性值。您只能在表格或列组中使用该元素。
func Col(attributes Attr) string {
	var s bytes.Buffer
	s.WriteString("<col")
	for k, v := range attributes {
		s.WriteString(" " + k + "=\"" + v + "\"")
	}
	s.WriteString("/>")
	return s.String()
}

// 定义表格中供格式化的列组
//
//<colgroup> 标签定义表格列的组。通过此标签，您可以对列进行组合，以便格式化。该元素只有在 <table> 中才是合法的。
func Colgroup(attributes Attr, html ...string) string {
	return d("colgroup", attributes, html)
}

// 定义命令按钮
//
// <command> 标签定义命令按钮，比如单选按钮、复选框或按钮。
func Command(attributes Attr, html ...string) string {
	return d("command", attributes, html)
}

// 定义可选数据列表
//
// <datagrid> 标签定义可选数据的列表。datagrid 作为树列表来显示。
func Datagrid(attributes Attr, html ...string) string {
	return d("datagrid", attributes, html)
}

// 定义下拉列表
//
// <datalist> 标签定义可选数据的列表。与 input 元素配合使用，就可以制作出输入值的下拉列表。
func Datalist(attributes Attr, html ...string) string {
	return d("datalist", attributes, html)
}

// 定义数据模板容器
//
// <datatemplate> 标签定义数据模板的一个容器。该元素必须有定义模板的子元素：<rule> 元素。
func Datatemplate(attributes Attr, html ...string) string {
	return d("datatemplate", attributes, html)
}

// 定义定义列表中项目的描述
//
// <dd> 标签定义一个定义列表中对项目的描述。
func Dd(attributes Attr, html ...string) string {
	return d("dd", attributes, html)
}

// 定义被删除文本
//
// <dd> 标签定义文档中已删除的文本。
func Del(attributes Attr, html ...string) string {
	return d("del", attributes, html)
}

// 定义元素的细节
//
// <details> 标签定义元素的细节，用户可进行查看，或通过点击进行隐藏。
func Details(attributes Attr, html ...string) string {
	return d("details", attributes, html)
}

// 定义对话
//
// <dialog> 标签定义对话，比如交谈。
func Dialog(attributes Attr, html ...string) string {
	return d("dialog", attributes, html)
}

// 定义一个定义项目
//
// <dfn> 标签可标记那些对特殊术语或短语的定义。
// 现在流行的浏览器通常用斜体来显示 <dfn> 中的文本。将来，<dfn> 还可能有助于创建文档的索引或术语表。
// 与其他许多基于内容的样式和物理样式标签一样，<dfn> 标签尽量少用为妙。
// 作为一种通用样式，尤其在技术文档中，当第一次新的术语时，应该将它们与普通文本分开，这样读者可以更好地理解文章当前的主题，而从那以后就不要再对这个术语进行任何标记了。
func Dfn(attributes Attr, html ...string) string {
	return d("dfn", attributes, html)
}

// 定义文档中的节
//
// <div> 可定义文档中的分区或节（division/section）。
// <div> 标签可以把文档分割为独立的、不同的部分。它可以用作严格的组织工具，并且不使用任何格式与其关联。
// 如果用 id 或 class 来标记 <div>，那么该标签的作用会变得更加有效。
func Div(attributes Attr, html ...string) string {
	return d("div", attributes, html)
}

// 定义定义列表
//
// <dl> 标签定义了定义列表（definition list）。
// <dl> 标签用于结合 <dt> （定义列表中的项目）和 <dd> （描述列表中的项目）。
func Dl(attributes Attr, html ...string) string {
	return d("dl", attributes, html)
}

// 定义定义列表中的项目
//
// <dt> 标签定义了定义列表中的项目（即术语部分）。
func Dt(attributes Attr, html ...string) string {
	return d("dt", attributes, html)
}

// 定义强调文本
//
// <em> 标签告诉浏览器把其中的文本表示为强调的内容。对于所有浏览器来说，这意味着要把这段文字用斜体来显示。
func Em(attributes Attr, html ...string) string {
	return d("em", attributes, html)
}

// 定义 HTML 文档
//
// 此元素可告知浏览器其自身是一个 HTML 文档。
// 强制HTML版本为5，自带<!DOCTYPE HTML>
func Html(attributes Attr, html ...string) string {
	var s bytes.Buffer
	s.WriteString("<!DOCTYPE HTML>")
	s.WriteString(d("html", attributes, html))
	return s.String()
}

func d(tag string, attributes Attr, html []string) string {
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

type Attr map[string]string
