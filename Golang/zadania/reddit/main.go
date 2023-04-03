package main

import (
	"io"
	"log"
	"os"
	"reddit/fetcher"
)

func main() {
	var f fetcher.RedditFetcher // do not change
	var w io.Writer             // do not change
	var w2 io.Writer

	f = &fetcher.MyFetcher{}

	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Nie można otwrzyc pliku!")
	}
	defer file.Close()
	w = file

	w2 = os.Stdout

	err = f.Fetch()
	if err != nil {
		log.Fatalf("%v", err)
	}
	f.Save(w2)
	f.Save(w)

	file.Close()
}
