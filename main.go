package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"log"
	"strings"
	"io"

	"github.com/radovskyb/watcher"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

/**
	Pushes a file change to Github.
*/
func push(gitp string, commitMsg string) {
	ref, err := git.PlainOpen(gitp);
	if err != nil {
		fmt.Println(err);
	}

	wtree, err := ref.Worktree();
	if err != nil {
		fmt.Println(err);
	}

	// git add . 
	_, err = wtree.Add(".");
	if err != nil {
		fmt.Println(err);
	}

	// git commit -m commitMsg
	_, err = wtree.Commit(commitMsg, &git.CommitOptions{});
	if err != nil {
		fmt.Println(err);
	}

	// git push	
	err = ref.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: os.Getenv("GH_USER"),
			Password: os.Getenv("GH_PASS"),
		},
		Progress: os.Stdout,
	});
	if err != nil {
		fmt.Println(err);
	}
}

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
func watch(subl string, gitp string) {
	w := watcher.New();
	w.SetMaxEvents(1);
	w.FilterOps(watcher.Write, watcher.Create);

	go func() {
	  for {
		  select {
			  case event := <-w.Event:

					// Name of file in corresponding git repo
				  dst := strings.Replace(event.Path, subl, gitp, 1);

					// Update git file
				  err := copy(event.Path, dst);
				  if err != nil {
				  	fmt.Println(err);
				  }

					commitMsg := event.String();

					push(gitp, commitMsg);

					fmt.Println(commitMsg);

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
