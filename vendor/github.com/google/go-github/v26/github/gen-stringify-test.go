// Copyright 2019 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// gen-stringify-test generates test methods to test the String methods.
//
// These tests eliminate most of the code coverage problems so that real
// code coverage issues can be more readily identified.
//
// It is meant to be used by go-github contributors in conjunction with the
// go generate tool before sending a PR to GitHub.
// Please see the CONTRIBUTING.md file for more information.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	ignoreFilePrefix1 = "gen-"
	ignoreFilePrefix2 = "github-"
	outputFileSuffix  = "-stringify_test.go"
)

var (
	verbose = flag.Bool("v", false, "Print verbose log messages")

	// blacklistStructMethod lists "struct.method" combos to skip.
	blacklistStructMethod = map[string]bool{}
	// blacklistStruct lists structs to skip.
	blacklistStruct = map[string]bool{
		"RateLimits": true,
	}

	funcMap = template.FuncMap{
		"isNotLast": func(index int, slice []*structField) string {
			if index+1 < len(slice) {
				return ", "
			}
			return ""
		},
		"processZeroValue": func(v string) string {
			switch v {
			case "Bool(false)":
				return "false"
			case "Float64(0.0)":
				return "0"
			case "0", "Int(0)", "Int64(0)":
				return "0"
			case `""`, `String("")`:
				return `""`
			case "Timestamp{}", "&Timestamp{}":
				return "github.Timestamp{0001-01-01 00:00:00 +0000 UTC}"
			case "nil":
				return "map[]"
			}
			log.Fatalf("Unhandled zero value: %q", v)
			return ""
		},
	}

	sourceTmpl = template.Must(template.New("source").Funcs(funcMap).Parse(source))
)

func main() {
	flag.Parse()
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, ".", sourceFilter, 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	for pkgName, pkg := range pkgs {
		t := &templateData{
			filename:     pkgName + outputFileSuffix,
			Year:         2019, // No need to change this once set (even in following years).
			Package:      pkgName,
			Imports:      map[string]string{"testing": "testing"},
			StringFuncs:  map[string]bool{},
			StructFields: map[string][]*structField{},
		}
		for filename, f := range pkg.Files {
			logf("Processing %v...", filename)
			if err := t.processAST(f); err != nil {
				log.Fatal(err)
			}
		}
		if err := t.dump(); err != nil {
			log.Fatal(err)
		}
	}
	logf("Done.")
}

func sourceFilter(fi os.FileInfo) bool {
	return !strings.HasSuffix(fi.Name(), "_test.go") &&
		!strings.HasPrefix(fi.Name(), ignoreFilePrefix1) &&
		!strings.HasPrefix(fi.Name(), ignoreFilePrefix2)
}

type templateData struct {
	filename     string
	Year         int
	Package      string
	Imports      map[string]string
	StringFuncs  map[string]bool
	StructFields map[string][]*structField
}

type structField struct {
	sortVal      string // Lower-case version of "ReceiverType.FieldName".
	ReceiverVar  string // The one-letter variable name to match the ReceiverType.
	ReceiverType string
	FieldName    string
	FieldType    string
	ZeroValue    string
	NamedStruct  bool // Getter for named struct.
}

func (t *templateData) processAST(f *ast.File) error {
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok {
			if fn.Recv != nil && len(fn.Recv.List) > 0 {
				id, ok := fn.Recv.List[0].Type.(*ast.Ident)
				if ok && fn.Name.Name == "String" {
					logf("Got FuncDecl: Name=%q, id.Name=%#v", fn.Name.Name, id.Name)
					t.StringFuncs[id.Name] = true
				} else {
					logf("Ignoring FuncDecl: Name=%q, Type=%T", fn.Name.Name, fn.Recv.List[0].Type)
				}
			} else {
				logf("Ignoring FuncDecl: Name=%q, fn=%#v", fn.Name.Name, fn)
			}
			continue
		}

		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			logf("Ignoring AST decl type %T", decl)
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			// Skip unexported identifiers.
			if !ts.Name.IsExported() {
				logf("Struct %v is unexported; skipping.", ts.Name)
				continue
			}
			// Check if the struct is blacklisted.
			if blacklistStruct[ts.Name.Name] {
				logf("Struct %v is blacklisted; skipping.", ts.Name)
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				logf("Ignoring AST type %T, Name=%q", ts.Type, ts.Name.String())
				continue
			}
			for _, field := range st.Fields.List {
				if len(field.Names) == 0 {
					continue
				}

				fieldName := field.Names[0]
				if id, ok := field.Type.(*ast.Ident); ok {
					t.addIdent(id, ts.Name.String(), fieldName.String())
					continue
				}

				if _, ok := field.Type.(*ast.MapType); ok {
					t.addMapType(ts.Name.String(), fieldName.String())
					continue
				}

				se, ok := field.Type.(*ast.StarExpr)
				if !ok {
					logf("Ignoring type %T for Name=%q, FieldName=%q", field.Type, ts.Name.String(), fieldName.String())
					continue
				}

				// Skip unexported identifiers.
				if !fieldName.IsExported() {
					logf("Field %v is unexported; skipping.", fieldName)
					continue
				}
				// Check if "struct.method" is blacklisted.
				if key := fmt.Sprintf("%v.Get%v", ts.Name, fieldName); blacklistStructMethod[key] {
					logf("Method %v is blacklisted; skipping.", key)
					continue
				}

				switch x := se.X.(type) {
				case *ast.ArrayType:
				case *ast.Ident:
					t.addIdentPtr(x, ts.Name.String(), fieldName.String())
				case *ast.MapType:
				case *ast.SelectorExpr:
				default:
					logf("processAST: type %q, field %q, unknown %T: %+v", ts.Name, fieldName, x, x)
				}
			}
		}
	}
	return nil
}

func (t *templateData) addMapType(receiverType, fieldName string) {
	t.StructFields[receiverType] = append(t.StructFields[receiverType], newStructField(receiverType, fieldName, "map[]", "nil", false))
}

func (t *templateData) addIdent(x *ast.Ident, receiverType, fieldName string) {
	var zeroValue string
	var namedStruct = false
	switch x.String() {
	case "int":
		zeroValue = "0"
	case "int64":
		zeroValue = "0"
	case "float64":
		zeroValue = "0.0"
	case "string":
		zeroValue = `""`
	case "bool":
		zeroValue = "false"
	case "Timestamp":
		zeroValue = "Timestamp{}"
	default:
		zeroValue = "nil"
		namedStruct = true
	}

	t.StructFields[receiverType] = append(t.StructFields[receiverType], newStructField(receiverType, fieldName, x.String(), zeroValue, namedStruct))
}

func (t *templateData) addIdentPtr(x *ast.Ident, receiverType, fieldName string) {
	var zeroValue string
	var namedStruct = false
	switch x.String() {
	case "int":
		zeroValue = "Int(0)"
	case "int64":
		zeroValue = "Int64(0)"
	case "float64":
		zeroValue = "Float64(0.0)"
	case "string":
		zeroValue = `String("")`
	case "bool":
		zeroValue = "Bool(false)"
	case "Timestamp":
		zeroValue = "&Timestamp{}"
	default:
		zeroValue = "nil"
		namedStruct = true
	}

	t.StructFields[receiverType] = append(t.StructFields[receiverType], newStructField(receiverType, fieldName, x.String(), zeroValue, namedStruct))
}

func (t *templateData) dump() error {
	if len(t.StructFields) == 0 {
		logf("No StructFields for %v; skipping.", t.filename)
		return nil
	}

	// Remove unused structs.
	var toDelete []string
	for k := range t.StructFields {
		if !t.StringFuncs[k] {
			toDelete = append(toDelete, k)
			continue
		}
	}
	for _, k := range toDelete {
		delete(t.StructFields, k)
	}

	var buf bytes.Buffer
	if err := sourceTmpl.Execute(&buf, t); err != nil {
		return err
	}
	clean, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("failed-to-format source:\n%v", buf.String())
		return err
	}

	logf("Writing %v...", t.filename)
	return ioutil.WriteFile(t.filename, clean, 0644)
}

func newStructField(receiverType, fieldName, fieldType, zeroValue string, namedStruct bool) *structField {
	return &structField{
		sortVal:      strings.ToLower(receiverType) + "." + strings.ToLower(fieldName),
		ReceiverVar:  strings.ToLower(receiverType[:1]),
		ReceiverType: receiverType,
		FieldName:    fieldName,
		FieldType:    fieldType,
		ZeroValue:    zeroValue,
		NamedStruct:  namedStruct,
	}
}

func logf(fmt string, args ...interface{}) {
	if *verbose {
		log.Printf(fmt, args...)
	}
}

const source = `// Copyright {{.Year}} The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-stringify-tests; DO NOT EDIT.

package {{ $package := .Package}}{{$package}}
{{with .Imports}}
import (
  {{- range . -}}
  "{{.}}"
  {{end -}}
)
{{end}}
func Float64(v float64) *float64 { return &v }
{{range $key, $value := .StructFields}}
func Test{{ $key }}_String(t *testing.T) {
  v := {{ $key }}{ {{range .}}{{if .NamedStruct}}
    {{ .FieldName }}: &{{ .FieldType }}{},{{else}}
    {{ .FieldName }}: {{.ZeroValue}},{{end}}{{end}}
  }
 	want := ` + "`" + `{{ $package }}.{{ $key }}{{ $slice := . }}{
{{- range $ind, $val := .}}{{if .NamedStruct}}{{ .FieldName }}:{{ $package }}.{{ .FieldType }}{}{{else}}{{ .FieldName }}:{{ processZeroValue .ZeroValue }}{{end}}{{ isNotLast $ind $slice }}{{end}}}` + "`" + `
	if got := v.String(); got != want {
		t.Errorf("{{ $key }}.String = %v, want %v", got, want)
	}
}
{{end}}
`
