package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
	"strconv"
	"strings"
)

func main(){
	length := 60
	var str string
	var jstr []rune

	jstr = []rune(str)
	f,e := os.Open("/tmp/twpipe")
	f.Close()
	if e != nil{
		exec.Command("mkfifo","/tmp/twpipe").Run()
	}
	go func(){
		sig := make(chan os.Signal,1)
		signal.Notify(sig,syscall.SIGTERM,syscall.SIGINT,syscall.SIGHUP)
		<- sig
		os.Remove("/tmp/twpipe")
		os.Exit(0)
	}()
	for {
		cmd := exec.Command("/home/saffron/go/bin/twty")
		out ,err := cmd.Output()

		if err!=nil{time.Sleep(time.Second * 20);continue}
		//out := strings.Repeat("abcd",300)
		str = string(out)
		str = strings.Replace(str,"\n"," ",-1)
		str = strings.Repeat(" ",length) + str + strings.Repeat(" ",10)
		jstr = []rune(str)
		for i := 0;i<len(jstr)-length;i++{
			func(){
				o , _ := os.OpenFile("/tmp/twpipe",os.O_RDWR,0755)
				defer o.Close()
				fmt.Fprintf(o,"%"+strconv.Itoa(length)+"."+strconv.Itoa(length)+"s\n",string(jstr[i:]))
				time.Sleep(time.Millisecond * 200)
			}()
			//c := exec.Command("clear")
			//c.Stdout = os.Stdout
			//c.Run()
		}
	//	time.Sleep(time.Second * 20)
	}
}
