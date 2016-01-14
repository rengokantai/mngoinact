package main
import (
	"runtime"
	"sync"
	"fmt"
)



func main(){
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for count:=0;count<2;count++ {
			for char:='a';char<'a'+26;char++{
				fmt.Printf("%c",char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count:=0;count<2;count++ {
			for char:='A';char<'A'+26;char++{
				fmt.Printf("%c",char)
			}
		}
	}()

	fmt.Println("Start")
	wg.Wait()
	fmt.Println("Done")

}
