package otto_wrapper

/*
	This file is not generated
*/

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/robertkrimen/otto"
)

func NewOttoFunctions() *OttoFunctions {
	o := &OttoFunctions{
		vm:   otto.New(),
		logs: make(chan string, 10),
	}
	o.registFunctions()
	return o
}

type OttoFunctions struct {
	vm   *otto.Otto
	logs chan string
}

func (o *OttoFunctions) GetVM() *otto.Otto {
	return o.vm
}

func (o *OttoFunctions) debugMessage(msg string) string {
	t := time.Now()
	ts := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d %03d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000000)
	c := o.vm.Context()
	return fmt.Sprintf("%s (debug: func: %s, row: %d, col: %d) %s", ts, c.Filename, c.Line, c.Column, msg)
}

func (o *OttoFunctions) handleError(err error) otto.Value {
	if o.logs != nil && len(o.logs) != cap(o.logs) {
		o.logs <- o.debugMessage(err.Error())
	}
	v, _ := o.vm.ToValue(o.debugMessage(err.Error()))
	return v
}

func bytesToBase64str(bs []byte) string {
	return base64.StdEncoding.EncodeToString(bs)
}

func base64strToBytes(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
