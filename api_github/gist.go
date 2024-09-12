package api_github

import (
	"fmt"
	"github.com/team-ide/go-tool/util"
	"net/http"
	"time"
)

var GetGistsUrl = "https://api.github.com/gists"

type GetGistsResponse []*Gist

type Gist struct {
	Url         string               `json:"url"`
	ForksUrl    string               `json:"forks_url"`
	CommitsUrl  string               `json:"commits_url"`
	Id          string               `json:"id"`
	NodeId      string               `json:"node_id"`
	GitPullUrl  string               `json:"git_pull_url"`
	GitPushUrl  string               `json:"git_push_url"`
	HtmlUrl     string               `json:"html_url"`
	Files       map[string]*GistFile `json:"files"`
	Public      bool                 `json:"public"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Description string               `json:"description"`
	Comments    int                  `json:"comments"`
	CommentsUrl string               `json:"comments_url"`
	Truncated   bool                 `json:"truncated"`
}

type GistFile struct {
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Language string `json:"language"`
	RawUrl   string `json:"raw_url"`
	Size     int64  `json:"size"`
}

type GetGistsRequest struct {
	Since   string `json:"since"`    // 仅显示在给定时间后最后更新的结果。这是ISO 8601格式的时间戳：YYYY-MM-DDTHH:MM:SSZ
	PerPage int    `json:"per_page"` // 每页数量 默认 30
	Page    int    `json:"page"`     // 页码 默认 1
}

func GetGists(accessToken string, request *GetGistsRequest) (res *GetGistsResponse, err error) {
	apiUrl := GetGistsUrl
	if request != nil {
		apiUrl = fmt.Sprintf("%s?since=%s&per_page=%d&page=%d", apiUrl, request.Since, request.PerPage, request.Page)
	}
	header := http.Header{
		"Accept":               []string{"application/vnd.github+json"},
		"Authorization":        []string{"Bearer " + accessToken},
		"X-GitHub-Api-Version": []string{"2022-11-28"},
	}
	res, _, err = util.GetJsonHeader(apiUrl, header, &GetGistsResponse{})
	if err != nil {
		return
	}
	return
}

var CreateGistUrl = "https://api.github.com/gists"

type CreateGistResponse Gist

type CreateGistRequest struct {
	Description string                     `json:"description"`
	Files       map[string]*CreateGistFile `json:"files"`
	Public      bool                       `json:"public"`
}

type CreateGistFile struct {
	Content string `json:"content"`
}

func CreateGist(accessToken string, request *CreateGistRequest) (res *CreateGistResponse, err error) {
	apiUrl := CreateGistUrl
	header := http.Header{
		"Accept":               []string{"application/vnd.github+json"},
		"Authorization":        []string{"Bearer " + accessToken},
		"X-GitHub-Api-Version": []string{"2022-11-28"},
	}
	res, _, err = util.PostJsonHeader(apiUrl, header, request, &CreateGistResponse{})
	if err != nil {
		return
	}
	return
}
