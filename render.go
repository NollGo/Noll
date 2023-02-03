package main

import (
	"fmt"
	"path/filepath"
	"text/template"

	"io.github.topages/assets"
)

// RenderData 渲染模板的结构体
type RenderData struct {
	Debug bool
	Data  interface{}
}

// TemplateExecute 模板渲染接口
type TemplateExecute func(string, *template.Template, interface{}) error

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

	indexTemplate := htmlTemplate.Lookup("index.html")
	if err = execute(indexTemplate.Name(), indexTemplate, &RenderData{true, &repo}); err != nil {
		return err
	}

	postTemplate := htmlTemplate.Lookup("post.html")
	for _, discussion := range repo.Discussions.Nodes {
		if err = execute(fmt.Sprintf(`p/%v.html`, discussion.Number), postTemplate, &RenderData{true, &discussion}); err != nil {
			return err
		}
	}

	return nil
}
