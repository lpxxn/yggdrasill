package code_gen

import (
	"bytes"
	"errors"
	"strings"
	"text/template"
	"time"

	"github.com/lpxxn/yggdrasill/code_gen/mysql"
	"github.com/lpxxn/yggdrasill/code_gen/postgresql"
	"github.com/lpxxn/yggdrasill/utils"
)

func NewDbCodeGen(t string) (IDBMetaData, error) {
	switch strings.ToLower(t) {
	case "mysql":
		return &mysql.Gen{}, nil
	case "pg", "postgresql":
		return &postgresql.PGGen{}, nil
	}
	return nil, errors.New("invalid type")
}

func GenerateTemplate(templateText string, templateData interface{}, params map[string]interface{}) ([]byte, error) {
	t, err := template.New("tableTemplate").Funcs(template.FuncMap{
		"CamelizeStr":    utils.CamelizeStr,
		"FirstCharacter": utils.FirstCharacter,
		"Replace": func(old, new, src string) string {
			return strings.ReplaceAll(src, old, new)
		},
		"now": func() string {
			return time.Now().Format(time.RFC3339)
		},
		"param": func(name string) interface{} {
			if v, ok := params[name]; ok {
				return v
			}
			return ""
		},
	}).Parse(templateText)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, templateData); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
