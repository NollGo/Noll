package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const exportTemplate = `---
title: {{ .Title }}
createAt: {{ .CreatedAt }}
updateAt: {{ .UpdatedAt }}
tags: {{ .LabelsString }}
categories: {{ .Category }}
---

{{ .Body }}
`

func export(config Config) (err error) {
	var data *GithubData
	if data, err = getRepository(config.Owner, config.Name, config.Token); err != nil {
		return err
	}

	discussions := data.Repository.Discussions
	if discussions == nil || discussions.TotalCount == 0 {
		return fmt.Errorf("no discussions")
	}

	discussionsMap := groupByCategory(discussions.Nodes)
	tmpl, err := parseTemplate(exportTemplate)
	if err != nil {
		return err
	}
	for category, discussions := range discussionsMap {
		_ = os.MkdirAll(filepath.Join(config.ExportDir, category), os.ModePerm)
		for _, discussion := range discussions {
			if err = exportDiscussion(config.ExportDir, discussion, tmpl); err != nil {
				return err
			}
		}
	}
	return err
}

func exportDiscussion(dir string, discussion *Discussion, tmpl *template.Template) error {
	filePath := filepath.Join(dir, discussion.Category.Name, fmt.Sprintf("%v.md", discussion.Number))
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, map[string]interface{}{
		"Title":        discussion.Title,
		"CreatedAt":    discussion.CreatedAt.Format("2006-01-02 15:04:05"),
		"UpdatedAt":    discussion.UpdatedAt.Format("2006-01-02 15:04:05"),
		"LabelsString": discussion.Labels.String(),
		"Labels":       discussion.Labels,
		"Category":     discussion.Category.Name,
		"Body":         discussion.Body,
	})
	if err != nil {
		return err
	}

	return nil
}

func groupByCategory(discussions []*Discussion) map[string][]*Discussion {
	group := make(map[string][]*Discussion)
	for i := range discussions {
		discussion := discussions[i]
		group[discussion.Category.Name] = append(group[discussion.Category.Name], discussion)
	}
	return group
}

func parseTemplate(tmp string) (*template.Template, error) {
	return template.New("export").Parse(tmp)
}
