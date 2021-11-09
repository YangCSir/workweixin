package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://work.weixin.qq.com/api/doc/90001/90143/90375#视频消息

func init() {
	// 添加可解析的回调事件
	supportCallback(Video{})
}

// XML was generated 2021-11-01 20:28:46 by insomnia on Insomnia.lan.
type Video struct {
	XMLName    xml.Name `xml:"xml"`
	Text       string   `xml:",chardata"`
	ToUserName struct {
		Text string `xml:",chardata"`
	} `xml:"ToUserName"`
	FromUserName struct {
		Text string `xml:",chardata"`
	} `xml:"FromUserName"`
	CreateTime struct {
		Text string `xml:",chardata"`
	} `xml:"CreateTime"`
	MsgType struct {
		Text string `xml:",chardata"`
	} `xml:"MsgType"`
	MediaId struct {
		Text string `xml:",chardata"`
	} `xml:"MediaId"`
	ThumbMediaId struct {
		Text string `xml:",chardata"`
	} `xml:"ThumbMediaId"`
	MsgId struct {
		Text string `xml:",chardata"`
	} `xml:"MsgId"`
	AgentID struct {
		Text string `xml:",chardata"`
	} `xml:"AgentID"`
}

func (Video) GetMessageType() string {
	return "video"
}

func (Video) GetEventType() string {
	return ""
}

func (Video) GetChangeType() string {
	return ""
}

func (m Video) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (Video) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp Video
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
