package main

import (
	"fmt"
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
	categories := &CategoryPage{}
	categories.Nodes = append(categories.Nodes, &Category{Name: "Announcements"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "General"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Ideas"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Polls"})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Q&A"})
	categories.TotalCount = len(categories.Nodes)
	render(&Repository{Categories: categories})
}
