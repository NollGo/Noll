package main

import (
	"fmt"
	"github.com/excing/goflag"
	"github.com/fsnotify/fsnotify"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Config is gd2b config
type Config struct {
	Owner    string `flag:"Github repository owner"`
	Name     string `flag:"Github repository name"`
	Token    string `flag:"Github authorization token (see https://docs.github.com/zh/graphql/guides/forming-calls-with-graphql)"`
	Pages    string `flag:"Your github pages repository name, If None, defaults to the repository where the discussion resides"`
	Debug    bool   `flag:"Debug mode if true"`
	BaseURL  string `flag:"Web site base url"`
	GamID    string `flag:"Google Analytics Measurement id, Defaults to empty to not load the Google Analytics script"`
	ThemeDir string `flag:"Filesystem path to themes directory, Defaults to embed assets/theme"`
	NewSite  bool   `flag:"Generate theme, Defaults to false"`
	Export   string `flag:"Export all Discussions to markdown, Value is the export directory"`
	Include  string `flag:"Include local files to preivew, Value is the local files directory"`
}

func main() {
	var config Config
	goflag.Var(&config)
	goflag.Parse("config", "Configuration file path.")

	if config.NewSite {
		if err := newSite(config.ThemeDir); err != nil {
			panic(err)
		}
		fmt.Println("New site success")
		return
	}
	if config.Export != "" {
		if err := export(config); err != nil {
			panic(err)
		}
		fmt.Println("Export success")
		return
	}

	fmt.Println("Start build noll siteweb")

	if config.Pages == "" {
		config.Pages = config.Name
	}

	pageDomain := fmt.Sprintf("%v.github.io", config.Owner)
	config.BaseURL = UnixPath(strings.ReplaceAll(config.BaseURL, pageDomain, "/"))

	var err error
	if _, err = os.Stat(config.Pages); os.IsNotExist(err) {
		os.MkdirAll(config.Pages, os.ModePerm)
	}

	var githubData *GithubData

	_getGithubData := func() error {
		githubData, err = getRepository(config.Owner, config.Name, config.Token, config.Include)
		return err
	}

	_render := func() error {
		return render(
			&RenderSite{
				BaseURL: config.BaseURL,
				GamID:   config.GamID,
			},
			githubData, config.ThemeDir,
			config.Debug, config.Include,
			func(s string, b []byte) error {
				fname := strings.ReplaceAll(s, ".gtpl", ".html")
				htmlPath := filepath.Join(config.Pages, fname)
				MkdirFileFolderIfNotExists(htmlPath)
				if config.Debug {
					//fmt.Println(s, string(b), "\n=========================================")
				}
				return os.WriteFile(htmlPath, b, os.ModePerm)
			})
	}

	_refreshLocalMarkdown := func(evnet fsnotify.Event) error {
		eventName := filepath.Clean(evnet.Name)
		include := filepath.Clean(config.Include)

		if !strings.HasPrefix(eventName, include) {
			return nil
		}
		if !strings.HasSuffix(eventName, ".md") {
			return nil
		}

		// map path nodes
		discussionMap := make(map[string]*Discussion)
		nodes := githubData.Repository.Discussions.Nodes
		for i := range nodes {
			discussion := nodes[i]
			if discussion.LocalPath != "" {
				discussionMap[discussion.LocalPath] = discussion
			}
		}

		if evnet.Has(fsnotify.Create) && discussionMap[eventName] == nil {
			newDis := includeLocal(eventName, githubData.Viewer, githubData.Repository.Labels, githubData.Repository.Categories, config.Token)
			appendDis(githubData, newDis)
			return nil
		}

		if evnet.Has(fsnotify.Write) {
			newDis := includeLocal(eventName, githubData.Viewer, githubData.Repository.Labels, githubData.Repository.Categories, config.Token)
			if len(newDis) <= 0 {
				return nil
			}

			if discussionMap[eventName] == nil {
				appendDis(githubData, newDis)
				return nil
			}

			discussionMap[eventName].Title = newDis[0].Title
			discussionMap[eventName].Body = newDis[0].Body
			discussionMap[eventName].BodyHTML = newDis[0].BodyHTML
			discussionMap[eventName].LocalPath = newDis[0].LocalPath
			discussionMap[eventName].CreatedAt = newDis[0].CreatedAt
			discussionMap[eventName].UpdatedAt = newDis[0].UpdatedAt
			discussionMap[eventName].Labels = newDis[0].Labels
			discussionMap[eventName].Category = newDis[0].Category
		}
		return nil
	}

	if err = _getGithubData(); err != nil {
		panic(err)
	}
	if err = _render(); err != nil {
		panic(err)
	}

	fmt.Println("Build noll siteweb finished")

	if config.Debug {
		port := ":20000"
		fs := &DirWithError{
			FS:     http.Dir(config.Pages),
			Status: map[int]string{http.StatusNotFound: "404.html"},
		}
		fmt.Println("Start noll debug mode in http://localhost" + port)

		http.Handle("/ws", debugWs(config, _render, _refreshLocalMarkdown))
		http.Handle("/", http.StripPrefix("/", http.FileServer(fs)))
		// 重新编译渲染接口
		// 调试使用
		http.Handle("/build", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			mode := query.Get("mode")
			switch mode {
			case "full":
				// 全量更新：
				// 删除本地所有文件，
				// 然后从网络上获取最新数据，
				// 再重新生成所有文件。
			case "increase":
				// 增量更新：
				// 从网络上获取最新数据，
				// 并检测本地数据是否需要更新，
				// 如果需要，则更新，否则跳过，此操作由渲染引擎处理。
				//
				// 增量更新和全量更新在流程，仅是否有删除本地所有文件的区别。
				if err = _getGithubData(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
			}
			if err = _render(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Build successed!"))
			}
		}))
		err = http.ListenAndServe(port, nil)
		if err != nil {
			panic(err)
		}
	}
}

func appendDis(githubData *GithubData, newDis []*Discussion) {
	githubData.Repository.Discussions.Nodes = append(githubData.Repository.Discussions.Nodes, newDis...)
	githubData.Repository.Discussions.TotalCount += len(newDis)
}

// DirWithError 带有错误状态页面的 http 文件系统
type DirWithError struct {
	FS     http.FileSystem
	Status map[int]string
}

// Open 返回指定名称（路径）的文件
func (d *DirWithError) Open(name string) (http.File, error) {
	f, err := d.FS.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			_404, ok := d.Status[http.StatusNotFound]
			if ok {
				return d.FS.Open(_404)
			}
		} else if os.IsPermission(err) {
			_403, ok := d.Status[http.StatusForbidden]
			if ok {
				return d.FS.Open(_403)
			}
		} else {
			// Default:
			_500, ok := d.Status[http.StatusInternalServerError]
			if ok {
				return d.FS.Open(_500)
			}
		}
		return f, err
	}

	return f, nil
}
