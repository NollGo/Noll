package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func includeLocal(path string, viewer *User, lables *LabelPage, categories *CategoryPage, token string) []*Discussion {
	// labels node to map
	labelsMap := make(map[string]*Label)
	for _, label := range lables.Nodes {
		labelsMap[label.Name] = label
	}

	// categories node to map
	categoriesMap := make(map[string]*Category)
	for _, category := range categories.Nodes {
		categoriesMap[category.Name] = category
	}

	discussions := make([]*Discussion, 0)
	// read local file to discussions
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if strings.HasSuffix(info.Name(), ".md") {
				discussion, err := fakeDiscussion(path, viewer, categoriesMap, token, labelsMap)
				if err != nil {
					return err
				}
				if discussion != nil {
					discussions = append(discussions, discussion)
				}

			}
		}
		return err
	})
	if err != nil {
		panic(err)
	}

	return discussions
}

func fakeDiscussion(path string, viewer *User, categoriesMap map[string]*Category, token string, labelsMap map[string]*Label) (*Discussion, error) {
	md, err := parseMarkdownFile(path)
	if err != nil {
		return nil, err
	}
	if md == nil {
		return nil, nil
	}

	discussion := &Discussion{}
	discussion.Number = rand.Int()
	discussion.Title = md.Metadata.Title
	discussion.CreatedAt = parseTime(md.Metadata.CreateAt)
	discussion.UpdatedAt = parseTime(md.Metadata.UpdateAt)
	discussion.Body = md.Body
	discussion.Author = viewer
	discussion.Category = categoriesMap[md.Metadata.Category]
	discussion.LocalPath = path

	markdownHtml, err := renderMarkdown(md.Body, token)
	if err != nil {
		return nil, err
	}
	discussion.BodyHTML = markdownHtml
	discussion.Labels = &LabelPage{
		Nodes: make([]*Label, 0),
	}
	for _, labelName := range md.Metadata.Tags {
		discussion.Labels.Nodes = append(discussion.Labels.Nodes, labelsMap[labelName])
	}
	discussion.Labels.TotalCount = len(discussion.Labels.Nodes)
	discussion.Comments = &CommentPage{Nodes: make([]*Comment, 0)}
	discussion.ReactionGroups = make([]*ReactionGroup, 0)

	if discussion.Category == nil {
		return nil, fmt.Errorf("category not found: %s", md.Metadata.Category)
	}

	return discussion, nil
}

func parseTime(timeStr string) time.Time {
	t, err := time.Parse(time.DateTime, timeStr)
	if err != nil {
		t = time.Now()
	}
	return t
}

type LocalMeta struct {
	Title    string   `yaml:"title"`
	CreateAt string   `yaml:"createAt"`
	UpdateAt string   `yaml:"updateAt"`
	Tags     []string `yaml:"tags"`
	Category string   `yaml:"category"`
}

type LocalMarkdown struct {
	Metadata LocalMeta
	Body     string
}

func parseMarkdownFile(filePath string) (*LocalMarkdown, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	sections := strings.SplitN(string(content), "---", 3)
	if len(sections) < 3 {
		fmt.Printf("markdown file %s format error, skip it\n", filePath)
		return nil, nil
	}

	var metadata LocalMeta
	err = yaml.Unmarshal([]byte(sections[1]), &metadata)
	if err != nil {
		return nil, err
	}

	md := LocalMarkdown{
		Metadata: metadata,
		Body:     sections[2],
	}

	return &md, nil
}
