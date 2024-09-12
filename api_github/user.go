package api_github

import (
	"github.com/team-ide/go-tool/util"
	"net/http"
)

var GetUserUrl = "https://api.github.com/user"

// GetUserResponse
/**
{
	"login": "toccata",
	"id": 1,
	"node_id": "MDQ6VXNlcjE=",
	"avatar_url": "https://github.com/images/error/octocat_happy.gif",
	"gravatar_id": "",
	"url": "https://api.github.com/users/octocat",
	"html_url": "https://github.com/octocat",
	"followers_url": "https://api.github.com/users/octocat/followers",
	"following_url": "https://api.github.com/users/octocat/following{/other_user}",
	"gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
	"starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
	"subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
	"organizations_url": "https://api.github.com/users/octocat/orgs",
	"repos_url": "https://api.github.com/users/octocat/repos",
	"events_url": "https://api.github.com/users/octocat/events{/privacy}",
	"received_events_url": "https://api.github.com/users/octocat/received_events",
	"type": "User",
	"site_admin": false,
	"name": "monalisa octocat",
	"company": "GitHub",
	"blog": "https://github.com/blog",
	"location": "San Francisco",
	"email": "octocat@github.com",
	"hireable": false,
	"bio": "There once was...",
	"twitter_username": "monatheoctocat",
	"public_repos": 2,
	"public_gists": 1,
	"followers": 20,
	"following": 0,
	"created_at": "2008-01-14T04:33:35Z",
	"updated_at": "2008-01-14T04:33:35Z",
	"private_gists": 81,
	"total_private_repos": 100,
	"owned_private_repos": 100,
	"disk_usage": 10000,
	"collaborators": 8,
	"two_factor_authentication": true,
	"plan": {
		"name": "Medium",
		"space": 400,
		"private_repos": 20,
		"collaborators": 0
	}
}
*/
type GetUserResponse struct {
	Login           string `json:"login"`
	Id              int    `json:"id"`
	Name            string `json:"name"`
	AvatarUrl       string `json:"avatar_url"`
	Email           string `json:"email"`
	HtmlUrl         string `json:"html_url"`
	Blob            string `json:"blob"`
	TwitterUsername string `json:"twitter_username"`
	Company         string `json:"company"`
	Location        string `json:"location"`
	Hireable        bool   `json:"hireable"`
	Bil             string `json:"bil"`
}

func GetUser(accessToken string) (res *GetUserResponse, err error) {
	apiUrl := GetUserUrl
	header := http.Header{
		"Accept":               []string{"application/vnd.github+json"},
		"Authorization":        []string{"Bearer " + accessToken},
		"X-GitHub-Api-Version": []string{"2022-11-28"},
	}
	res, _, err = util.GetJsonHeader(apiUrl, header, &GetUserResponse{})
	if err != nil {
		return
	}
	return
}
