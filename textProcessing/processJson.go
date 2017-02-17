package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName string
	ServerIp   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":
	"Shanghai_VPN",
	"serverIP":"127.0.0.1"},
	{"serverName":"Beijing_VPN",
	"serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s.Servers[0])

	str2 := `{"Name":"Wednesday","Age":6,
	"Parents":["Gomez","Morticia"]}`

	js, err := NewJson([]byte(str2))
	name := js.get("Name")
	fmt.Println(name)

}
