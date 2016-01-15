package main
import (
	"log"
	"os"
	"io/ioutil"
	"io"
)

var (
	Trace *log.Logger
	Info  *log.Logger
	Warning *log.Logger
	Error *log.Logger
)

func init(){
	file,err := os.OpenFile("err.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("Failed",err)
	}
	Trace = log.New(ioutil.Discard,"TRACE: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	Info = log.New(os.Stdout,"Info:",log.Ldate|log.Ltime|log.Lshortfile)
	Warning=log.New(os.Stdout,"Warning: ",log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file,os.Stderr),"Error: ",log.Ldate|log.Lshortfile)
}

func main(){
	Trace.Println("trace")
	Info.Println("info")
	Warning.Println("warning")
	Error.Println("error")
}
