//Use this test with a started fireserver
//TODO: Inegrate fireserver here
package server

import (
//	"io/ioutil"
	"net/http"
	"testing"
)

const ADRESS = "http://localhost"

func TestShutdown(t *testing.T) {
	port := ":8080"
	fireserver := NewFireServer(port)
	go fireserver.ListenAndServe()
	fireserver.Close()

	_, err := http.Get("http://localhost"+port)
	if err == nil{
		t.Fail()
	}
}

//api POST /api/c
func TestAddChatRoom(t *testing.T){
	port := ":8081"
	fireserver := NewFireServer(port)
	go fireserver.ListenAndServe()

	t.Error("Not implemented yet")
}

/*func TestHello(t *testing.T) {
	//defer srv.Close()
	response, err := http.Get("http://localhost" + port + "/hello")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if string(body) != "Hello World" {
		t.Log("Answer is not as expected")
		t.Fail()
	}

	response.Body.Close()

	//srv.Close()

}*/
