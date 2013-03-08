//Use this test with a started fireserver
//TODO: Inegrate fireserver here
package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const ADRESS = "http://localhost"

//First Api Test for /api
func TestOveview(t *testing.T) {
	resp, err := http.Get(ADRESS + ":8080/api")
	if err != nil {
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

	if string(json) != string(body) {
		t.Log("not the same")
		t.Fail()
	}

	fmt.Println(string(body))
}
//Test for /api/c
func TestChatOverview(t *testing.T) {
	resp, err := http.Get(ADRESS + ":8080/api/c")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	if string(body) != string(http.StatusNotFound) {
		t.Log("Not the same")
		t.Fail()
	}

}
