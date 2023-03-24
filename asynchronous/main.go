package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func explore(dirPath string, search *string, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()

	dir, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Could not read dir", dirPath)
		return
	}

	for _, entry := range dir {
		entryPath := filepath.Join(dirPath, entry.Name())

		if strings.Contains(entry.Name(), *search) {
			resultChan <- entryPath
		}

		if entry.IsDir() {
			wg.Add(1)
			go explore(entryPath, search, wg, resultChan)
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
	resultChan := make(chan string)

	wg.Add(1)
	go explore(dir, &search, &wg, resultChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Println("FOUND:", result)
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
