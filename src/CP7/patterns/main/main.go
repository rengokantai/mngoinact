package main

import(
	"log"
	"sync"
	"time"

	"CP7/patterns"
	"CP7/patterns/work"
)

var names=[]string{
	"steve",
	"jason",
}

type namePrinter struct {
	name string
}

func(m *namePrinter)Task(){
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main(){
	p:=work.New(2)
}


