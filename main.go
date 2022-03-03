package main

// Copyright (C) 2022 - rootqa
// GPL License

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func get_file(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
} /* file parser and get request */

func set_target(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	println("Downloading...")
	target_url := "TARGET_URL"
	err := get_file("FILE_NAME", target_url)
	if err != nil {
		panic(err)
	}
} /* download loop */

func main() {
	// println("GOMAXPROCS", runtime.GOMAXPROCS(2))
	var wg sync.WaitGroup
	start := time.Now()

	for i := 1; i < n; i++ {
		println("Generating thread: ", i)
		wg.Add(1)
		go set_target(i, &wg)
	}

	wg.Wait()
	fmt.Println("Finished the thread: ", time.Since(start))

	finish := time.Now()
	log.Println("start: ", start)
	log.Println("finish: ", finish)
	fmt.Println("All worker have finished.")
} /* main */
