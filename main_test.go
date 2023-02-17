package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gosimple/slug"
)

func TestURLSlug(t *testing.T) {
	text := slug.Make("Hellö Wörld хелло ворлд")
	fmt.Println(text) // Will print: "hello-world-khello-vorld"

	text2 := slug.Make("o(*￣▽￣*)ブ")
	fmt.Println(text2) // Will print: "o-v-bu"

	text3 := slug.Make("💡 (_　_)。゜zｚＺ")
	fmt.Println(text3) // Will print: "zzz"

	text4 := slug.Make("( ´･･)ﾉ(._.`)")
	fmt.Println(text4) // Will print: "no"

	text5 := slug.Make("ヾ(⌐■_■)ノ♪")
	fmt.Println(text5) // Will print: "no"

	someText := slug.Make("影師")
	fmt.Println(someText) // Will print: "ying-shi"

	enText := slug.MakeLang("This & that", "en")
	fmt.Println(enText) // Will print: "this-and-that"

	deText := slug.MakeLang("Diese & Dass", "de")
	fmt.Println(deText) // Will print: "diese-und-dass"

	slug.Lowercase = false // Keep uppercase characters
	deUppercaseText := slug.MakeLang("Diese & Dass", "fa")
	fmt.Println(deUppercaseText) // Will print: "Diese-und-Dass"

	slug.CustomSub = map[string]string{
		"water": "sand",
	}
	textSub := slug.Make("water is hot")
	fmt.Println(textSub) // Will print: "sand-is-hot"
}

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
		&RenderSite{"/"},
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
