package dingding_sdk_golang

type GetCorpTokenResp struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type GetAuthInfoResp struct {
	AuthCorpInfo    AuthCorpInfo    `json:"auth_corp_info"`
	AuthUserInfo    AuthUserInfo    `json:"auth_user_info"`
	AuthInfo        AuthInfo        `json:"auth_info"`
	ChannelAuthInfo ChannelAuthInfo `json:"channel_auth_info"`
	ErrCode         int             `json:"errcode"`
	ErrMsg          string          `json:"errmsg"`
}

type AuthCorpInfo struct {
	CorpLogoUrl     string `json:"corp_logo_url"`
	CorpName        string `json:"corp_name"`
	CorpId          string `json:"corpid"`
	Industry        string `json:"industry"`
	InviteCode      string `json:"invite_code"`
	LicenseCode     string `json:"license_code"`
	AuthChannel     string `json:"auth_channel"`
	AuthChannelType string `json:"auth_channel_type"`
	IsAuthenticated bool   `json:"is_authenticated"`
	AuthLevel       int    `json:"auth_level"`
	InviteUrl       string `json:"invite_url"`
	CorpProvince    string `json:"corp_province"`
	CorpCity        string `json:"corp_city"`
}

type AuthUserInfo struct {
	UserId string `json:"userId"`
}

type AuthInfo struct {
	Agent []Agent `json:"agent"`
}

type Agent struct {
	AgentName string   `json:"agent_name"`
	AgentId   int      `json:"agentid"`
	AppId     int      `json:"appid"`
	LogoUrl   int      `json:"logo_url"`
	AdminList []string `json:"admin_list"`
}

type ChannelAuthInfo struct {
	ChannelAgent []ChannelAgent `json:"channelAgent"`
}

type ChannelAgent struct {
	AgentName string `json:"agent_name"`
	AgentId   int    `json:"agentid"`
	AppId     int    `json:"appid"`
	LogoUrl   int    `json:"logo_url"`
}