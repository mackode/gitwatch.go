package main

import (
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type Notifier struct {
	watcher *fsnotify.Watcher
}

func NewNotifier() *Notifier {
	return &Notifier{}
}

func (n *Notifier) Wait() {
	select {
	case <-n.watcher.Events:
	case <-n.watcher.Errors:
	}
}

func (n *Notifier) Start(rootDir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	watcher.Add(filepath.Join(gitTopDir(), ".git/HEAD"))
	watcher.Add(filepath.Join(gitTopDir(), ".git/refs/remotes/origin"))

	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if filepath.Base(path) == ".git" {
				return filepath.SkipDir // avoid .git loops
			}

			if err := watcher.Add(path); err != nil {
				return err
			}
		}

		return nil
	})

	n.watcher = watcher
}
