package main

import (
	"testing"
)

const (
	subl = "/home/vardaro/.config/sublime-text-3/Packages/User/"
	git = "~/projects/sublime_text_settings/"
)

func TestWatch(t *testing.T) {

	watch(subl, git);
}
