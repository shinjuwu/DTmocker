package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type Callback func(w *http.ResponseWriter, r *http.Request) error

type MethodHandler struct {
	w            *http.ResponseWriter
	r            *http.Request
	callbackFunc map[string]Callback
}

var hd MethodHandler

func init() {
	fmt.Println("initlize...")
	hd.callbackFunc = map[string]Callback{}
	hd.register("loginByPwd", LoginByPwd)
	hd.register("login", Login)
}

func (hd *MethodHandler) register(id string, f Callback) {
	if _, ok := hd.callbackFunc[id]; ok {
		panic(fmt.Sprintf("function id %v: already registered", id))
	}
	hd.callbackFunc[id] = f
}

func (hd *MethodHandler) ProcessMethod() error {
	hd.r.ParseForm()
	id := ""
	for k, v := range hd.r.Form {
		switch k {
		case "method":
			id = strings.Join(v, "")
			break
		default:
			break
		}
	}
	if id == "" {
		return nil
	}
	if _, ok := hd.callbackFunc[id]; !ok {
		return errors.New(fmt.Sprintf("Callback no this function id: %v", id))
		//panic(fmt.Sprintf("Callback no this function id: %v", id))
	}
	callback := hd.callbackFunc[id]

	return callback(hd.w, hd.r)
}

func ShowApiList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Api List:")
	fmt.Fprintln(w, "Api List:")
	hd.w = &w
	hd.r = r
	err := hd.ProcessMethod()
	if err != nil {
		fmt.Fprintln(w, err)
	}
}

func LoginByPwd(w *http.ResponseWriter, r *http.Request) error {
	fmt.Println("Callback: LoginByPwd")
	fmt.Fprintln(*w, "LoginByPwd")
	return nil
}

func Login(w *http.ResponseWriter, r *http.Request) error {
	fmt.Println("Callback: login")
	fmt.Fprintln(*w, "login")
	data := map[string]interface{}{
		"id":            3225815,
		"bossCode":      "NNTI",
		"playerPrice":   955,
		"platformCode":  "NNTI_SUN_LONG8",
		"playerName":    "TEST0.23345346",
		"playerStatus":  0,
		"partitionId":   5,
		"Encode":        "sdfdgdfyrtuytjtyfgsdgdfgdfg",
		"agentCode":     "NNTI_SUN",
		"playerVersion": 92,
		"lineBets":      "|0|0.01|0.02|0.05",
		"currency":      "EUR",
	}
	usi := map[string]interface{}{
		"result": "00000",
		"data":   data,
	}
	result, err := json.Marshal(usi)
	if err != nil {
		return err
	}
	fmt.Printf("%s", result)
	fmt.Fprintln(*w, result)
	return nil
}
