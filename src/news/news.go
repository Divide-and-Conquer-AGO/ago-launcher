package news

import (
	"ago-launcher/api"
	"ago-launcher/utils"
	"encoding/json"
	"io"
	"net/http"
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
	utils.Logger().Println("[News] Starting GetNewsItems")

	// Local
	// jsonFile, err := os.Open("resources/newsItems.json")
	// if err != nil {
	// 	utils.Logger().Println("[News] could not open resources/newsItems.json file")
	// }
	// defer jsonFile.Close()

	// Remote
	resp, err := http.Get("https://raw.githubusercontent.com/Divide-and-Conquer-AGO/ago-launcher/refs/heads/main/src/resources/newsItems.json")
	if err != nil {
		utils.Logger().Println("[News] could not fetch modVersions file from GitHub")
		return
	}
	defer resp.Body.Close()
	jsonFile := resp.Body

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		utils.Logger().Println("[News] could not open modVersions file")
	}

	jsonErr := json.Unmarshal(byteValue, newsReader)
	if jsonErr != nil {
		utils.Logger().Println("[News] could not unmarshal file")
	}

	utils.Logger().Printf("[News] Found %d news items\n", len(newsReader.RemoteNewsItems))

	// Initialize NewsItems slice to the correct length
	newsReader.NewsItems = make([]NewsItem, len(newsReader.RemoteNewsItems))
	for i, remoteItem := range newsReader.RemoteNewsItems {
		utils.Logger().Printf("[News] Processing news item %d: %s\n", i, remoteItem.Title)
		newsReader.NewsItems[i].Date = remoteItem.Date
		newsReader.NewsItems[i].Published = remoteItem.Published
		newsReader.NewsItems[i].Title = remoteItem.Title

		// Read the text from the given remote urls
		url := remoteItem.Url
		utils.Logger().Printf("[News] Fetching markdown text from URL: %s\n", url)
		newsText, err := api.GetRemoteText(url)
		if err != nil {
			utils.Logger().Printf("[News] Failed to get read markdown text from remote url %s\n", url)
			newsReader.NewsItems[i].MarkdownText = "**Error reading Markdown text from URL**"
		} else {
			utils.Logger().Printf("[News] Successfully fetched markdown text for: %s\n", remoteItem.Title)
			newsReader.NewsItems[i].MarkdownText = newsText
		}
	}
	utils.Logger().Println("[News] Finished GetNewsItems")
}
