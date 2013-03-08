package server

import(
	"testing"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

const ADRESS = "http://localhost"

//First Api Test for /api
func TestOveview(t *testing.T){
	resp, err := http.Get(ADRESS+":8080/api")
	if err != nil{
		t.Log(err)
		t.Fail()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	json, er := json.Marshal(restCommands)
	
	if er != nil {
		t.Log(er)
		t.Fail()
	}

	if string(json) != string(body){ 
		t.Log("not the same")
		t.Fail()
	}


	fmt.Println(string(body))
}
