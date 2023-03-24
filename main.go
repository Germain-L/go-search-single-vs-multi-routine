package main

import (
	"errors"
	"fmt"
	"os"
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

func explore(dirPath string, search *string) {
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Could not read dir", dirPath)
	}

	for _, entry := range dir {
		entrypath := dirPath + "/" + entry.Name()

		if entry.Name() == *search {
			fmt.Println("FOUND" + entrypath)
		}

		if entry.IsDir() {
			explore(entrypath, search)
		}
	}
}

func main() {
	start := time.Now()

	dir, search, err := getArgs()
	if err != nil {
		panic(err)
	}

	explore(dir, &search)

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)

}
