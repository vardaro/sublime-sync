package main

import (
	"flag"
	"fmt"
)

// Ubuntu path
const UBUNTU_SUBL_PATH string = "~/.config/sublime-text-3/Packages/User"

var SUBL_SETTING_FILENAMES = [2]string{"Preferences.sublime-settings", "Package Control.sublime-settings"}

func main() {
	dirPtr := flag.String("subl", UBUNTU_SUBL_PATH, "File directory containing subl setting files.");

	gitPtr := flag.String("git", "", "File directory containing .git");


	fmt.Println(*dirPtr);
	fmt.Println(*gitPtr)
}
