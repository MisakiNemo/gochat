package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Rows  interface{}
	Total interface{}
}

func resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}
func respList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}
func RespFail(w http.ResponseWriter, msg string) {
	resp(w, -1, nil, msg)
}
func RespOK(w http.ResponseWriter, data interface{}, msg string) {
	resp(w, 0, data, msg)
}
func RespOKList(w http.ResponseWriter, data interface{}, total interface{}) {
	respList(w, 0, data, total)
}