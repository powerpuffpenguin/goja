package number

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type Template struct {
	dir string
}

func NewTemplate(dir string) *Template {
	return &Template{
		dir: dir,
	}
}
func (t *Template) Serve() {
	outGo := filepath.Join(t.dir, "..", "..", "zpp_expand_number.go")
	tGo := filepath.Join(t.dir, "internal", "number", "number.template")
	f, e := os.Create(outGo)
	if e != nil {
		log.Fatalln(e)
	}
	t.generation(f, tGo)
	f.Close()
}
func (t *Template) generation(w io.Writer, filename string) {
	b, e := ioutil.ReadFile(filename)
	if e != nil {
		log.Fatalln(e)
	}
	tm := template.New(filename)
	tm = tm.Funcs(builtins)
	tm, e = tm.Parse(string(b))
	if e != nil {
		log.Fatalln(e)
	}
	ctx := types
	e = tm.Execute(w, ctx)
	if e != nil {
		log.Fatalln(e)
	}
}

var builtins = template.FuncMap{
	"typeconv":  funcTypeconv,
	"typeconv2": funcTypeconv2,
}

func funcTypeconv(t Type, dstType, value string) string {
	if t.Type == dstType {
		return value
	}
	return fmt.Sprintf("%s(%s)", dstType, value)
}
func funcTypeconv2(t, dstType, value string) string {
	if t == dstType {
		return value
	}
	return fmt.Sprintf("%s(%s)", dstType, value)
}
