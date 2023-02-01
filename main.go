package main

import (
	"fmt"

	"github.com/excing/goflag"
)

// Config is gd2b config
type Config struct {
	Owner string `flag:"github repository owner"`
	Name  string `flag:"github repository name"`
	Token string `flag:"github authorization token (see https://docs.github.com/zh/graphql/guides/forming-calls-with-graphql)"`
}

func main() {
	var config Config
	goflag.Var(&config)
	goflag.Parse("config", "Configuration file path.")

	// 标签集合
	lables, err := getLabels(config.Owner, config.Name, config.Token)
	if err != nil {
		panic(err)
	}
	for _, lable := range lables.Nodes {
		lable.Discussions = &DiscussionPage{}
	}

	// 分类集合
	categories, err := getCategories(config.Owner, config.Name, config.Token)
	if err != nil {
		panic(err)
	}
	for _, category := range categories.Nodes {
		category.Discussions = &DiscussionPage{}
	}

	// 讨论集合
	hasNextPage := true
	endCursor := ""
	discussions := &DiscussionPage{}
	for hasNextPage {
		// 获取所有的讨论
		discussionPage, err := getDiscussionPage(config.Owner, config.Name, config.Token, endCursor)
		if err != nil {
			panic(err)
		}

		for _, discussion := range discussionPage.Nodes {
			// 获取所有的评论
			hasNextCommentPage := true
			endCommentCursor := ""
			discussion.Comments = &CommentPage{}
			for hasNextCommentPage {
				commentPage, err := getCommentPage(config.Owner, config.Name, config.Token, discussion.Number, endCommentCursor)
				if err != nil {
					panic(err)
				}

				if 0 < commentPage.TotalCount {
					discussion.Comments.Nodes = append(discussion.Comments.Nodes, commentPage.Nodes...)
					discussion.Comments.PageInfo = commentPage.PageInfo
					discussion.Comments.TotalCount += commentPage.TotalCount
				}

				// 是否有下一页评论
				hasNextCommentPage = commentPage.PageInfo.HasNextPage
				endCommentCursor = commentPage.PageInfo.EndCursor
			}

			for _, category := range categories.Nodes {
				if category.Name == discussion.Category.Name {
					category.Discussions.Nodes = append(category.Discussions.Nodes, discussion)
					category.Discussions.TotalCount++
				}
			}

			for _, discussLabel := range discussion.Labels.Nodes {
				for _, label := range lables.Nodes {
					if discussLabel.Name == label.Name {
						label.Discussions.Nodes = append(label.Discussions.Nodes, discussion)
						label.Discussions.TotalCount++
					}
				}
			}
		}

		if 0 < discussionPage.TotalCount {
			discussions.Nodes = append(discussions.Nodes, discussionPage.Nodes...)
			discussions.PageInfo = discussionPage.PageInfo
			discussions.TotalCount += discussionPage.TotalCount
		}

		// 是否有下一页
		hasNextPage = discussionPage.PageInfo.HasNextPage
		endCursor = discussionPage.PageInfo.EndCursor
	}

	for _, label := range lables.Nodes {
		fmt.Println(label.Name, label.Discussions.TotalCount)
	}

	for _, category := range categories.Nodes {
		fmt.Println(category.Name, category.Discussions.TotalCount)
	}

	for _, discussion := range discussions.Nodes {
		fmt.Println(discussion.Title, discussion.Comments.TotalCount)
	}
}
