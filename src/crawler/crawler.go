package crawler

import (
	"learngo/src/crawler/engine"
	"learngo/src/crawler/scheduler"
	"learngo/src/crawler/zhenai/parser"
)

const targetUrl = "http://www.zhenai.com/zhenghun"

func MainRun() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        targetUrl,
		ParserFunc: parser.ParserCityList,
	})
}
