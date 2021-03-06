package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"model"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("コマンドライン引数エラー")
	}

	storeId := args[0]

	// Fetch
	url := fmt.Sprintf("https://itunes.apple.com/jp/rss/customerreviews/id=%s/json", storeId)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := &model.Review{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("tmp.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Write CSV
	w := csv.NewWriter(f)
	colmuns := []string{"id", "version", "name", "date", "title", "content"}
	if err := w.Write(colmuns); err != nil {
		log.Fatal(err)
	}

	for _, entry := range data.Feed.Entry {
		id := entry.ID.Label
		version := entry.ImVersion.Label
		name := entry.Author.Name.Label
		date := entry.Updated.Label
		title := strings.Replace(entry.Title.Label, "\n", "", -1)
		content := strings.Replace(entry.Content.Label, "\n", "", -1)

		line := []string{id, version, name, date, title, content}
		if err := w.Write(line); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("SUCCESSED!")
}
