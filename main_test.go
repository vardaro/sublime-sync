package main

import (
	"fmt"
	"os/user"
	"testing"
)

const (
	subl = "/home/vardaro/.config/sublime-text-3/Packages/User/"
	gitp = "/home/vardaro/projects/sublime_text_settings/"
)

/**
Excuse for my laziness
*/
func init() {
	me, err := user.Current()
	if err != nil {
		fmt.Println("Error accessing logon user")
	}

	if me.Username == "vardaro" {
		return
	}

	fmt.Println("Tests won't work on your machine because test paths are hardcoded to Anthony's machine (not that I expect anybody but me to be running tests though.")
}

func TestWatch(t *testing.T) {
	watch(subl, gitp)
}
