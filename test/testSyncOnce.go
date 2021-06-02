package main

import "sync"

var loadOnce sync.Once
var Content map[string]string

func Load(key string) string {

	loadOnce.Do(initContent)
	return Content[key]
}

func initContent() {
	Content = make(map[string]string)
	Content["tom"] = "tom"
}

func main() {
	i := 10
	for i > 0 {
		go func() {
			Load("tom")
		}()

		i--
	}

	select {}
}
