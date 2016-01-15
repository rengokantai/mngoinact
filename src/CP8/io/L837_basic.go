package main
import (
	"bytes"
	"fmt"
	"os"
)

func main(){
	var b bytes.Buffer

	b.Write([]byte("test"))
	//concat string into buffer
	fmt.Fprintf(&b, " success")
	b.WriteTo(os.Stdout)
}