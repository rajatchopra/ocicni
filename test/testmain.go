
package main

import (
	"fmt"
	"github.com/rajatchopra/ocicni"
)

func main() {
	plugin, err := ocicni.InitCNI("")
	if err != nil {
		fmt.Sprintf("Error in finding/initializing plugin: %v", err)
	} else {
		fmt.Sprintf("Plugin %v found!", plugin)
	}
	return
}
