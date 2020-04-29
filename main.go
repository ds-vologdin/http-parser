package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ds-vologdin/http-parser/counter"
)

type StatUrl struct {
	URL   string
	Count int
}

func main() {
	MaxWorkers := flag.Int("max-workers", 5, "max count of worker")
	Word := flag.String("word", "Go", "word for count")
	flag.Parse()

	urls := readUrls(os.Stdin)

	statIn, statDone := statCollector()

	taskCounter := counter.NewTaskCounter(*MaxWorkers)
	for url := range urls {
		// блокирует, если уже запущено максимально допустимое количество задач
		taskCounter.Inc()
		go func(url string) {
			defer taskCounter.Done()
			count, err := httpWordCounter(url, *Word)
			if err != nil {
				fmt.Printf("error: %v", err)
			}
			statIn <- StatUrl{URL: url, Count: count}
		}(url)
	}
	taskCounter.Wait()
	close(statIn)
	<-statDone
}

func statCollector() (chan StatUrl, chan struct{}) {
	in := make(chan StatUrl)
	done := make(chan struct{})
	go func() {
		defer close(done)
		total := 0
		for stat := range in {
			total += stat.Count
			fmt.Printf("Count for %s: %d\n", stat.URL, stat.Count)
		}
		fmt.Printf("Total: %d\n", total)
		done <- struct{}{}
	}()
	return in, done
}

func readUrls(f *os.File) chan string {
	urls := make(chan string)
	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			// TODO: need to validate url
			urls <- scanner.Text()
		}
		close(urls)
	}()
	return urls
}

func httpWordCounter(url, word string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	// resp.Body может быть очень большим, поэтому обрабатываем его построчно
	return countWord(resp.Body, word), nil
}

func countWord(r io.Reader, word string) int {
	count := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		count += strings.Count(line, word)
	}
	return count
}
