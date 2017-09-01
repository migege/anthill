package main

import (
	"fmt"
	"os"
	"time"
)

type Writer struct {
	UserFileMap map[string]struct {
		F  *os.File
		Ts int64
	}
}

func (this *Writer) GC() {
	var l []string
	now := time.Now().Unix()
	for k, v := range this.UserFileMap {
		if v.F != nil && now-v.Ts > 10*60 {
			v.F.Close()
			l = append(l, k)
		}
	}
	for _, k := range l {
		delete(this.UserFileMap, k)
	}
}

func (this *Writer) Write(k, info string) {
	var f *os.File
	if v, ok := this.UserFileMap[k]; ok && v.F != nil {
		f = v.F

		this.UserFileMap[k] = struct {
			F  *os.File
			Ts int64
		}{f, time.Now().Unix()}
	} else {
		fn := fmt.Sprintf("/home/aib/apps/ant/logs/%s.log", k)
		f, err := os.OpenFile(fn, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

		this.UserFileMap[k] = struct {
			F  *os.File
			Ts int64
		}{f, time.Now().Unix()}
	}

	f.WriteString(info + "\n")
}
