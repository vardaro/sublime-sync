package main

import (
	"testing"
	"fmt"
	"os/user"
)

const (
	subl = "/home/vardaro/.config/sublime-text-3/Packages/User/"
	git = "~/projects/sublime_text_settings/"
)

func init() {
	me, err := user.Current();
	if err != nil {
		fmt.Println("Error accessing logon user");
	}

	if me.Username == "vardaro" {
		return;
	}

	fmt.Println();
	fmt.Println("Tests won't work on your machine because test paths are hardcoded to Anthony's machine (not that I expect anybody but me to be running tests though.");
	fmt.Println();
}

func TestWatch(t *testing.T) {

	watch(subl, git);
}
