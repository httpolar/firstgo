package main

import (
	"firstgo/cats"
	"sync"
)

func main() {
	ch := make(chan cats.CatImage)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go cats.GetCatUrl(ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for catImage := range ch {
		cats.PrintCat(&catImage)
	}
}
