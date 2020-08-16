package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityParser_test_data.html")
	if err != nil {
		panic(err)
	}
	parserResult := ParserCityList(contents)
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"闃垮潩", "闃垮厠鑻廫n", "闃挎媺鍠勭洘",
	}
	for i, Url := range expectedUrls {
		if parserResult.Requests[i].Url != Url {
			t.Errorf("expected Url #%d: %s, but was %s", i, Url, parserResult.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		if parserResult.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s, but was %s", i, city, parserResult.Items[i].(string))
		}
	}
	if len(parserResult.Requests) != resultSize {
		t.Errorf("Result should be %d, but had %d", resultSize, len(parserResult.Requests))
	}
	if len(parserResult.Items) != resultSize {
		t.Errorf("Result should be %d, but had %d", resultSize, len(parserResult.Items))
	}
}
