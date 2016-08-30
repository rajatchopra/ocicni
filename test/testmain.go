
package main

import (
	"fmt"
	"github.com/rajatchopra/ocicni"
)

func Main() error {
	plugin, err := InitCNI()
	if err != nil {
		fmt.Sprintf("Error in finding/initializing plugin: %v", err)
	} else {
		fmt.Sprintf("Plugin %v found!", plugin)
	}
	return err
}
