package model

type UserInfo struct {
	ID            int64  `json:"id"`
	BossCode      string `json:"bossCode"`
	PlayerPrice   int64  `json:"playerPrice"`  //帳戶餘額
	PlatformCode  string `json:"platformCode"` //平台編碼
	PlayerName    string `json:"playerName"`   //玩家帳戶
	PlayerStatus  int    `json:"playerStatus"` //狀態  0-正常  1-凍結
	PartitionID   int    `json:"partitionId"`  //分區編號
	Encode        string `json:"Encode"`       //玩家信息簽名
	AgentCode     string `json:"agentCode"`    //代理編碼
	PlayerVersion int    `json:"playerVersion"`
	LineBets      string `json:"lineBets"` //線注信息
	Currency      string `json:"currency"` //貨幣類型
}

type Response struct {
	Result string   `json:"result"`
	Data   UserInfo `json:"data"`
}
