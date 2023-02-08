package main

import (
	"fmt"
	"path/filepath"
	"text/template"

	"io.github.topages/assets"
)

// RenderData 渲染模板的结构体
type RenderData struct {
	Debug      bool
	Viewer     *User
	Labels     *LabelPage
	Categories *CategoryPage
	Data       interface{}
}

// TemplateExecute 模板渲染接口
type TemplateExecute func(string, *template.Template, interface{}) error

// TemplateReader 模板文件读取接口
type TemplateReader func(name string, debug bool) (*template.Template, error)

// Support syntax highlighting for Go Template files: *.go.txt, *.go.tpl, *.go.tmpl, *.gtpl.
func readTemplates(name, templateDir string, debug bool) (*template.Template, error) {
	rootTmpl := template.New(name)
	if debug && templateDir != "" {
		tmplFilesPath, err := filepath.Glob(filepath.Join(templateDir, "*.gtpl"))
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

func render(data *GithubData, debug bool, reader TemplateReader, execute TemplateExecute) error {
	htmlTemplate, err := reader("__ToPagesTemplate__", debug)
	if err != nil {
		return err
	}

	indexTemplate := htmlTemplate.Lookup("index.gtpl")
	if err = execute(indexTemplate.Name(), indexTemplate, &RenderData{true, data.Viewer, data.Repository.Labels, data.Repository.Categories, data.Repository.Discussions}); err != nil {
		return err
	}

	archiveTemplate := htmlTemplate.Lookup("archive.gtpl")
	if err = execute(archiveTemplate.Name(), archiveTemplate, &RenderData{true, data.Viewer, data.Repository.Labels, data.Repository.Categories, data.Repository.Discussions}); err != nil {
		return err
	}

	categoriesTemplate := htmlTemplate.Lookup("categories.gtpl")
	if err = execute(categoriesTemplate.Name(), categoriesTemplate, &RenderData{true, data.Viewer, data.Repository.Labels, data.Repository.Categories, nil}); err != nil {
		return err
	}

	categoryTemplate := htmlTemplate.Lookup("category.gtpl")
	for i, category := range data.Repository.Categories.Nodes {
		if err = execute(fmt.Sprintf(`category/%2v.gtpl`, i+1), categoryTemplate, &RenderData{true, data.Viewer, data.Repository.Labels, data.Repository.Categories, category}); err != nil {
			return err
		}
	}

	labelsTemplate := htmlTemplate.Lookup("labels.gtpl")
	if err = execute(labelsTemplate.Name(), labelsTemplate, &RenderData{true, data.Viewer, data.Repository.Labels, data.Repository.Categories, nil}); err != nil {
		return err
	}

	labelTemplate := htmlTemplate.Lookup("label.gtpl")
	for i, label := range data.Repository.Labels.Nodes {
		if err = execute(fmt.Sprintf(`label/%2v.gtpl`, i+1), labelTemplate, &RenderData{true, data.Viewer, data.Repository.Labels, data.Repository.Categories, label}); err != nil {
			return err
		}
	}

	aboutTemplate := htmlTemplate.Lookup("about.gtpl")
	if err = execute(aboutTemplate.Name(), aboutTemplate, &RenderData{true, data.Viewer, data.Repository.Labels, data.Repository.Categories, data.Repository.Discussions}); err != nil {
		return err
	}

	postTemplate := htmlTemplate.Lookup("post.gtpl")
	for _, discussion := range data.Repository.Discussions.Nodes {
		if err = execute(fmt.Sprintf(`p/%v.gtpl`, discussion.Number), postTemplate, &RenderData{true, data.Viewer, data.Repository.Labels, data.Repository.Categories, discussion}); err != nil {
			return err
		}
	}

	return nil
}
