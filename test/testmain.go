
package main

import (
	"fmt"
	"github.com/rajatchopra/ocicni"
)

func main() {
	plugin, err := ocicni.InitCNI("")
	if err != nil {
		fmt.Printf("Error in finding/initializing plugin: %v", err)
	} else {
		fmt.Printf("Plugin %v found!", plugin)
	}
	return
}
