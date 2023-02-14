package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"io.github.topages/assets"
)

// RenderData 渲染模板的结构体
type RenderData struct {
	Viewer     *User
	Labels     *LabelPage
	Categories *CategoryPage
	Data       interface{}
}

// JsRenderLoader js 渲染加载器
// 包含数学公式、图表、地图和三维模型
type JsRenderLoader struct {
	HTML       string
	HasMermaid bool
	HasMathjax bool
	HasGeojson bool
	HasSTL3D   bool
}

// Has 返回 Html 中是否包含需要 js 渲染的内容
func (l *JsRenderLoader) Has() bool {
	if strings.Contains(l.HTML, `data-type="geojsin"`) || strings.Contains(l.HTML, `data-type="topojson"`) {
		l.HasGeojson = true
	}
	if strings.Contains(l.HTML, `</math-renderer>`) {
		l.HasMathjax = true
	}
	if strings.Contains(l.HTML, `data-type="mermaid"`) {
		l.HasMermaid = true
	}
	if strings.Contains(l.HTML, `data-type="stl"`) {
		l.HasSTL3D = true
	}
	return l.HasGeojson || l.HasMathjax || l.HasMermaid || l.HasSTL3D
}

// WriterFunc 向指定文件写入内容
type WriterFunc func(string, []byte) error

// StringWriter 可以写入到字符串的 Writer
type StringWriter struct {
	Data []byte
}

// Reset 重置资源
func (w *StringWriter) Reset() *StringWriter {
	w.Data = make([]byte, 0)
	return w
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
	templateFuncMap := template.FuncMap{
		"time": func() time.Time { return time.Time{} },
		"isd": func(d1, d2 time.Time) bool {
			return d1.Year() == d2.Year() && d1.YearDay() == d2.YearDay()
		},
		"ism": func(d1, d2 time.Time) bool {
			return d1.Year() == d2.Year() && d1.Month() == d2.Month()
		},
		"isy": func(d1, d2 time.Time) bool {
			return d1.Year() == d2.Year()
		},
	}
	themeTemplate, err := readTemplates(
		template.New("__ToPagesTemplate__").Funcs(templateFuncMap), r, ".")
	if err != nil {
		return err
	}

	jsRenderLoaderBS, err := readGlobalFile("js-render-loader.gtpl")
	if err != nil {
		return err
	}
	jsRenderTemplate, err := template.New("jsRenderLoader").Parse(string(jsRenderLoaderBS))
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
	if err = indexTemplate.Execute(stringWriter.Reset(), _data); err != nil {
		return err
	}
	htmlPages[indexTemplate.Name()] = stringWriter.String()

	categoriesTemplate := themeTemplate.Lookup("categories.gtpl")
	if err = categoriesTemplate.Execute(stringWriter.Reset(), _data); err != nil {
		return err
	}
	htmlPages[categoriesTemplate.Name()] = stringWriter.String()

	labelsTemplate := themeTemplate.Lookup("labels.gtpl")
	if err = labelsTemplate.Execute(stringWriter.Reset(), _data); err != nil {
		return err
	}
	htmlPages[labelsTemplate.Name()] = stringWriter.String()

	aboutTemplate := themeTemplate.Lookup("about.gtpl")
	if err = aboutTemplate.Execute(stringWriter.Reset(), _data); err != nil {
		return err
	}
	htmlPages[aboutTemplate.Name()] = stringWriter.String()

	categoryTemplate := themeTemplate.Lookup("category.gtpl")
	for _, category := range data.Repository.Categories.Nodes {
		_data.Data = category
		if err = categoryTemplate.Execute(stringWriter.Reset(), _data); err != nil {
			return err
		}
		htmlPages[fmt.Sprintf(`category/%v.gtpl`, category.Name)] = stringWriter.String()
	}

	labelTemplate := themeTemplate.Lookup("label.gtpl")
	for _, label := range data.Repository.Labels.Nodes {
		_data.Data = label
		if err = labelTemplate.Execute(stringWriter.Reset(), _data); err != nil {
			return err
		}
		htmlPages[fmt.Sprintf(`label/%v.gtpl`, label.Name)] = stringWriter.String()
	}

	postTemplate := themeTemplate.Lookup("post.gtpl")
	for _, discussion := range data.Repository.Discussions.Nodes {
		_data.Data = discussion
		if err = postTemplate.Execute(stringWriter.Reset(), _data); err != nil {
			return err
		}
		jrl := &JsRenderLoader{HTML: stringWriter.String()}
		if jrl.Has() {
			jsRenderTemplate.Execute(stringWriter, jrl)
		}
		htmlPages[fmt.Sprintf(`post/%v.gtpl`, discussion.Number)] = stringWriter.String()
	}

	archiveTemplate := themeTemplate.Lookup("archive.gtpl")
	totalCount := data.Repository.Discussions.TotalCount
	pageIndex := 1 // 编号从 1 开始
	pageSize := 30
	pageCount := totalCount / pageSize
	if totalCount%pageSize > 0 {
		pageCount++
	}
	for start := 0; start < totalCount; {
		end := start + pageSize
		if end > totalCount {
			end = totalCount
		}
		nodes := data.Repository.Discussions.Nodes[start:end]
		_pageInfo := &PageInfo{end < totalCount, fmt.Sprintf("%v", pageIndex+1), 0 < start, fmt.Sprintf("%v", pageIndex-1)}
		_data.Data = &DiscussionPage{end - start, nodes, _pageInfo}
		if err = archiveTemplate.Execute(stringWriter.Reset(), _data); err != nil {
			return err
		}
		htmlPages[fmt.Sprintf("archive/%v.gtpl", pageIndex)] = stringWriter.String()
		pageIndex++
		start = end
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
