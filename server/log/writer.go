package main

import (
	"fmt"
	"os"
)

type Writer struct {
	UserFileMap map[string]*os.File
}

func (this *Writer) Write(user, pid, ts, info string) {
	k := fmt.Sprintf("%s|%s|%s", user, pid, ts)
	var f *os.File
	if v, ok := this.UserFileMap[k]; ok && v != nil {
		f = v
	} else {
		f, err := os.OpenFile(fmt.Sprintf("/home/aib/apps/ant/logs/%s-%s.%s.log", user, pid, ts), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		this.UserFileMap[k] = f
	}

	f.WriteString(info + "\n")
}
