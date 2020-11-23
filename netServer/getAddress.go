package main

import (
	"fmt"
	"net"
)

func main() {
	ifaces, err := net.Interfaces()
	if err!=nil {
		panic(err)
	}
	fmt.Println("Getting local Ip address: ")
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err!=nil {
			panic(err)
		}
		for _, addr := range addrs {
			switch v:= addr.(type){
			case *net.IPAddr:
				if v.String() !="0.0.0.0"{
					fmt.Println(v.String())
				}
			}
		}
	}
}
