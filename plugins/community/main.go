package main

import (
	"encoding/json"
	"errors"
	"github.com/ld-2022/authorize"
	"github.com/ld-2022/authorize/encoding"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	AuthorizeUrl string
)

type CommunityAuthorize struct {
}

// CheckLogin 检查用户是否登录
func (c *CommunityAuthorize) CheckLogin(parameter authorize.RequestParameter) (bool, error) {
	log.Println("社区版登录信息校验")
	req := parameter.Request
	communityUsername := encoding.GetHeaderVal(req.Header, "communityUsername")
	communityToken := encoding.GetHeaderVal(req.Header, "communityToken")
	if communityUsername == "" || communityToken == "" {
		return false, errors.New("用户名和token不能为空")
	}

	v := url.Values{}
	v.Add("token", communityToken)
	postForm, err := http.PostForm(AuthorizeUrl, v)
	if err != nil {
		log.Println("授权服务异常 Err:", err)
		return false, errors.New("授权服务异常")
	}
	defer postForm.Body.Close()
	all, err := ioutil.ReadAll(postForm.Body)
	if err != nil {
		log.Println("读取授权数据异常 Err:", err)
		return false, errors.New("读取授权数据异常")
	}
	m := new(AuthMsg)
	err = json.Unmarshal(all, m)
	if err != nil {
		log.Println("解析授权返回数据 Err:", err, " Data:", string(all))
		return false, errors.New("解析授权返回数据异常")
	}
	if m.Status != 2 {
		return false, errors.New("还未申请通过")
	}
	return true, nil
}

// FindUserProjectTeamList 查询用户项目团队列表
func (c *CommunityAuthorize) FindUserProjectTeamList(parameter authorize.RequestParameter) ([]authorize.ProjectTeam, error) {
	log.Println("社区版查询用户项目团队列表")
	return []authorize.ProjectTeam{}, nil
}

func BuildPlugin() authorize.Authorize {
	return new(CommunityAuthorize)
}

type AuthMsg struct {
	Errormessage string `json:"errormessage"`
	Errorcode    string `json:"errorcode"`
	Status       int    `json:"status"`
}
