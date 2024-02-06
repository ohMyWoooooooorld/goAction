package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

//run执行搜索逻辑  perform【做】
func Run(searchTerm string)  {
	//检索要搜索的源列表。Retrieve(取回) the list of feeds to search through.
	//feed 可以被翻译为“信息源”或“数据源
	feeds, err := RetrieveFeeds()
	if err != nil {
		//致命错误
		log.Fatal(err)
	}

	//创建一个无缓冲通道以接收要显示的匹配结果
	results := make(chan *Result)

	//sync.WaitGroup 是一个用于等待一组并发操作完成的同步原语。
	//它提供了一种简单的方法来阻塞当前goroutine，直到一组其他的goroutine都完成执行
	var waitGroup sync.WaitGroup

	//增加WaitGroup的计数器。通常，在启动一个新的goroutine时调用此方法，传递一个正整数（通常是1），表示有一个goroutine即将开始执行。
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {

		//matchers是一个键值对，feed.type 是key，matcher获取值
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		//传入入matcher, feed； 立即调用
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)


	}
	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)
}