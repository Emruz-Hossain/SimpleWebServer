package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request)  {
	name, exist :=request.URL.Query()["name"]
	if !exist{
		http.NotFound(writer,request)

	}else {
		for i:=0;i<len(name);i++{
			if(len(name[i])>0){
				fmt.Fprintln(writer,"Hello, "+name[i])
			}else {
				fmt.Fprintln(writer,"No name specified")
			}

		}
	}

}
func main()  {
	fmt.Println("Server is running......")
    http.HandleFunc("/hello",handler)
	http.ListenAndServe(":9000",nil)
}