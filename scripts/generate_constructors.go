package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var curAppVersion = `"0.0.1"`
var curSchemaVersion = `"0.0.1"`

var genTypesFile = "gen_types"

var tmpl = template.Must(template.New("constructor").Parse(`
func New{{ .Name }}() {{ .Name }} {
	return {{ .Name }}{
		{{- range .Fields }}
		{{ .Name }}: {{ .ZeroValue }},
		{{- end }}
	}
}
`))

var header = `
package models

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"time"
	{imports}
)

func Ptr[T any](v T) *T {
    return &v
}

func getUuid() string {
	id, err := gonanoid.New()
	if err != nil {
		id = ""
	}
	return id
}

`

type Field struct {
	Name      string
	ZeroValue string
}

type StructData struct {
	Name   string
	Fields []Field
}

var debug = false
var useTime = false

var enumProps = []string{"TabType", "WidgetType", "TabDefinitionType", "WidgetDefinitionType", "TabDefType", "WidgetDefType"}
var typeSliceProps = []string{"EditorSession{}", "AppStateChange{}", "TabDefinition{}", "WidgetDefinition{}", "TabDef{}", "WidgetDef{}"}
var timeStampZero = `time.Now().Format(time.RFC3339)`
var timeStampZeroPtr = `Ptr(time.Now().Format(time.RFC3339))`
var metaZeros = map[string]string{
	"UUID":           `getUuid()`,
	"SchemaVersion":  curSchemaVersion,
	"AppVersion":     curAppVersion,
	"CreatedOn":      timeStampZero,
	"LastModifiedOn": timeStampZero,
	"CreatedBy":      `Ptr("bot-builder")`,
	"LastModifiedBy": `Ptr("bot-builder")`,
	"Status":         `Ptr(Active)`,
}

func debugPrint(items ...any) {
	if debug {
		fmt.Println(items...)
	}
}

func contains(items []string, value string) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false
}

func zeroValue(expr ast.Expr) string {

	switch t := expr.(type) {
	case *ast.Ident:
		switch t.Name {
		case "string":
			return `""`
		case "int", "int32", "int64", "float64", "float32":
			return "0"
		case "bool":
			return "false"
		default:
			if contains(enumProps, t.Name) {
				return `""`
			}
			debugPrint("Unknown type ident:", t.Name)
			return t.Name + "{}"
		}

	case *ast.StarExpr:
		switch inner := t.X.(type) {
		case *ast.Ident:
			switch inner.Name {
			case "string":
				return `Ptr("")`
			case "int", "int32", "int64", "float64", "float32":
				return `Ptr(` + fmt.Sprintf("%s(0)", inner.Name) + `)`
			case "bool":
				return `Ptr(false)`
			default:
				debugPrint("Unknown type star:", inner.Name)
				return "nil"
			}
		default:
			return "nil"
		}

	case *ast.ArrayType:
		zero := zeroValue(t.Elt)
		debugPrint("Array type:", zero, t.Elt)
		if zero == "nil" {
			zero = "any{}"
		} else if zero == `""` {
			zero = "string"
		} else if contains(typeSliceProps, zero) {
			zero = strings.Replace(zero, "{}", "", 1)
		}
		return "make([]" + zero + ", 0)"

	case *ast.SelectorExpr:
		debugPrint("Selector type:", t.Sel.Name)
		selName := t.Sel.Name

		if selName == "Time" {
			useTime = true
			selName = "time.Time"

		}
		return selName + "{}"

	default:
		debugPrint("Unknown type:", expr)
		return "nil"
	}
}

func main() {
	fset := token.NewFileSet()
	modelsDir := ""
	if len(os.Args) > 1 {
		modelsDir = os.Args[1]
	} else {
		panic("Models directory not provided")
	}

	path := filepath.Join(modelsDir, genTypesFile+".go")
	node, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	out := &strings.Builder{}
	out.WriteString(header)
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			st, ok2 := ts.Type.(*ast.StructType)
			if !ok || !ok2 {
				continue
			}

			fields := []Field{}
			for _, f := range st.Fields.List {
				if len(f.Names) == 0 {
					continue
				}
				name := f.Names[0].Name
				zero := zeroValue(f.Type)
				/* if name == "UUID" {
					zero = `getUuid()`
				} else if name == "SchemaVersion" {
					zero = curSchemaVersion
				} else if name == "AppVersion" {
					zero = curAppVersion
				} else if name == "CreatedOn" || name == "LastModifiedOn" {
					zero = `time.Now().Format(time.RFC3339)`
				} else if name == "Status" {
					zero = `Ptr(Active)`
				} */
				for key, val := range metaZeros {
					if name == key {
						zero = val
					}
				}

				if name == "UUID" {
					debug = true
					fmt.Println("\n", name)
				}
				fields = append(fields, Field{
					Name:      name,
					ZeroValue: zero,
				})
				debug = false
			}

			data := StructData{
				Name:   ts.Name.Name,
				Fields: fields,
			}

			tmpl.Execute(out, data)
		}
	}

	imports := ""
	output := strings.Replace(out.String(), "{imports}", imports+"\n", 1)
	outPath := filepath.Join(modelsDir, fmt.Sprintf("%s_constructors.go", genTypesFile))

	if err := os.WriteFile(outPath, []byte(output), 0644); err != nil {
		fmt.Println("Error writing to file")
	}
}
