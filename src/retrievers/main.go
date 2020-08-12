package retrievers

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"

func Download(r Retriever) string {
	return r.Get(url)
}
func Session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "another faked imooc.com"})
	return s.Get(url)
}
