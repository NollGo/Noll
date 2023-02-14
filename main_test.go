package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestStringContains(t *testing.T) {
	if strings.Contains(`{
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
	}`, `first: 11`) {
		t.Log()
	} else {
		t.Fail()
	}
}

func TestTimeConvert(t *testing.T) {
	timeStr := ""
	time, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		// Result: failed
		t.Fatal(err)
	}
	t.Log(time.String())
}

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
		"assets/theme",
		true,
		func(s string, b []byte) error {
			fmt.Println(s)
			_, err := os.Stdout.Write(b)
			return err
		},
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOsStat(t *testing.T) {
	if _, err := os.Stat(""); os.IsNotExist(err) {
		t.Fatal(err)
	} else {
		t.Log("PASS")
	}
}

func TestGetEmoji4GEmoji(t *testing.T) {
	gemojiFormat := `<div></div>`
	gemoji := `<g-emoji class="g-emoji" alias="mega" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f4e3.png">📣</g-emoji>`
	t.Log(getGemoji(fmt.Sprintf(gemojiFormat, gemoji)) == gemoji)
}

func TestPref(t *testing.T) {
	str := "example.txt"
	suffix := "txt"
	t.Log(strings.HasSuffix(str, suffix))
}

func testRepository() *GithubData {
	labels := &LabelPage{}
	labels.Nodes = append(labels.Nodes, &Label{Name: "bug", Discussions: &DiscussionPage{}})
	labels.TotalCount = len(labels.Nodes)

	categories := &CategoryPage{}
	categories.Nodes = append(categories.Nodes, &Category{Name: "Announcements", Discussions: &DiscussionPage{}})
	categories.Nodes = append(categories.Nodes, &Category{Name: "General", Discussions: &DiscussionPage{}})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Ideas", Discussions: &DiscussionPage{}})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Polls", Discussions: &DiscussionPage{}})
	categories.Nodes = append(categories.Nodes, &Category{Name: "Q&A", Discussions: &DiscussionPage{}})
	categories.TotalCount = len(categories.Nodes)

	discussions := &DiscussionPage{}
	discussions.Nodes = append(discussions.Nodes, &Discussion{Title: "关于模板版本的一些思考", GitHubURL: "https://github.com/ThreeTenth/GitHub-Discussions-to-Blog/discussions/8", Category: &Category{Name: "Ideas"}, Comments: &CommentPage{}})
	discussions.TotalCount = len(discussions.Nodes)

	return &GithubData{
		&Repository{Labels: labels, Categories: categories, Discussions: discussions},
		&User{Login: "excing"},
	}
}
