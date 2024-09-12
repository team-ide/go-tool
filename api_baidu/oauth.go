package api_baidu

import "github.com/team-ide/go-tool/util"

var GetAccessTokenUrl = "https://aip.baidubce.com/oauth/2.0/token"

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetAccessToken(apiKey string, secretKey string) (res *GetAccessTokenResponse, err error) {
	apiUrl := GetAccessTokenUrl + "?grant_type=client_credentials&client_id=" + apiKey + "&client_secret=" + secretKey
	res, _, err = util.GetJson(apiUrl, &GetAccessTokenResponse{})
	if err != nil {
		return
	}
	return
}
