package cats

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

const url = "https://api.thecatapi.com/v1/images/search"

type CatImage struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

func GetCatUrl(ch chan<- CatImage, wg *sync.WaitGroup) CatImage {
	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating a request to the Cat API: %s\n", err)
		return CatImage{}
	}

	catKey := os.Getenv("CAT_KEY")
	if len(catKey) == 0 {
		fmt.Println("CAT_KEY env variable is not set!")
		return CatImage{}
	}

	req.Header.Set("x-api-key", catKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending a request to the Cat API: %s\n", err)
		return CatImage{}
	}

	defer res.Body.Close()

	var cats []CatImage
	err = json.NewDecoder(res.Body).Decode(&cats)
	if err != nil {
		fmt.Println("Error decoding response JSON body")
		return CatImage{}
	}

	ch <- cats[0]
	return cats[0]
}

func PrintCat(catImage *CatImage) {
	fmt.Printf("ID: %s\nURL: %s\nWidth: %v\nHeight: %v\n", catImage.ID, catImage.URL, catImage.Width, catImage.Height)
}
