package main

import (
	"log"
	"os"
)

func main() {
	rootDir := "."
	if len(os.Args) > 1 {
		rootDir = os.Args[1]
	}
	err := os.Chdir(rootDir)
	if err != nil {
		panic(err)
	}

	notify := NewNotifier()
	notify.Start(rootDir)
	app, cmds := ui()

	go func() {
		for {
			statuses, err := gitStatus()
			if err != nil {
				log.Printf("Status error: %v\n", err)
				break
			}
			pstatus, err := gitPushStatus()
			if err != nil {
				log.Printf("Push status error: %v\n", err)
				break
			}
			cmds <- Cmd{fs: statuses, pstatus: pstatus}
			notify.Wait()
		}
	}()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
