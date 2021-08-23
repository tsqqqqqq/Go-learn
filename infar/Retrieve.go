package infar

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type CodeRetrieve struct {
	url string
}

func (c CodeRetrieve) String() string {
	return fmt.Sprintf("%s", c.url)
}

func (CodeRetrieve) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
