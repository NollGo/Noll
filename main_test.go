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
