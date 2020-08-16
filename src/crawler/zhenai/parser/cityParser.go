package parser

import (
	"fmt"
	"learngo/src/crawler/engine"
	"regexp"
)

const regCity = `(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

var reCity = regexp.MustCompile(regCity)

func ParserCityList(contents []byte) engine.ParserResult {
	matchs := reCity.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	limit := 10
	for _, m := range matchs {
		name := string(m[2])
		result.Items = append(result.Items, "City"+name)
		fmt.Printf("get city: %s;", name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: singleCityParser,
		})
		limit--
		if limit <= 0 {
			break
		}
	}
	return result
}
