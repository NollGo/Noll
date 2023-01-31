package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getDiscussions(owner, name, token string) (*DiscussionPage, error) {
	queryFormat := `{
		repository(owner: "excing", name: "find-roots-of-word") {
			discussions(first: 10) {
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
					reactions(first: 100) {
						totalCount
						nodes {
							content
							user {
								login
								avatarUrl
								url
							}
						}
					}
					labels(first: 10) {
						totalCount
						nodes {
							color
							name
						}
					}
					comments(first: 100) {
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
							reactions(first: 100) {
								totalCount
								nodes {
									content
									createdAt
									user {
										login
										avatarUrl
										url
									}
								}
							}
						}
					}
				}
			}
		}
	}`
	var result Body
	if err := query(fmt.Sprintf(queryFormat, owner, name), token, &result); err != nil {
		return nil, err
	}
	return result.Data.Repository.Discussions, nil
}

func getCategories(owner, name, token string) (*CategoryPage, error) {
	queryFormat := `{
		repository(owner: "%v", name: "%v") {
			discussionCategories(first: 10) {
				nodes {
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
	}`
	var result Body
	if err := query(fmt.Sprintf(queryFormat, owner, name), token, &result); err != nil {
		return nil, err
	}
	return result.Data.Repository.DiscussionCategories, nil
}

func query(body string, token string, result interface{}) error {
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
