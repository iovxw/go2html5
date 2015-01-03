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

// <a> 标签定义超链接，用于从一张页面链接到另一张页面。
func A_(attributes Attr, html ...string) string {
	return d("a", attributes, html)
}

// <abbr> 标签指示简称或缩写，比如 "WWW" 或 "NATO"。
// 通过对缩写进行标记，您能够为浏览器、拼写检查和搜索引擎提供有用的信息。
// <abbr> 标签最初是在 HTML 4.0 中引入的，表示它所包含的文本是一个更长的单词或短语的缩写形式。
func Abbr_(attributes Attr, html ...string) string {
	return d("abbr", attributes, html)
}

// <address> 标签定义文档或文章的作者/拥有者的联系信息。
// 如果 <address> 元素位于 <body> 元素内，则它表示文档联系信息。
// 如果 <address> 元素位于 <article> 元素内，则它表示文章的联系信息。
// <address> 元素中的文本通常呈现为斜体。大多数浏览器会在 address 元素前后添加折行。
func Address_(attributes Attr, html ...string) string {
	return d("address", attributes, html)
}

// <area> 标签定义图像映射中的区域（注：图像映射指得是带有可点击区域的图像）。
// area 元素总是嵌套在 <map> 标签中。
func Area_(attributes Attr, html ...string) string {
	return d("area", attributes, html)
}

// <article> 标签规定独立的自包含内容。
// 一篇文章应有其自身的意义，应该有可能独立于站点的其余部分对其进行分发。
// <article> 元素的潜在来源：
//   - 论坛帖子
//   - 报纸文章
//   - 博客条目
//   - 用户评论
func Article_(attributes Attr, html ...string) string {
	return d("article", attributes, html)
}

// <audio> 标签定义声音，比如音乐或其他音频流。
func Audio_(attributes Attr, html ...string) string {
	return d("audio", attributes, html)
}

// <b> 标签定义强调的文本。
// <b> 标签定义了文本中的部分比其余的部分更重要，并呈现为粗体。
func B_(attributes Attr, html ...string) string {
	return d("b", attributes, html)
}

// <base> 标签为页面上的所有链接规定默认地址或默认目标。
// 通常情况下，浏览器会从当前文档的 URL 中提取相应的元素来填写相对 URL 中的空白。
// 使用 <base> 标签可以改变这一点。浏览器随后将不再使用当前文档的 URL，而使用指定的基本 URL 来解析所有的相对 URL。这其中包括 <a>、<img>、<link>、<form> 标签中的 URL。
func Base_(attributes Attr, html ...string) string {
	return d("base", attributes, html)
}

// <bdi> 标签允许您设置一段文本，使其脱离其父元素的文本方向设置。
// bdi 指的是 bidi 隔离。
// 在发布用户评论或其他您无法完全控制的内容时，该标签很有用。
func Bdi_(attributes Attr, html ...string) string {
	return d("bdi", attributes, html)
}

// <bdo> 标签可覆盖默认的文本方向。
func Bdo_(attributes Attr, html ...string) string {
	return d("bdo", attributes, html)
}

// <blockquote> 标签定义块引用。
// <blockquote> 与 </blockquote> 之间的所有文本都会从常规文本中分离出来，经常会在左、右两边进行缩进（增加外边距），而且有时会使用斜体。
// 也就是说，块引用拥有它们自己的空间。
func Blockquote_(attributes Attr, html ...string) string {
	return d("blockquote", attributes, html)
}

// <body> 元素定义文档的主体。
// <body> 元素包含文档的所有内容，比如文本、超链接、图像、表格、列表等等。
func Body_(attributes Attr, html ...string) string {
	return d("body", attributes, html)
}

// <br/> 标签插入简单的换行符。
// <br/> 标签是一个空标签，意味着它没有结束标签。所以这样是错误的：<br></br>
func Br_(attributes Attr) string {
	var s bytes.Buffer
	s.WriteString("<br")
	for k, v := range attributes {
		s.WriteString(" " + k + "=\"" + v + "\"")
	}
	s.WriteString("/>")
	return s.String()
}

// <button> 标签定义按钮。
// 您可以在 button 元素放置内容，比如文本或图像。这是该元素与通过 input 元素创建的按钮的不同之处。
// 请始终为按钮规定 type 属性。不同的浏览器根据 type 属性使用不同的默认值。
func Button_(attributes Attr, html ...string) string {
	return d("button", attributes, html)
}

// <canvas> 标签定义图形，比如图表和其他图像。
// <canvas> 标签只是图形容器，您必须使用脚本来绘制图形。
func Canvas_(attributes Attr, html ...string) string {
	return d("canvas", attributes, html)
}

// <cite> 标签定义作品（比如书籍、歌曲、电影、电视节目、绘画、雕塑等等）的标题。
func Cite_(attributes Attr, html ...string) string {
	return d("cite", attributes, html)
}

// <code> 标签用于表示计算机源代码或者其他机器可以阅读的文本内容。
// 软件代码的编写者已经习惯了编写源代码时文本表示的特殊样式。<code> 标签就是为他们设计的。
// 包含在该标签内的文本将用等宽、类似电传打字机样式的字体（Courier）显示出来，对于大多数程序员来说，这应该是十分熟悉的。
// 只应该在表示计算机程序源代码或者其他机器可以阅读的文本内容上使用 <code> 标签。虽然 <code> 标签通常只是把文本变成等宽字体，但它暗示着这段文本是源程序代码。
// 将来的浏览器有可能会加入其他显示效果。例如，程序员的浏览器可能会寻找 <code> 片段，并执行某些额外的文本格式化处理，如循环和条件判断语句的特殊缩进等。
func Code_(attributes Attr, html ...string) string {
	return d("code", attributes, html)
}

// <col> 标签为表格中的一个或多个列定义属性值。您只能在表格或列组中使用该元素。
func Col_(attributes Attr) string {
	var s bytes.Buffer
	s.WriteString("<col")
	for k, v := range attributes {
		s.WriteString(" " + k + "=\"" + v + "\"")
	}
	s.WriteString("/>")
	return s.String()
}

//<colgroup> 标签定义表格列的组。通过此标签，您可以对列进行组合，以便格式化。该元素只有在 <table> 中才是合法的。
func Colgroup_(attributes Attr, html ...string) string {
	return d("colgroup", attributes, html)
}

// <command> 标签定义命令按钮，比如单选按钮、复选框或按钮。
func Command_(attributes Attr, html ...string) string {
	return d("command", attributes, html)
}

// <datagrid> 标签定义可选数据的列表。datagrid 作为树列表来显示。
func Datagrid_(attributes Attr, html ...string) string {
	return d("datagrid", attributes, html)
}

// <datalist> 标签定义可选数据的列表。与 input 元素配合使用，就可以制作出输入值的下拉列表。
func Datalist_(attributes Attr, html ...string) string {
	return d("datalist", attributes, html)
}

// <datatemplate> 标签定义数据模板的一个容器。该元素必须有定义模板的子元素：<rule> 元素。
func Datatemplate_(attributes Attr, html ...string) string {
	return d("datatemplate", attributes, html)
}

// <dd> 标签定义一个定义列表中对项目的描述。
func Dd_(attributes Attr, html ...string) string {
	return d("dd", attributes, html)
}

// <dd> 标签定义文档中已删除的文本。
func Del_(attributes Attr, html ...string) string {
	return d("del", attributes, html)
}

// <details> 标签定义元素的细节，用户可进行查看，或通过点击进行隐藏。
func Details_(attributes Attr, html ...string) string {
	return d("details", attributes, html)
}

// <dialog> 标签定义对话，比如交谈。
func Dialog_(attributes Attr, html ...string) string {
	return d("dialog", attributes, html)
}

// <dfn> 标签可标记那些对特殊术语或短语的定义。
// 现在流行的浏览器通常用斜体来显示 <dfn> 中的文本。将来，<dfn> 还可能有助于创建文档的索引或术语表。
// 与其他许多基于内容的样式和物理样式标签一样，<dfn> 标签尽量少用为妙。
// 作为一种通用样式，尤其在技术文档中，当第一次新的术语时，应该将它们与普通文本分开，这样读者可以更好地理解文章当前的主题，而从那以后就不要再对这个术语进行任何标记了。
func Dfn_(attributes Attr, html ...string) string {
	return d("dfn", attributes, html)
}

// <div> 可定义文档中的分区或节（division/section）。
// <div> 标签可以把文档分割为独立的、不同的部分。它可以用作严格的组织工具，并且不使用任何格式与其关联。
// 如果用 id 或 class 来标记 <div>，那么该标签的作用会变得更加有效。
func Div_(attributes Attr, html ...string) string {
	return d("div", attributes, html)
}

// <dl> 标签定义了定义列表（definition list）。
// <dl> 标签用于结合 <dt> （定义列表中的项目）和 <dd> （描述列表中的项目）。
func Dl_(attributes Attr, html ...string) string {
	return d("dl", attributes, html)
}

// <dt> 标签定义了定义列表中的项目（即术语部分）。
func Dt_(attributes Attr, html ...string) string {
	return d("dt", attributes, html)
}

// <em> 标签告诉浏览器把其中的文本表示为强调的内容。对于所有浏览器来说，这意味着要把这段文字用斜体来显示。
func Em_(attributes Attr, html ...string) string {
	return d("em", attributes, html)
}

// <embed> 标签定义嵌入的内容，比如插件。
func Embed_(attributes Attr) string {
	var s bytes.Buffer
	s.WriteString("<embed")
	for k, v := range attributes {
		s.WriteString(" " + k + "=\"" + v + "\"")
	}
	s.WriteString("/>")
	return s.String()
}

// <event-source> 标签定义由服务器发送的事件的来源。
func EventSource_(attributes Attr, html ...string) string {
	return d("event-source", attributes, html)
}

// <fieldset> 标签可将表单内的相关元素分组。
// <fieldset> 标签将表单内容的一部分打包，生成一组相关表单的字段。<fieldset> 标签没有必需的或唯一的属性。
// 当一组表单元素放到 <fieldset> 标签内时，浏览器会以特殊方式来显示它们，它们可能有特殊的边界、3D 效果，或者甚至可创建一个子表单来处理这些元素。
func Fieldset_(attributes Attr, html ...string) string {
	return d("fieldset", attributes, html)
}

// <figcaption> 标签定义 figure 元素的标题（caption）。
// figcaption 元素应该被置于 figure 元素的第一个或最后一个子元素的位置。
func Figcaption_(attributes Attr, html ...string) string {
	return d("figcaption", attributes, html)
}

// <figure> 标签规定独立的流内容（图像、图表、照片、代码等等）。
// figure 元素的内容应该与主内容相关，但如果被删除，则不应对文档流产生影响。
func Figure_(attributes Attr, html ...string) string {
	return d("figure", attributes, html)
}

// <footer> 标签定义 section 或 document 的页脚。
// 典型地，它会包含创作者的姓名、文档的创作日期以及/或者联系信息。
func Footer_(attributes Attr, html ...string) string {
	return d("footer", attributes, html)
}

// <form> 标签创建供用户输入的表单。
// 表单可包含文本域，复选框，单选按钮等等。表单用于向指定的 URL 传递用户数据。
func Form_(attributes Attr, html ...string) string {
	return d("form", attributes, html)
}

// <h1> - <h6> 标签可定义标题。<h1> 定义最大的标题。<h6> 定义最小的标题
func H1_(attributes Attr, html ...string) string {
	return d("h1", attributes, html)
}

// <h1> - <h6> 标签可定义标题。<h1> 定义最大的标题。<h6> 定义最小的标题
func H2_(attributes Attr, html ...string) string {
	return d("h2", attributes, html)
}

// <h1> - <h6> 标签可定义标题。<h1> 定义最大的标题。<h6> 定义最小的标题
func H3_(attributes Attr, html ...string) string {
	return d("h3", attributes, html)
}

// <h1> - <h6> 标签可定义标题。<h1> 定义最大的标题。<h6> 定义最小的标题
func H4_(attributes Attr, html ...string) string {
	return d("h4", attributes, html)
}

// <h1> - <h6> 标签可定义标题。<h1> 定义最大的标题。<h6> 定义最小的标题
func H5_(attributes Attr, html ...string) string {
	return d("h5", attributes, html)
}

// <h1> - <h6> 标签可定义标题。<h1> 定义最大的标题。<h6> 定义最小的标题
func H6_(attributes Attr, html ...string) string {
	return d("h6", attributes, html)
}

// <head> 标签用于定义文档的头部，它是所有头部元素的容器。<head> 中的元素可以引用脚本、指示浏览器在哪里找到样式表、提供元信息等等。
// 文档的头部描述了文档的各种属性和信息，包括文档的标题、在 Web 中的位置以及和其他文档的关系等。绝大多数文档头部包含的数据都不会真正作为内容显示给读者。
// 下面这些标签可用在 head 部分：<base>, <link>, <meta>, <script>, <style>, 以及 <title>。
// <title> 定义文档的标题，它是 head 部分中唯一必需的元素。
func Head_(attributes Attr, html ...string) string {
	return d("head", attributes, html)
}

// <header> 标签定义 section 或 document 的页眉。
func Header_(attributes Attr, html ...string) string {
	return d("header", attributes, html)
}

// <hr> 标签在 HTML 页面中创建一条水平线。
// 水平分隔线（horizontal rule）可以在视觉上将文档分隔成各个部分。
func Hr_(attributes Attr) string {
	var s bytes.Buffer
	s.WriteString("<hr")
	for k, v := range attributes {
		s.WriteString(" " + k + "=\"" + v + "\"")
	}
	s.WriteString("/>")
	return s.String()
}

// <html> 标签定义 HTML 文档
// 此元素可告知浏览器其自身是一个 HTML 文档。
// 强制HTML版本为5，自带<!DOCTYPE HTML>
func Html_(attributes Attr, html ...string) string {
	var s bytes.Buffer
	s.WriteString("<!DOCTYPE HTML>")
	s.WriteString(d("html", attributes, html))
	return s.String()
}

// <i> 标签呈现斜体的文本。
// <i> 标签定义与文本中其余部分不同的部分，并把这部分文本呈现为斜体文本。
// 在没有其他元素可以使用时，比如 <b>, <cite>, <dfn>, <em>, <q>, <small>, <strong>，请使用 <i> 标签。
func I_(attributes Attr, html ...string) string {
	return d("i", attributes, html)
}

// <iframe> 标签创建包含另一个文档的行内框架。
func Iframe_(attributes Attr, html ...string) string {
	return d("iframe", attributes, html)
}

// <img> 标签定义图像。
func Img_(attributes Attr) string {
	var s bytes.Buffer
	s.WriteString("<img")
	for k, v := range attributes {
		s.WriteString(" " + k + "=\"" + v + "\"")
	}
	s.WriteString("/>")
	return s.String()
}

// <input> 标签定义输入字段，用户可在其中输入数据。
func Input_(attributes Attr, html ...string) string {
	return d("input", attributes, html)
}

// <ins> 标签定义文档的插入的新文本
func Ins_(attributes Attr, html ...string) string {
	return d("ins", attributes, html)
}

// <kbd> 标签定义键盘文本。
// 说到技术概念上的特殊样式时，就要提到 <kbd> 标签。正如你已经猜到的，它用来表示文本是从键盘上键入的。
// 浏览器通常用等宽字体来显示该标签中包含的文本。
// <kbd> 标签经常用在于计算机相关的文档和手册中。例如：
//   键入 <kbd>quit</kbd> 来退出程序，或者键入 <kbd>menu</kbd> 来返回主菜单。
func Kbd_(attributes Attr, html ...string) string {
	return d("kbd", attributes, html)
}

// <label> 标签定义控件的标注。如果您在 label 元素内点击文本，就会触发此控件。
// 就是说，当用户选择该标签时，浏览器就会自动将焦点转到和标签相关的表单控件上。
func Label_(attributes Attr, html ...string) string {
	return d("label", attributes, html)
}

// <legend> 标签为 <fieldset>、<figure> 以及 <details> 元素定义标题。
func Legend_(attributes Attr, html ...string) string {
	return d("legend", attributes, html)
}

// <p> 标签会自动在其前后创建一些空白。浏览器会自动添加这些空间，您也可以在样式表中规定。
func P_(attributes Attr, html ...string) string {
	return d("p", attributes, html)
}

// <title> 元素可定义文档的标题。
// 浏览器会以特殊的方式来使用标题，并且通常把它放置在浏览器窗口的标题栏或状态栏上。同样，当把文档加入用户的链接列表或者收藏夹或书签列表时，标题将成为该文档链接的默认名称。
func Title_(attributes Attr, html ...string) string {
	return d("title", attributes, html)
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

// HTML标签属性
type Attr map[string]string
