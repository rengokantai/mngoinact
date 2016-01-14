package main
import (
	"sync"
	"fmt"
	"sync/atomic"
	"runtime"
)
//atomic operation


var wg sync.WaitGroup
var counter int64

func main(){
	wg.Add(2)
	go incre(1)
	go incre(2)
	wg.Wait()
	fmt.Println("Final: ",counter)
}

func incre(id int){

	defer wg.Done()
	for i:=0;i<2;i++{
		atomic.AddInt64(&counter,1)
		runtime.Gosched()
	}

}

