package engine

import (
	"learngo/src/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e *SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)
		for _, m := range parserResult.Items {
			log.Printf("Got item %v", m)
		}
	}
}

func worker(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher err fetching Url %s: %v", r.Url, err)
		return ParserResult{}, nil
	}
	return r.ParserFunc(body), nil
}
