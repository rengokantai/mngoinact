package main
import (
	"sync"
	"time"
	"fmt"
	"sync/atomic"
)
var (
	shutdown int64
	wg sync.WaitGroup
)

func main(){

	wg.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(3*time.Second)
	fmt.Println("Shut")
	atomic.StoreInt64(&shutdown,1)
	wg.Wait()
}


func doWork(name string){
	defer wg.Done()

	for{
		time.Sleep(250*time.Millisecond)
		fmt.Printf("%s",name)
		if atomic.LoadInt64(&shutdown)==1{
			fmt.Printf("%s get",name)
			break
		}
	}
}