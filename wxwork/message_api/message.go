package message_api

// https://github.com/fastwego/wxwork/tree/master/corporation/apis/message
// Copyright 2021 FastWeGo
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

// Package message 消息推送

import (
	"bytes"
	"context"
	"net/url"

	"github.com/lixinio/weixin/utils"
	"github.com/lixinio/weixin/wxwork/agent"
)

const (
	apiSend                  = "/cgi-bin/message/send"
	apiUpdateTaskcard        = "/cgi-bin/message/update_taskcard"
	apiAppchatCreate         = "/cgi-bin/appchat/create"
	apiAppchatUpdate         = "/cgi-bin/appchat/update"
	apiAppchatGet            = "/cgi-bin/appchat/get"
	apiAppchatSend           = "/cgi-bin/appchat/send"
	apiLinkedcorpMessageSend = "/cgi-bin/linkedcorp/message/send"
	apiGetStatistics         = "/cgi-bin/message/get_statistics"
)

type MessageApi struct {
	*utils.Client
}

func NewAgentApi(agent *agent.Agent) *MessageApi {
	return &MessageApi{
		Client: agent.Client,
	}
}

/*
发送应用消息
应用支持推送文本、图片、视频、文件、图文等类型。
See: https://work.weixin.qq.com/api/doc/90000/90135/90236
POST https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN
*/
func (api *MessageApi) Send(ctx context.Context, payload []byte) (resp []byte, err error) {
	return api.Client.HTTPPost(ctx, apiSend, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新任务卡片消息状态
应用可以发送任务卡片消息，发送之后可再通过接口更新用户任务卡片消息的选择状态。
See: https://work.weixin.qq.com/api/doc/90000/90135/91579
POST https://qyapi.weixin.qq.com/cgi-bin/message/update_taskcard?access_token=ACCESS_TOKEN
*/
func (api *MessageApi) UpdateTaskcard(ctx context.Context, payload []byte) (resp []byte, err error) {
	return api.Client.HTTPPost(ctx, apiUpdateTaskcard, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
创建群聊会话
See: https://work.weixin.qq.com/api/doc/90000/90135/90245
POST https://qyapi.weixin.qq.com/cgi-bin/appchat/create?access_token=ACCESS_TOKEN
*/
func (api *MessageApi) AppchatCreate(ctx context.Context, payload []byte) (resp []byte, err error) {
	return api.Client.HTTPPost(ctx, apiAppchatCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改群聊会话
See: https://work.weixin.qq.com/api/doc/90000/90135/90246
POST https://qyapi.weixin.qq.com/cgi-bin/appchat/update?access_token=ACCESS_TOKEN
*/
func (api *MessageApi) AppchatUpdate(ctx context.Context, payload []byte) (resp []byte, err error) {
	return api.Client.HTTPPost(ctx, apiAppchatUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取群聊会话
See: https://work.weixin.qq.com/api/doc/90000/90135/90247
GET https://qyapi.weixin.qq.com/cgi-bin/appchat/get?access_token=ACCESS_TOKEN&chatid=CHATID
*/
func (api *MessageApi) AppchatGet(ctx context.Context, params url.Values) (resp []byte, err error) {
	return api.Client.HTTPGet(ctx, apiAppchatGet+"?"+params.Encode())
}

/*
应用推送消息
应用支持推送文本、图片、视频、文件、图文等类型。
See: https://work.weixin.qq.com/api/doc/90000/90135/90248
POST https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=ACCESS_TOKEN
*/
func (api *MessageApi) AppchatSend(ctx context.Context, payload []byte) (resp []byte, err error) {
	return api.Client.HTTPPost(ctx, apiAppchatSend, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
互联企业消息推送
互联企业的应用支持推送文本、图片、视频、文件、图文等类型。
See: https://work.weixin.qq.com/api/doc/90000/90135/90250
POST https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=ACCESS_TOKEN
*/
func (api *MessageApi) LinkedcorpMessageSend(ctx context.Context, payload []byte) (resp []byte, err error) {
	return api.Client.HTTPPost(ctx, apiLinkedcorpMessageSend, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询应用消息发送统计
See: https://work.weixin.qq.com/api/doc/90000/90135/92369
POST https://qyapi.weixin.qq.com/cgi-bin/message/get_statistics?access_token=ACCESS_TOKEN
*/
func (api *MessageApi) GetStatistics(ctx context.Context, payload []byte) (resp []byte, err error) {
	return api.Client.HTTPPost(ctx, apiGetStatistics, bytes.NewReader(payload), "application/json;charset=utf-8")
}
