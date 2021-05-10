package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/radovskyb/watcher"
	"github.com/go-git/go-git/v5"
	"time"
	"log"
	"strings"
	"io"
)

/**
	Given a src path and a copy path,
	copy the contents of the src file into the copy file.
*/
func copy(src string, dst string) error {
	in, err := os.Open(src);
	if err != nil {
		return err;
	}
	defer in.Close();

	out, err := os.Create(dst);
	if err != nil {
		return err;
	}
	defer out.Close();

	_, err = io.Copy(out, in);
	if err != nil {
		return err;
	}

	return out.Close();
}

/**
	Watches the subl directory

	If theres a change, updates the git directory and performs a
		- add
		- commit
		- push

		(if possible)
*/
func watch(subl string, git string) {
	w := watcher.New();
	w.SetMaxEvents(1);
	w.FilterOps(watcher.Write, watcher.Create);

	go func() {
		for {
			select {
				case event := <-w.Event:
					fmt.Println(event);

					// Name of file in corresponding git repo
					dst := strings.Replace(event.Path, subl, git, 1);

					// Update git file
					err := copy(event.Path, dst);
					if err != nil {
						fmt.Println(err);
					}



				case err := <-w.Error:
					log.Fatalln(err);

				case <-w.Closed:
					return;
			}
		}
	}()

	// Begin watching subl dir
	err := w.Add(subl);
	if err != nil {
		log.Fatalln(err);
	}

	for path, f := range w.WatchedFiles() {
		fmt.Printf(" Watching %s: %s\n", path, f.Name())
	}

	fmt.Println();

	err = w.Start(time.Millisecond * 100);
	if err != nil {
		log.Fatalln(err);
	}	
}

func main() {
	dirPtr := flag.String("subl", "", "File directory containing subl setting files. (REQUIRED)");

	gitPtr := flag.String("git", "", "File directory containing .git. (REQUIRED)");

	flag.Parse();

	if *dirPtr == "" || *gitPtr == "" {
		fmt.Println("Missing required params.");
		flag.PrintDefaults();
		os.Exit(1);
	}

	watch(*dirPtr, *gitPtr);
}
