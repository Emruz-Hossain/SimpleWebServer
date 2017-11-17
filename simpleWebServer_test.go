package main

import (
	"net/http"
	"testing"
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

func TestServer(t *testing.T) {
	t.Run("Test-1", testServer(1,"http://127.0.0.1:9000","ok",404,"404 page not found")) //only connection test with currect url
	t.Run("Test-2", testServer(2,"http://127.0.0.1:8080","refused",0,""))// connection test with wrong url
	t.Run("Test-3", testServer(3,"http://127.0.0.1:9000/hello?name=emruz","ok",200,"Hello, emruz"))//response test with valid data
	t.Run("Test-4", testServer(4,"http://127.0.0.1:9000/hello?name=","ok",200,"No name specified"))//response test with invalid data
	}

func testServer(testNo int, url string, expectedConnection string, expectedStatusCode int, expectedResponse string) func(*testing.T) {
	return func(t *testing.T) {
		fmt.Println("Running test ",testNo)
		response,err := http.Get(url)
		if err != nil{
				if(expectedConnection=="ok"){
					t.Error("Expected sucessfull connection but connection refused.")
				}else{
					fmt.Println("Test "+strconv.Itoa(testNo)+" PASS")
				}
		}else {
			if expectedStatusCode!=response.StatusCode{
				t.Error("Expected status code ",expectedStatusCode," found ",response.StatusCode)
			}else {
				bodyText,err:=ioutil.ReadAll(response.Body)
				if err!=nil{
					t.Error("Expected text on body but found empty")
				}else{
					tmp:=string(bodyText)
					tmp=strings.TrimSuffix(tmp,"\n")

					if tmp!=expectedResponse{
						t.Error("Expected ", expectedResponse," found ",tmp)
					}else{
						fmt.Println("Test "+strconv.Itoa(testNo)+" PASS")
					}
				}
			}


			//fmt.Println(response.Body)
			defer response.Body.Close()
		}

	}
}