package main

import (
	c "github.com/float1251/echo_sample/controller"
	"os"
	"reflect"
	"text/template"
)

type (
	TemplateArgument struct {
		Name   string
		Fields []StructField
	}

	StructField struct {
		reflect.StructField
		JsonTag string
	}
)

func main() {
	// 作成するStructの設定
	var list []interface{}

	list = append(list, c.UserCreateRequest{})
	list = append(list, c.UserLoginRequest{})

	var res []TemplateArgument
	for i := 0; i < len(list); i++ {
		t := reflect.TypeOf(list[i])
		arg := TemplateArgument{Name: t.Name()}
		num := t.NumField()
		arg.Fields = make([]StructField, num, num)
		for i := 0; i < num; i++ {
			// フィールドの取得
			f := t.Field(i)
			arg.Fields[i] = StructField{f, f.Tag.Get("json")}
		}
		res = append(res, arg)
	}

	tpl, err := template.New("Client Code").Parse(`
{{- range . }}
	[System.Serializable]
	public class {{.Name}}  {
		{{- range $v := .Fields }}
		public {{.Type}} {{.JsonTag}} { get; set;}
		{{- end }}
	}
{{- end }}
`)

	if err != nil {
		panic(err)
	}

	tpl.Execute(os.Stdout, res)
}
