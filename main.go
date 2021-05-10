package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/radovskyb/watcher"
	"time"
	"log"
)

func main() {
	dirPtr := flag.String("subl", "", "File directory containing subl setting files. (REQUIRED)");

	gitPtr := flag.String("git", "", "File directory containing .git. (REQUIRED)");

	flag.Parse();

	if *dirPtr == "" || *gitPtr == "" {
		fmt.Println("Missing required params.");
		flag.PrintDefaults();
		os.Exit(1);
	}

	// Create a file watcher for each file that matters.
	w := watcher.New();
	w.SetMaxEvents(1);
	w.FilterOps(watcher.Write, watcher.Create);

	go func() {
		for {
			select {
				case event := <-w.Event:
					fmt.Println(event);

				case err := <-w.Error:
					fmt.Println(err);

				case <-w.Closed:
					return;
			}
		}
	}()

	// Begin watching subl dir
	err := w.Add(*dirPtr);
	if err != nil {
		fmt.Println(err);
	}

	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	fmt.Println();

	err = w.Start(time.Millisecond * 100);
	if err != nil {
		log.Fatalln(err);
	}
}
