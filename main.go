package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func getArgs() (dir string, search string, err error) {
	if len(os.Args) == 2 {
		dir = "."
		search = os.Args[1]
	} else if len(os.Args) == 3 {
		dir = os.Args[1]
		search = os.Args[2]
	} else {
		return "", "", errors.New("no args provided")
	}

	return dir, search, nil
}

func explore(dirPath string, search *string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("searching in", dirPath)

	dir, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Could not read dir", dirPath)
	}

	for _, entry := range dir {
		entrypath := filepath.Join(dirPath, entry.Name())

		if entry.Name() == *search {
			fmt.Println("FOUND" + entrypath)
		}

		if entry.IsDir() {
			wg.Add(1)
			go explore(entrypath, search, wg)
		}
	}
}

func main() {
	start := time.Now()

	dir, search, err := getArgs()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	dir, err = filepath.Abs(dir)
	if err != nil {
		panic("invalid starting directory")
	}

	go explore(dir, &search, &wg)
	wg.Add(1)

	go func() {
		wg.Wait()
	}()

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
