package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func walkDir(dir string, fileSize chan<- int64) {

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}

}
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprint(os.Stderr, "dul: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	/*	f := func() {
		dir := "D:\\"
		fileSize := make(chan int64)
		n := &sync.WaitGroup{}
		go func() {
			walkDir(dir, n, fileSize)
		}()
		go func() {
			n.Wait()
			close(fileSize)
		}()

		tick := time.Tick(500 * time.Millisecond)
		var nFiles, nbytes int64

		for fileSize != nil {
			select {
			case size, ok := <-fileSize:
				if ok {
					nFiles++
					nbytes += size
				} else {
					fileSize = nil
				}
			case <-tick:

				printDisUsage(nFiles, nbytes)
			}
		}
	}*/

	f := func() {
		dir := "D:\\"
		fileSize := make(chan int64)
		go func() {
			walkDir(dir, fileSize)
			close(fileSize)
		}()

		tick := time.Tick(500 * time.Millisecond)
		var nFiles, nbytes int64

		for fileSize != nil {
			select {
			case size, ok := <-fileSize:
				if ok {
					nFiles++
					nbytes += size
				} else {
					fileSize = nil
				}
			case <-tick:

				printDisUsage(nFiles, nbytes)
			}
		}
	}
	rTime(f)
}

func rTime(f func()) {
	start := time.Now()
	defer func() {
		end := time.Now()
		rtime := end.Sub(start)
		fmt.Println("runtime :", rtime.Seconds())
	}()
	f()
}

func printDisUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
