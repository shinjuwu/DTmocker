package main

import (
	"DTmocker/controller"
	"DTmocker/model"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// DT bet 回傳物件
type DTResBet struct {
	Result string       `json:"result"` // 成功或失敗代碼  00000 成功 其他失敗
	Data   DTResBetData `json:"data"`   // 資料區段
}

//=========================================================================================================================================================================
// DT 回傳物件2-bet
type DTResBetData struct {
	PlatformCode string `json:"platformCode"` // 平台编码
	AgentCode    string `json:"agentCode"`    // 代理编码
	PartitionID  int    `json:"partitionId"`  // 分区编号
	BossCode     string `json:"bossCode"`     // bossCode

	ID            int64       `json:"id"`            // 玩家id
	PlayerName    string      `json:"playerName"`    // 玩家账户
	PlayerPrices  interface{} `json:"playerPrice"`   // 账户余额 (詢問是否有小數點)
	PlayerStatus  int         `json:"playerStatus"`  // 状态0-正常1-冻结
	PlayerVersion int         `json:"playerVersion"` // 玩家版本

	Encode string `json:"Encode"` // 玩家信息签名

	BetId         string `json:"betId"` // 押注ID
	IsCompetition bool   `json:"isCompetition"`

	Currency string `json:"currency"` // 货币类型
}

func main() {
	http.HandleFunc("/api/", controller.ShowApiList)
	http.HandleFunc("/login", login)
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	var res model.Response
	res.Result = "00000"
	res.Data.ID = 345345
	res.Data.BossCode = "NNTI"
	res.Data.PlayerPrice = 955
	res.Data.PlatformCode = "NNTI_SUN_LONG8"
	res.Data.PlayerName = "TEST0.23345346"
	res.Data.PlayerStatus = 0
	res.Data.PartitionID = 5
	res.Data.Encode = "sdfsdgdfg"
	res.Data.AgentCode = "NNTI_SUN"
	res.Data.PlayerVersion = 92
	res.Data.LineBets = "|0|0.01|0.02|0.05"
	res.Data.Currency = "EUR"

	result, err := json.Marshal(res)
	if err != nil {
		return
	}

	fmt.Println(string(result))
	fmt.Fprintln(w, string(result))
}
