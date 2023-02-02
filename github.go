package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getRepository(owner, name, token string) (*Repository, error) {
	// 标签集合
	lables, err := getLabels(owner, name, token)
	if err != nil {
		return nil, err
	}
	for _, lable := range lables.Nodes {
		lable.Discussions = &DiscussionPage{}
	}

	// 分类集合
	categories, err := getCategories(owner, name, token)
	if err != nil {
		return nil, err
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
		discussionPage, err := getDiscussionPage(owner, name, token, endCursor)
		if err != nil {
			return nil, err
		}

		for _, discussion := range discussionPage.Nodes {
			// 获取所有的评论
			hasNextCommentPage := true
			endCommentCursor := ""
			discussion.Comments = &CommentPage{}
			for hasNextCommentPage {
				commentPage, err := getCommentPage(owner, name, token, discussion.Number, endCommentCursor)
				if err != nil {
					return nil, err
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

	return &Repository{
		Labels:      lables,
		Categories:  categories,
		Discussions: discussions,
	}, nil
}

func getDiscussionPage(owner, name, token string, afterCursor string) (*DiscussionPage, error) {
	queryFormat := `{
		repository(owner: "%v", name: "%v") {
			discussions(first: 10, %v) {
				totalCount
				nodes {
					number
					title
					body
					upvoteCount
					locked
					createdAt
					updatedAt
					url
					author {
						login
						avatarUrl
						url
					}
					category {
						emoji
						name
					}
					labels(first: 10) {
						totalCount
						nodes {
							color
							name
						}
					}
					reactionGroups {
						content
						reactors(first: 1) {
							totalCount
						}
					}
				}
        pageInfo {
          hasNextPage
          endCursor
        }
			}
		}
	}`
	var result Body
	if err := query(fmt.Sprintf(queryFormat, owner, name, afterQuery(afterCursor)), token, &result); err != nil {
		return nil, err
	}
	return result.Data.Repository.Discussions, nil
}

func getCommentPage(owner, name, token string, discussionNumber int, afterCursor string) (*CommentPage, error) {
	queryFormat := `{
		repository(owner: "%v", name: "%v") {
			discussion(number: %v) {
				comments(first: 100, %v) {
					totalCount
					nodes {
						body
						createdAt
						author {
							avatarUrl
							login
							url
						}
						authorAssociation
						updatedAt
						upvoteCount
						reactionGroups {
							content
							reactors(first: 1) {
								totalCount
							}
						}
					}
					pageInfo {
						hasNextPage
						endCursor
					}
				}
			}
		}
	}`
	var result Body
	if err := query(fmt.Sprintf(queryFormat, owner, name, discussionNumber, afterQuery(afterCursor)), token, &result); err != nil {
		return nil, err
	}
	return result.Data.Repository.Discussion.Comments, nil
}

func getCategories(owner, name, token string) (*CategoryPage, error) {
	queryFormat := `{
		repository(owner: "%v", name: "%v") {
			discussionCategories(first: 100) {
				nodes {
					name
					emoji
					description
				}
				totalCount
			}
		}
	}`
	var result Body
	if err := query(fmt.Sprintf(queryFormat, owner, name), token, &result); err != nil {
		return nil, err
	}
	return result.Data.Repository.Categories, nil
}

func getLabels(owner, name, token string) (*LabelPage, error) {
	queryFormat := `{
		repository(owner: "%v", name: "%v") {
			labels(first: 100) {
				totalCount
				nodes {
					color
					name
					description
					createdAt
					updatedAt
				}
			}
		}
	}`
	var result Body
	if err := query(fmt.Sprintf(queryFormat, owner, name), token, &result); err != nil {
		return nil, err
	}
	return result.Data.Repository.Labels, nil
}

func query(body string, token string, result *Body) error {
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", strings.NewReader(queryf(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "bearer "+token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	resBodyBytes, err := ioutil.ReadAll(response.Body)

	if http.StatusOK != response.StatusCode {
		return fmt.Errorf("GraphQL query failed: %v\n%v", response.Status, string(resBodyBytes))
	}

	if err = json.Unmarshal(resBodyBytes, &result); err != nil {
		return err
	}

	if result.Data == nil {
		return fmt.Errorf("GraphQL query error: %v", string(resBodyBytes))
	}

	return nil
}

// queryf 参数的值来源 https://docs.github.com/zh/graphql/overview/explorer
func queryf(query string) string {
	query = strings.ReplaceAll(query, "\n", "")
	query = strings.ReplaceAll(query, "\t", " ")
	query = strings.ReplaceAll(query, `"`, `\"`)
	fields := strings.FieldsFunc(query, func(c rune) bool {
		return c == ' '
	})
	return fmt.Sprintf(`{"query": "query %v" }`, strings.Join(fields, " "))
}

func afterQuery(afterCursor string) string {
	after := ""
	if afterCursor != "" {
		after = fmt.Sprintf(`after: "%v"`, afterCursor)
	}
	return after
}
