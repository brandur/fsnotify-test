package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

const fileToWatch = "./foo"

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			log.Println("start loop")

			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Println("not ok")
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(fileToWatch)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
