package stmpl

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
	"testing"
	"text/template"
)

//go:embed tmpltest.go.tmpl
var tpl string

func TestTemplate(t *testing.T) {
	// 创建模板
	fmt.Println(tpl)
	tmpl, err := template.New("my.gotpl").Parse(tpl)
	if err != nil || tmpl == nil {
		t.Errorf("tmpl create fail!")
		return
	}

	// 根据模板输出内容
	buf := new(bytes.Buffer)
	meta := &goFileMeta{
		Name: "quaint",
		Pkg:  "gosdk",
		Imports: []string{
			"fmt",
		},
	}
	if err := tmpl.Execute(buf, meta); err != nil {
		t.Errorf("tmpl create fail!")
		return
	}

	// 写go文件
	src, err := format.Source(buf.Bytes())
	filename := fmt.Sprintf("gen_%s.go", strings.ToLower(meta.Name))
	err = ioutil.WriteFile(filename, src, 0644)
	if err != nil {
		t.Errorf("write file fail!")
		return
	}
}

type goFileMeta struct {
	Name    string
	Pkg     string
	Imports []string
}
