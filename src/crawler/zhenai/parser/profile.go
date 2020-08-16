package parser

import (
	"fmt"
	"learngo/src/crawler/engine"
	"learngo/src/crawler/model"
	"regexp"
)

var IdRe = regexp.MustCompile(`<div .+ class="id">ID：([^<]+)</div>`)

// 阿坝 | 21岁 | 大专 | 未婚 | 178cm | 3001-5000元
var infoRe = regexp.MustCompile(`<div .+ class="des f-cl">([\w]+) | ([\d]+)岁 | ([\w]+) | ([\w]+) | ([\d]+)cm | ([\w]+)元</div>`)

func ProfileParser(contents []byte, name string) engine.ParserResult {
	provfile := model.Profile{}
	provfile.Name = name
	provfile.Id = extractString(contents, IdRe)
	//info := extractString(contents, infoRe)
	fmt.Println("123123")
	fmt.Println(extractString(contents, IdRe))
	fmt.Println("345345")
	fmt.Println(extractString(contents, infoRe))
	fmt.Println("======")
	return engine.ParserResult{}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) > 2 {
		return string(match[1])
	} else {
		return ""
	}
}
