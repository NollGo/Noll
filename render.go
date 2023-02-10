package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"io.github.topages/assets"
)

// RenderData 渲染模板的结构体
type RenderData struct {
	Viewer     *User
	Labels     *LabelPage
	Categories *CategoryPage
	Data       interface{}
}

// WriterFunc 向指定文件写入内容
type WriterFunc func(string, []byte) error

// StringWriter 可以写入到字符串的 Writer
type StringWriter struct {
	Data []byte
}

// Write 向字符串中写入
func (w *StringWriter) Write(p []byte) (n int, err error) {
	w.Data = append(w.Data, p...)
	return len(p), nil
}

func (w *StringWriter) String() string {
	return string(w.Data)
}

// FileReader 是文件读取接口
type FileReader interface {
	ReadDir(name string) ([]os.DirEntry, error)
	ReadFile(name string) ([]byte, error)
}

// LocalFileReader 本地文件读取器
type LocalFileReader struct {
	DirPath string
}

// ReadDir 读取本地文件夹
func (r *LocalFileReader) ReadDir(name string) ([]os.DirEntry, error) {
	return os.ReadDir(filepath.Join(r.DirPath, name))
}

// ReadFile 读取本地文件，并返回文件内容
func (r *LocalFileReader) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(r.DirPath, name))
}

// EmbedFileReader embed 文件读取器
type EmbedFileReader struct {
	DirEmbed embed.FS
	DirPath  string
}

// ReadDir 读取 embed 打包里的文件夹
func (r *EmbedFileReader) ReadDir(name string) ([]os.DirEntry, error) {
	return r.DirEmbed.ReadDir(filepath.Join(r.DirPath, name))
}

// ReadFile 读取 embed 文件，并返回文件内容
func (r *EmbedFileReader) ReadFile(name string) ([]byte, error) {
	return r.DirEmbed.ReadFile(filepath.Join(r.DirPath, name))
}

func render(data *GithubData, themeTmplDir string, debug bool, writer WriterFunc) error {
	// 1. 获取全局资源（assets 文件夹）文件
	readGlobalFile := func(name string) ([]byte, error) {
		var fname = filepath.Join("assets", name)
		if _, err := os.Stat(fname); err != nil {
			return assets.Dir.ReadFile(fname)
		}
		return os.ReadFile(fname)
	}

	var r FileReader
	if _, err := os.Stat(themeTmplDir); os.IsNotExist(err) {
		r = &EmbedFileReader{assets.Dir, "theme"}
	} else {
		r = &LocalFileReader{themeTmplDir}
	}

	// 2. 获取主题模板
	themeTemplate, err := readTemplates(template.New("__ToPagesTemplate__"), r, ".")
	if err != nil {
		return err
	}

	// 3. 拷贝无需渲染的主题文件到目标文件夹
	if err = copyNonRenderFiles(r, "", writer); err != nil {
		return err
	}

	// 4. 渲染模板
	htmlPages := make(map[string]string)
	stringWriter := &StringWriter{}
	indexTemplate := themeTemplate.Lookup("index.gtpl")
	_data := &RenderData{
		Viewer:     data.Viewer,
		Labels:     data.Repository.Labels,
		Categories: data.Repository.Categories,
	}
	_data.Data = data.Repository.Discussions
	if err = indexTemplate.Execute(stringWriter, _data); err != nil {
		return err
	}
	htmlPages[indexTemplate.Name()] = stringWriter.String()

	archiveTemplate := themeTemplate.Lookup("archive.gtpl")
	if err = archiveTemplate.Execute(stringWriter, _data); err != nil {
		return err
	}
	htmlPages[archiveTemplate.Name()] = stringWriter.String()

	categoriesTemplate := themeTemplate.Lookup("categories.gtpl")
	if err = categoriesTemplate.Execute(stringWriter, _data); err != nil {
		return err
	}
	htmlPages[categoriesTemplate.Name()] = stringWriter.String()

	labelsTemplate := themeTemplate.Lookup("labels.gtpl")
	if err = labelsTemplate.Execute(stringWriter, _data); err != nil {
		return err
	}
	htmlPages[labelsTemplate.Name()] = stringWriter.String()

	aboutTemplate := themeTemplate.Lookup("about.gtpl")
	if err = aboutTemplate.Execute(stringWriter, _data); err != nil {
		return err
	}
	htmlPages[aboutTemplate.Name()] = stringWriter.String()

	categoryTemplate := themeTemplate.Lookup("category.gtpl")
	for i, category := range data.Repository.Categories.Nodes {
		_data.Data = category
		if err = categoryTemplate.Execute(stringWriter, _data); err != nil {
			return err
		}
		htmlPages[fmt.Sprintf(`category/%v.gtpl`, i+1)] = stringWriter.String()
	}

	labelTemplate := themeTemplate.Lookup("label.gtpl")
	for i, label := range data.Repository.Labels.Nodes {
		_data.Data = label
		if err = labelTemplate.Execute(stringWriter, _data); err != nil {
			return err
		}
		htmlPages[fmt.Sprintf(`label/%v.gtpl`, i+1)] = stringWriter.String()
	}

	postTemplate := themeTemplate.Lookup("post.gtpl")
	for _, discussion := range data.Repository.Discussions.Nodes {
		_data.Data = discussion
		if err = postTemplate.Execute(stringWriter, _data); err != nil {
			return err
		}
		htmlPages[fmt.Sprintf(`post/%v.gtpl`, discussion.Number)] = stringWriter.String()
	}

	// 5. 全局渲染，比如调试模式
	bs, err := readGlobalFile("debug.tmpl.gtpl")
	for name, page := range htmlPages {
		// 6. 输出到目标文件夹
		if debug {
			if err = writer(name, []byte(page+"\n\n"+string(bs))); err != nil {
				return err
			}
		} else {
			if err = writer(name, []byte(page)); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyNonRenderFiles(r FileReader, name string, writer WriterFunc) error {
	entities, err := r.ReadDir(name)
	if err != nil {
		return err
	}
	for _, entity := range entities {
		fname := filepath.Join(name, entity.Name())
		if entity.IsDir() {
			err = copyNonRenderFiles(r, fname, writer)
			if err != nil {
				return err
			}
		} else if !strings.HasSuffix(fname, ".gtpl") {
			bs, err := r.ReadFile(fname)
			if err != nil {
				return err
			}
			err = writer(fname, bs)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Support syntax highlighting for Go Template files: *.go.txt, *.go.tpl, *.go.tmpl, *.gtpl.
func readTemplates(rootTmpl *template.Template, r FileReader, name string) (*template.Template, error) {
	dirEntries, err := r.ReadDir(name)
	if err != nil {
		return nil, err
	}
	for _, entity := range dirEntries {
		fname := filepath.Join(name, entity.Name())
		if entity.IsDir() {
			if _, err = readTemplates(rootTmpl, r, fname); err != nil {
				return nil, err
			}
		} else if strings.HasSuffix(fname, ".gtpl") {
			bs, err := r.ReadFile(fname)
			if err != nil {
				return nil, err
			}
			// 可能会覆盖同名的模板
			_, err = rootTmpl.New(fname).Parse(string(bs))
			if err != nil {
				return nil, err
			}
		}
	}
	return rootTmpl, nil
}
