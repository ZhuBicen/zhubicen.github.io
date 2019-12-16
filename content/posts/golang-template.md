---
title: "Golang Template"
date: 2011-12-08T15:52:30+08:00
draft: false
tags: [golang]
---

这个 package 的目的是使用结构化的数据作为输入，来生成一些具有固定格式的
文本。从而结构化的数据可以转化为我们需要的文本。

## 最简单的一个用法
生成一个template, 然后使用一个结构体变量，去替换template中的内容。
比如一个template中的内容如下，{{.NUmber}}, 如果使用一个结构体（该结构体有一个.Number成员），
{{.Number}}就会被，该变量的值就会替换。
##  概念
### template可以做什么？

template可以解析一些含有Action的UTF-8编码的文本或文件。并且借助一些参
数执行其中的action，以生成处理后的文件。其中使用的参数执行action所需要的。

action是输入的文本中有以"{{"和"}}"分隔起来的内容。
输入文本中除了action以外的所有内容都会原封不动的转换为输出


在下面的一个最简单的例子中 ，{{.Count}}和{{.Material}}就是action。而
sweaters就是在执行这些action时所使用的参数。template执行后的输出到
stdout 中。
```golang
  type Inventory struct {
      Material string
      Count    uint
  }
  sweaters := Inventory{"wool", 17}
  tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
  if err != nil { panic(err) }
  err = tmpl.Execute(os.Stdout, sweaters)
  if err != nil { panic(err) }
```
## Action是什么？
go中定义了多种形式的action,比如{{pipeline}}, 或者{{ if pipeline}} T1
{{end}}。其中pipeline是由一系列的command组成，command之间可以使用管道
连接，command可以是一个参数或者函数调用。
下面这段代码的输出是"output",原因是pipeline的输出是pipelin中最后一个
command的输出，

```
// {{with pipeline}} T1 {{end}}
	tmp := template.Must(template.New("test").Parse(`{{with $x := "output" | printf "%q"}}{{$x}}{{end}}`))
	err := tmp.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
```

Execute函数中传入的参数，在执行开始后，在action中被$接收，
```
templateText := `<ul>{{ range $e := $ }}<li> {{ $e }} </li>{{ end }}</ul>`
	tmp, _ := template.New("test").Parse(templateText)
	tmp.Execute(os.Stdout, []string{"item1", "item2"})
```
在action中一般不包括代码逻辑，但是在action中可以调用外部的函数。
比如下面的代码输出 "The Go Programming Language"。
```
var text string = `{{pipeline}}`
func Title() string{
	return "The Go Programming Language"
}

func main() {
	funcMap := template.FuncMap {
		"pipeline":Title,
	}
	tmp := template.Must(template.New("test").Funcs(funcMap).Parse(text))
	err := tmp.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
```
当然也可以调用有输入参数的参数。比如`{{ pipeline . }}`或者 `{{
pipeline $}}`。

另外，使用action中也可以做一些抽象，类似我们代码中的函数。
```golang
text := `{{ define "title" }} {{ $ }} {{ end }}  {{ template "title" $ }} `
	tmp := template.Must(template.New("test").Parse(text))
	tmp.Execute(os.Stdout, "The Go Programming Language")
```