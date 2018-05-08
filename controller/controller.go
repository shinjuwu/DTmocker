package controller

import (
	"fmt"
	"net/http"
	"strings"
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
	if _, ok := hd.callbackFunc[id]; !ok {
		panic(fmt.Sprintf("Callback no this function id: %v", id))
	}
	callback := hd.callbackFunc[id]

	return callback(hd.w, hd.r)
}

func ShowApiList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Api List:")
	fmt.Fprintln(w, "Api List:")
	hd.w = &w
	hd.r = r
	hd.ProcessMethod()
}

func LoginByPwd(w *http.ResponseWriter, r *http.Request) error {
	fmt.Println("Callback: LoginByPwd")
	fmt.Fprintln(*w, "LoginByPwd")
	return nil
}
