package main

import (
	"flag"
	"fmt"
	"github.com/rajatchopra/ocicni"
)

func init() {
	flag.Parse()
}

func main() {
	plugin, err := ocicni.InitCNI("")
	if err != nil {
		fmt.Printf("Error in finding/initializing plugin: %v", err)
		return
	} else {
		fmt.Printf("Plugin %v found!\n", plugin)
	}

	err = plugin.SetUpPod("my_netnamespace", "my_namespace", "my_name", "my_containerid")
	if err != nil {
		fmt.Printf("Error in calling SetUpPod: %v\n", err)
		return
	} else {
		fmt.Println("Setup called without error")
	}
	return
}
