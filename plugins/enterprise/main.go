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
	AuthorizeUrl = "http://testserver3.nnthink.com:9080"
)

type RespMessage struct {
	Msg  string                  `json:"msg"`
	Code string                  `json:"code"`
	Data []authorize.ProjectTeam `json:"data"`
}
type EnterpriseAuthorize struct {
}

// CheckLogin 检查用户是否登录
func (e *EnterpriseAuthorize) CheckLogin(parameter authorize.RequestParameter) (bool, error) {
	header := parameter.Request.Header
	enterpriseUserName := encoding.GetHeaderVal(header, "enterpriseUserName")
	enterpriseToken := encoding.GetHeaderVal(header, "enterpriseToken")
	enterpriseCode := encoding.GetHeaderVal(header, "enterpriseCode")
	v := url.Values{}
	v.Add("userName", enterpriseUserName)
	v.Add("code", enterpriseCode)
	v.Add("token", enterpriseToken)
	log.Println("enterpriseAuthorize ->", v.Encode())
	resp, err := http.PostForm(AuthorizeUrl+"/user/getProjectList", v)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	by, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	log.Println("enterpriseAuthorize <-", string(by))
	respJson := make(map[string]interface{})
	err = json.Unmarshal(by, &respJson)
	if err != nil {
		return false, err
	}
	if code, ok := respJson["code"]; !ok || code != "200" {
		return false, errors.New("code nil or != 200")
	}
	dataJson := respJson["data"].(map[string]interface{})
	checkMachine, checkMachineOK := dataJson["checkMachine"]
	checkToken, checkTokenOK := dataJson["checkToken"]
	if !checkMachineOK || !checkTokenOK {
		return false, errors.New("授权服务器接口返回数据不对，缺少 checkMachine | checkToken | projectListOK")
	}
	//登录账号有问题
	if !checkMachine.(bool) || !checkToken.(bool) {
		return false, errors.New("用户没有登录")
	}
	return true, nil
}

// FindUserProjectTeamList 查询用户项目团队列表
func (e *EnterpriseAuthorize) FindUserProjectTeamList(parameter authorize.RequestParameter) ([]authorize.ProjectTeam, error) {
	token := encoding.GetHeaderVal(parameter.Request.Header, "enterpriseToken")
	v := url.Values{}
	v.Add("token", token)
	postForm, err := http.PostForm(AuthorizeUrl+"/user/findUserProjectTeamList", v)
	if err != nil {
		return nil, err
	}
	defer postForm.Body.Close()
	readAll, err := ioutil.ReadAll(postForm.Body)
	if err != nil {
		return nil, err
	}
	m := new(RespMessage)
	err = json.Unmarshal(readAll, &m)
	if err != nil {
		return nil, err
	}
	if m.Code == "200" {
		return m.Data, nil
	}
	return []authorize.ProjectTeam{}, errors.New(m.Msg)
}

// FindProjectTeamList 查询项目团队列表
func (e *EnterpriseAuthorize) FindProjectTeamList(parameter authorize.RequestParameter) ([]authorize.ProjectTeam, error) {
	postForm, err := http.PostForm(AuthorizeUrl+"/user/findProjectTeamList", nil)
	if err != nil {
		return nil, err
	}
	defer postForm.Body.Close()
	readAll, err := ioutil.ReadAll(postForm.Body)
	if err != nil {
		return nil, err
	}
	m := new(RespMessage)
	err = json.Unmarshal(readAll, &m)
	if err != nil {
		return nil, err
	}
	if m.Code == "200" {
		return m.Data, nil
	}
	return []authorize.ProjectTeam{}, errors.New(m.Msg)
}

func BuildPlugin() authorize.Authorize {
	return new(EnterpriseAuthorize)
}

func BuildPluginType() authorize.PluginType {
	return authorize.Plugin_AUTH
}
