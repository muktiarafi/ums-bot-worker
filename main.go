package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/chromedp/chromedp"
)

type Payload struct {
	URL string `json:"url"`
}

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()
	http.HandleFunc("/screenshots", func(w http.ResponseWriter, r *http.Request) {
		payload := new(Payload)
		json.NewDecoder(r.Body).Decode(payload)

		var buf []byte
		if err := chromedp.Run(ctx, fullScreenshot("http://"+payload.URL, 90, &buf)); err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	})

	log.Println("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
