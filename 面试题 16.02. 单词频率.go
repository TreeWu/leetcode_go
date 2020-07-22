package main

func main() {

}

type WordsFrequency struct {
	count map[string]int
}

func Constructor(book []string) WordsFrequency {
	w := WordsFrequency{
		count: make(map[string]int),
	}
	for i := 0; i < len(book); i++ {
		if _, ok := w.count[book[i]]; ok {
			w.count[book[i]]++
		} else {
			w.count[book[i]] = 1
		}
	}

	return w
}

func (this *WordsFrequency) Get(word string) int {
	if count, ok := this.count[word]; ok {
		return count
	}
	return 0
}
