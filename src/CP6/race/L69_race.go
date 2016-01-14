
package main
import (
	"sync"
	"runtime"
	"fmt"
)

var wg sync.WaitGroup
var counter int

func main(){
	//runtime.GOMAXPROCS(2)
	wg.Add(2)



	go incre(1)
	go incre(2)
	wg.Wait()
	fmt.Print(counter)
}

func incre(id int){
	defer wg.Done()
	for i:=0;i<2;i++{
		value:=counter
		runtime.Gosched()
		value++;
		counter=value
	}

}

