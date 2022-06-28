package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

//初始化变量
var
(
	targets string
	radName	string
	xrayName string
	xrayProxy string
	)

func init()  {
	//命令行获取参数
	flag.StringVar(&targets,"t","target.txt","set target file")
	flag.StringVar(&radName,"r","rad.exe","set rad name")
	flag.StringVar(&xrayName,"x","xray.exe","set xray name")
	flag.StringVar(&xrayProxy,"proxy","127.0.0.1:7777","set you  proxy")
}

//rad爬取函数
func Scan(target string) {
	fmt.Println(target)
	scanshell := radName+" -t "+target+" -http-proxy "+xrayProxy
	fmt.Println(scanshell)
	cmd:= exec.Command("cmd","/C",scanshell)
	err:=cmd.Start()
	if err!=nil {
		log.Fatal(err)
	}
	//exec.Command("cmd","/c",scanshell)
}

func main()  {
	flag.Parse()
	file,err := os.OpenFile(targets,os.O_RDONLY,0666)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	cmd_ex :="start "+xrayName+" webscan"+" --listen "+xrayProxy+" --html-output "+string(timestamp)+".html"
	cmd:= exec.Command("cmd","/C",cmd_ex)
	err=cmd.Start()
	if err!=nil {
		log.Fatal(err)
	}
	for {
		line,_,err :=reader.ReadLine()
		if err == io.EOF {
			break
		}
		Scan(string(line))
	}
	
}