package parser

import (
	"fmt"
	"learngo/src/crawler/engine"
	"regexp"
)

// <a href="http://album.zhenai.com/u/1749192296" target="_blank">喝牛奶吐泡泡</a>
var regSigleCity = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`
var resingleCity = regexp.MustCompile(regSigleCity)

func singleCityParser(contents []byte) engine.ParserResult {
	matchs := resingleCity.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matchs {
		name := string(m[2])
		Url := string(m[1])
		result.Items = append(result.Items, "User"+name)
		fmt.Printf("get User: %s; get User_Url: %s;\n", name, Url)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(Url),
			ParserFunc: engine.NilParserFunc,
			//	func(bytes []byte) engine.ParserResult {
			//	return ProfileParser(bytes, name)
			//},
		})
	}
	return result
}
