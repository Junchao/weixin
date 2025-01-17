package agent

// https://github.com/fastwego/wxwork/blob/master/corporation/apis/oauth/oauth.go
// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package oauth 网页授权登录(oauth)

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

const (
	apiAuthorize    = "https://open.weixin.qq.com/connect/oauth2/authorize"
	apiSSOAuthorize = "https://open.work.weixin.qq.com/wwopen/sso/qrConnect"
	apiUserInfo     = "/cgi-bin/user/getuserinfo"
)

type UserInfo struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	UserID   string `json:"UserId"`
	DeviceID string `json:"DeviceId"`
}

/*
构造网页授权链接
如果用户同意授权，页面将跳转至 redirect_uri/?code=CODE&state=STATE
See: https://work.weixin.qq.com/api/doc/90000/90135/91022
GET https://open.weixin.qq.com/connect/oauth2/authorize?appid=CORPID&redirect_uri=REDIRECT_URI&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect
*/
func (agent *Agent) GetAuthorizeUrl(redirectUri string, state string) (authorizeUrl string) {
	params := url.Values{}
	params.Add("appid", agent.wxwork.Config.Corpid)
	params.Add("redirect_uri", redirectUri)
	params.Add("response_type", "code")
	params.Add("scope", "snsapi_base")
	params.Add("state", state)
	return apiAuthorize + "?" + params.Encode() + "#wechat_redirect"
}

// 构造单点登录  授权链接
// 如果用户同意授权，页面将跳转至 redirect_uri/?code=CODE&state=STATE
// https://work.weixin.qq.com/api/doc/90000/90135/91019
func (agent *Agent) GetSSOAuthorizeUrl(redirectUri string, state string) (authorizeUrl string) {
	params := url.Values{}
	params.Add("appid", agent.wxwork.Config.Corpid)
	params.Add("agentid", agent.Config.AgentId)
	params.Add("redirect_uri", redirectUri)
	params.Add("state", state)
	return apiSSOAuthorize + "?" + params.Encode()
}

/*
获取访问用户身份
该接口用于根据code获取成员信息
See: https://work.weixin.qq.com/api/doc/90000/90135/91023
GET https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=ACCESS_TOKEN&code=CODE
*/
func (agent *Agent) GetUserInfo(ctx context.Context, code string) (userInfo UserInfo, err error) {
	params := url.Values{}
	params.Add("code", code)

	body, err := agent.Client.HTTPGetWithParams(ctx, apiUserInfo, params)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		err = fmt.Errorf("%s", string(body))
		return
	}

	return
}
