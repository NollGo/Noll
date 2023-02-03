package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/excing/goflag"
)

// Config is gd2b config
type Config struct {
	Owner string `flag:"Github repository owner"`
	Name  string `flag:"Github repository name"`
	Token string `flag:"Github authorization token (see https://docs.github.com/zh/graphql/guides/forming-calls-with-graphql)"`
	Pages string `flag:"Your github pages repository name, If None, defaults to the repository where the discussion resides"`
	Debug bool   `flag:"Debug mode if true"`
}

func main() {
	var config Config
	goflag.Var(&config)
	goflag.Parse("config", "Configuration file path.")

	if config.Pages == "" {
		config.Pages = config.Name
	}

	var err error
	if _, err = os.Stat(config.Pages); os.IsNotExist(err) {
		os.MkdirAll(config.Pages, os.ModePerm)
	}

	var repository *Repository

	_getRepository := func() error {
		repository, err = getRepository(config.Owner, config.Name, config.Token)
		return err
	}

	_render := func() error {
		return render(repository, config.Debug, func(name string, debug bool) (*template.Template, error) {
			return readTemplates(name, "assets", debug)
		}, func(s string, t *template.Template, i interface{}) error {
			htmlPath := filepath.Join(config.Pages, s)
			MkdirFileFolderIfNotExists(htmlPath)
			dist, err := os.Create(htmlPath)
			if err != nil {
				return err
			}
			return t.Execute(dist, i)
		})
	}
	if err = _getRepository(); err != nil {
		panic(err)
	}
	if err = _render(); err != nil {
		panic(err)
	}

	fmt.Println("Start toPages package finished")

	if config.Debug {
		fmt.Println("Start toPages debug mode")
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(config.Pages))))
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
				if err = _getRepository(); err != nil {
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
		http.ListenAndServe(":20000", nil)
	}
}
