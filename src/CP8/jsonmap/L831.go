package main
import (
	"encoding/json"
	"log"
	"fmt"
)

func main(){
	c:=make(map[string]interface{})
	c["name"]="Go"
	c["title"]="title"
	c["contact"]=map[string]interface{}{
		"home":"12345",
		"cell":"67890",
	}

	data,err:=json.MarshalIndent(c,"","   ")
	if err!=nil{
		log.Println("Err:",err)
		return
	}
	fmt.Println(string(data))
}
