package main

import (
	"fmt"
	"html/template"
	"os"
	"testing"
)

func TestQueryf(t *testing.T) {
	query := queryf(`{
		repository(owner: "excing", name: "find-roots-of-word") {
			discussionCategories(first: 10) {
				nodes {
					id 
					name
					emoji
					description
				}
				totalCount
			}
		}
		viewer {
			login
		}
	}`)
	fmt.Println(query)
	// {"query": "query { repository(owner: \"excing\", name: \"find-roots-of-word\") { discussionCategories(first: 10) { nodes { id name emoji description } totalCount } } viewer { login } }" }
}

func TestRender(t *testing.T) {
	err := render(
		testRepository(),
		true,
		func(name string, debug bool) (*template.Template, error) {
			return readTemplates(name, "assets", debug)
		},
		func(t *template.Template, i interface{}) error {
			fmt.Println(t.Name())
			return t.Execute(os.Stdout, i)
		})
	if err != nil {
		t.Fatal(err)
	}
}

func testRepository() *Repository {
	labels := &LabelPage{}
	labels.Nodes = append(labels.Nodes, &Label{Name: "bug"})
	labels.TotalCount = len(labels.Nodes)

	categories := &CategoryPage{}
	categories.Nodes = append(categories.Nodes, &Category{Name: "Announcements"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "General"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Ideas"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Polls"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Q&A"})
	categories.TotalCount = len(categories.Nodes)

	discussions := &DiscussionPage{}
	discussions.Nodes = append(discussions.Nodes, &Discussion{Title: "关于模板版本的一些思考", GitHubURL: "https://github.com/ThreeTenth/GitHub-Discussions-to-Blog/discussions/8"})
	discussions.TotalCount = len(discussions.Nodes)

	return &Repository{Labels: labels, Categories: categories, Discussions: discussions}
}
