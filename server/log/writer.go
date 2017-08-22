package main

import (
	"fmt"
	"os"
)

type Writer struct {
	UserFileMap map[string]*os.File
}

func (this *Writer) Write(user, info string) {
	var f *os.File
	if v, ok := this.UserFileMap[user]; ok && v != nil {
		f = v
	} else {
		f, err := os.OpenFile(fmt.Sprintf("/home/aib/apps/ant/logs/%s.log", user), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		this.UserFileMap[user] = f
	}

	f.WriteString(info + "\n")
}
