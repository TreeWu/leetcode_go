package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Memo struct {
	f     Func
	cache map[string]result
}

type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f, make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]

	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() []string {
	return []string{"http://www.baidu.com"}
}
func main() {
	m := New(httpGetBody)
	i := 10
	for i > 0 {
		start:=time.Now()
		get, _ := m.Get("http://localhost:8000/")
		fmt.Println(get)
		i--
		fmt.Printf("%s\n",time.Since(start))
	}
}
