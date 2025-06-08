package news

import (
	"ago-launcher/api"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type NewsReader struct {
	RemoteNewsItems []RemoteNewsItem `json:"newsItems"`
	NewsItems       []NewsItem
}

type RemoteNewsItem struct {
	Title     string `json:"title"`
	Url       string `json:"url"`
	Date      string `json:"date"`
	Published bool   `json:"published"`
}

type NewsItem struct {
	MarkdownText string
	Title        string
	Date         string
	Published    bool
}

func (newsReader *NewsReader) GetNewsItems() {
	fmt.Println("[News] Starting GetNewsItems")

	// Local
	jsonFile, err := os.Open("resources/newsItems.json")
	if err != nil {
		fmt.Println("[News] could not open resources/newsItems.json file")
	}
	defer jsonFile.Close()

	// Remote
	// resp, err := http.Get("https://raw.githubusercontent.com/EddieEldridge/ago-launcher/refs/heads/main/src/resources/modVersions.json?token=<>")
	// if err != nil {
	// 	fmt.Println("[News] could not fetch modVersions file from GitHub")
	// 	return
	// }
	// defer resp.Body.Close()
	// jsonFile := resp.Body

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("[News] could not open modVersions file")
	}

	jsonErr := json.Unmarshal(byteValue, newsReader)
	if jsonErr != nil {
		fmt.Println("[News] could not unmarshal file")
	}

	fmt.Println("[News] Found " + fmt.Sprintf("%d", len(newsReader.RemoteNewsItems)) + " news items")

	// Initialize NewsItems slice to the correct length
	newsReader.NewsItems = make([]NewsItem, len(newsReader.RemoteNewsItems))
	for i, remoteItem := range newsReader.RemoteNewsItems {
		fmt.Printf("[News] Processing news item %d: %s\n", i, remoteItem.Title)
		newsReader.NewsItems[i].Date = remoteItem.Date
		newsReader.NewsItems[i].Published = remoteItem.Published
		newsReader.NewsItems[i].Title = remoteItem.Title

		// Read the text from the given remote urls
		url := remoteItem.Url
		fmt.Printf("[News] Fetching markdown text from URL: %s\n", url)
		newsText, err := api.GetRemoteText(url)
		if err != nil {
			fmt.Printf("[News] Failed to get read markdown text from remote url %s\n", url)
			newsReader.NewsItems[i].MarkdownText = "**Error reading Markdown text from URL**"
		} else {
			fmt.Printf("[News] Successfully fetched markdown text for: %s\n", remoteItem.Title)
			newsReader.NewsItems[i].MarkdownText = newsText
		}
	}
	fmt.Println("[News] Finished GetNewsItems")
}
