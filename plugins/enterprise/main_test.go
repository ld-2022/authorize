package main

import (
	"errors"
	"fmt"
	"github.com/ld-2022/authorize"
	"github.com/valyala/fastjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

var (
	authUrl = "http://testserver3.nnthink.com:9080"
)

func Login() (string, error) {
	v := url.Values{}
	v.Add("userName", "zxw")
	v.Add("password", "123456")
	v.Add("code", "123456")
	postForm, err := http.PostForm(authUrl+"/user/userLogin", v)
	if err != nil {
		return "", err
	}
	defer postForm.Body.Close()
	readAll, err := ioutil.ReadAll(postForm.Body)
	if err != nil {
		return "", err
	}
	parseBytes, err := fastjson.ParseBytes(readAll)
	if err != nil {
		return "", err
	}
	if string(parseBytes.GetStringBytes("code")) != "200" {
		return "", errors.New(string(readAll))
	}
	return string(parseBytes.Get("data").GetStringBytes("token")), nil
}
func TestEnterpriseAuthorize_CheckLogin(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Error(err)
		return
	}
	en := new(EnterpriseAuthorize)
	request, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Error(err)
	}
	request.Header.Add("enterpriseUserName", "zxw")
	request.Header.Add("enterpriseToken", token)
	request.Header.Add("enterpriseCode", "123456")
	checkLogin, err := en.CheckLogin(authorize.RequestParameter{Request: request})
	if err != nil {
		t.Error(err)
	}
	if !checkLogin {
		t.Error("登录失败")
	}
	fmt.Println("登录:", checkLogin)
}

func TestEnterpriseAuthorize_FindUserProjectTeamList(t *testing.T) {
	token, err := Login()
	if err != nil {
		t.Error(err)
		return
	}
	en := new(EnterpriseAuthorize)
	request, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Error(err)
	}
	request.Header.Add("token", token)
	projectTeamList, err := en.FindUserProjectTeamList(authorize.RequestParameter{Request: request})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(projectTeamList)
}

func TestEnterpriseAuthorize_FindProjectTeamList(t *testing.T) {
	en := new(EnterpriseAuthorize)
	request, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Error(err)
	}
	projectTeamList, err := en.FindProjectTeamList(authorize.RequestParameter{Request: request})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(projectTeamList)
}
