package crawler

import (
	"learngo/src/crawler/engine"
	"learngo/src/crawler/zhenai/parser"
)

const targetUrl = "http://www.zhenai.com/zhenghun"

func MainRun() {
	engine.Run(engine.Request{
		Url:        targetUrl,
		ParserFunc: parser.ParserCityList,
	})
}
