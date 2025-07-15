package main

import (
	"fmt"
	"net/http"
)

func main(){
	resp, err:=http.Get("http://google.com");
	if(err!=nil){
		fmt.Println("err ", err);
	}
	fmt.Println("Resp ", resp);
}