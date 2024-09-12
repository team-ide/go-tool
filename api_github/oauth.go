package api_github

import (
	"github.com/team-ide/go-tool/util"
	"net/http"
	"net/url"
)

var GetAccessTokenUrl = "https://github.com/login/oauth/access_token"

type GetAccessTokenResponse struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	Scope                 string `json:"scope"`
	TokenType             string `json:"token_type"`
	Error                 string `json:"error"`
	ErrorDescription      string `json:"error_description"`
}

func GetAccessToken(clientId string, clientSecret string, code string) (res *GetAccessTokenResponse, err error) {
	apiUrl := GetAccessTokenUrl
	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	header := http.Header{}
	header.Set("Accept", "application/json")
	res, _, err = util.PostFormHeader(apiUrl, header, data, &GetAccessTokenResponse{})
	if err != nil {
		return
	}
	return
}
