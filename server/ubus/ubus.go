package ubus

import (
	"fmt"
	"log"
	"github.com/cdavid14/goubus"
)
func test_ubus(){
	ubus := goubus.Ubus{
	  Username: "root",
	  Password: "admin",
	  URL:      "http://127.0.0.1/ubus",
	}
	result, err := ubus.Login()
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println(result)
}


