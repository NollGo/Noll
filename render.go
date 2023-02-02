package main

import (
	"path/filepath"
	"text/template"

	"io.github.topages/assets"
)

// TemplateExecute 模板渲染接口
type TemplateExecute func(*template.Template, interface{}) error

// TemplateReader 模板文件读取接口
type TemplateReader func(name string, debug bool) (*template.Template, error)

func readTemplates(name, debugTemplateDir string, debug bool) (*template.Template, error) {
	rootTmpl := template.New(name)
	if debug {
		tmplFilesPath, err := filepath.Glob(filepath.Join(debugTemplateDir, "*.html"))
		if err != nil {
			return nil, err
		}
		return template.Must(rootTmpl.ParseFiles(tmplFilesPath...)), nil
	}
	dirEntries, err := assets.Dir.ReadDir(".")
	if err != nil {
		return nil, err
	}
	for _, entity := range dirEntries {
		if !entity.IsDir() {
			bs, err := assets.Dir.ReadFile(entity.Name())
			if err != nil {
				return nil, err
			}
			var tmplName = entity.Name()
			var tmpl *template.Template
			if name == tmplName {
				tmpl = rootTmpl
			} else {
				tmpl = rootTmpl.New(tmplName)
			}
			_, err = tmpl.Parse(string(bs))
			if err != nil {
				return nil, err
			}
		}
	}
	return rootTmpl, nil
}

func render(repo *Repository, debug bool, reader TemplateReader, execute TemplateExecute) error {
	htmlTemplate, err := reader("__ToPagesTemplate__", debug)
	if err != nil {
		return err
	}

	data := struct {
		Repo  *Repository
		Debug bool
	}{
		Repo:  repo,
		Debug: debug,
	}
	indexTemplate := htmlTemplate.Lookup("index.html")
	return execute(indexTemplate, &data)
}
