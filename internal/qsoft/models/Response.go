package models

import "errors"

var (
	ErrBadParam    = errors.New("bad param")
	ErrParamNotNum = errors.New("param not num")
)

type Response struct {
	Error     bool        `json:"error"`
	ErrorText string      `json:"errorText"`
	Data      interface{} `json:"data"`
	Code      int         `json:"code"`
}
