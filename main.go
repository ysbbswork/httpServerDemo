package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)



type Page struct {
	Qua string
	Ans string
}


//func SimpleServer(w http.ResponseWriter, request *http.Request) {
//	io.WriteString(w, "hello, world")
//}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	strSrc,err :=os.Getwd()
	fileSrc :=strSrc+"/templet.html"
	t, err := template.ParseFiles(fileSrc)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	switch request.Method {
	case "GET":
		p := Page{Qua:"你好，我是AI，你可以向我提问哦！",Ans:"例如：今天我帅吗？"}
		if err := t.Execute(w, p); err != nil {
			fmt.Println("There was an error:", err.Error())
		}
	case "POST":
		input:= request.FormValue("in")
		fmt.Println((input))
		output := AIcore(input)
		input = "我: "+input
		output = "AI: "+output
		p := Page{Qua:input,Ans:output}
		if err := t.Execute(w, p); err != nil {
			fmt.Println("There was an error:", err.Error())
		}

	}

}

func AIcore(input string) (output string) {
	output = strings.Replace(input,"吗","",-1)
	output = strings.Replace(output,"?","!",-1)
	output = strings.Replace(output,"？","!",-1)
	output = strings.Replace(output,"我","你",-1)
	output = strings.Replace(output,"么","",-1)
	return
}


func main() {
	//http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/ai_chat", logPanics(FormServer))
	if err := http.ListenAndServe(":8088", nil); err != nil {
	}
}

func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}
