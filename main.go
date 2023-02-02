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

	if _, err := os.Stat(config.Pages); os.IsNotExist(err) {
		os.MkdirAll(config.Pages, os.ModePerm)
	}

	repository, err := getRepository(config.Owner, config.Name, config.Token)
	if err != nil {
		panic(err)
	}

	_render := func() error {
		return render(repository, config.Debug, func(name string, debug bool) (*template.Template, error) {
			return readTemplates(name, "assets", debug)
		}, func(t *template.Template, i interface{}) error {
			dist, err := os.Create(filepath.Join(config.Pages, t.Name()))
			if err != nil {
				return err
			}
			return t.Execute(dist, i)
		})
	}
	if err = _render(); err != nil {
		panic(err)
	}

	if config.Debug {
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(config.Pages))))
		// 重新编译渲染接口
		// 调试使用
		http.Handle("/build", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("Start toPages finished")
}
