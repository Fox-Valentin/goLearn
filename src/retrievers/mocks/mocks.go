package mocks

type MockRetrievers struct {
	Contents string
}

func (r *MockRetrievers) Get(url string) string {
	return r.Contents
}

func (r *MockRetrievers) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
