package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://open.work.weixin.qq.com/api/doc/90001/90143/92654#删除成员事件

func init() {
	// 添加可解析的回调事件
	supportCallback(ThirdChangeContactDeleteUser{})
}

// XML was generated 2021-10-30 09:20:04 by insomnia on Insomnia.lan.
type ThirdChangeContactDeleteUser struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	SuiteId struct {
		Text string `xml:",chardata"`
	} `xml:"SuiteId"`
	AuthCorpId struct {
		Text string `xml:",chardata"`
	} `xml:"AuthCorpId"`
	InfoType struct {
		Text string `xml:",chardata"`
	} `xml:"InfoType"`
	TimeStamp struct {
		Text string `xml:",chardata"`
	} `xml:"TimeStamp"`
	ChangeType struct {
		Text string `xml:",chardata"`
	} `xml:"ChangeType"`
	UserID struct {
		Text string `xml:",chardata"`
	} `xml:"UserID"`
	OpenUserID struct {
		Text string `xml:",chardata"`
	} `xml:"OpenUserID"`
}

func (ThirdChangeContactDeleteUser) GetMessageType() string {
	return "third"
}

func (ThirdChangeContactDeleteUser) GetEventType() string {
	return "change_contact"
}

func (ThirdChangeContactDeleteUser) GetChangeType() string {
	return "delete_user"
}

func (m ThirdChangeContactDeleteUser) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeContactDeleteUser) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
	var temp ThirdChangeContactDeleteUser
	err := xml.Unmarshal(data, &temp)
	return temp, err
}
