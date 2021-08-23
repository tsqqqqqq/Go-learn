package infar

import "fmt"

type TestRetrieve struct {
	Url string
}

func (t TestRetrieve) String() string {
	return fmt.Sprintf("%s", t.Url)
}

func (TestRetrieve) Get(url string) string {
	return "test"
}
