package search

import "log"

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	//searchTerm 搜索词
	search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result)  {

	searchResult, err := matcher.search(feed, searchTerm)

	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResult {
		results <- result
	}
}

func Display(results chan *Result)  {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}


