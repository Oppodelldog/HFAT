package main

import (
	"HFAT/server"
	"fmt"
	"HFAT/config"
)

func main() {
	conf := config.ReadConfigFromFile()
	fmt.Printf("starting HFAT server with following targets:\n")
	fmt.Printf("%+v\n", conf)

	server.StartHFATServer(conf.Port,conf.ForwardingTargets)
}

