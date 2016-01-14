package main
import (
	"sync"
	"runtime"
	"fmt"
)
//single logic processor

var wg sync.WaitGroup
func main(){
	runtime.GOMAXPROCS(2)
	wg.Add(2)
	go printPrime("A")
	go printPrime("B")
	wg.Wait()
}

func printPrime(pre string){
	defer wg.Done()
	next:
	for i:=0;i<50;i++{
		for j:=2;j<i;j++{
			if i%j==0{
				continue next;
			}
		}
		fmt.Printf("%s:%d",pre,i)
	}
	fmt.Println("Done",pre)
}