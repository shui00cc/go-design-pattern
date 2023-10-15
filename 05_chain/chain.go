/*
校园论坛敏感过滤举例：过滤广告、政治敏感
*/
package main

import "fmt"

type WordFilter interface {
	Filter(content string) bool
}

type WordFileterChain struct {
	filters []WordFilter
}

func (c *WordFileterChain) AddFilter(filter WordFilter) {
	c.filters = append(c.filters, filter)
}

func (c *WordFileterChain) Filter(content string) bool {
	for _, filter := range c.filters {
		// 过滤沿链传递，有一个敏感信息直接返回true
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

type AdWordFilter struct{}

func (f *AdWordFilter) Filter(content string) bool {
	// TODO:具体广告过滤
	if content == "ad" {
		return true
	}
	return false
}

type PoWordFilter struct{}

func (f *PoWordFilter) Filter(content string) bool {
	// TODO:具体政治过滤
	if content == "po" {
		return true
	}
	return false
}

func main() {
	chain := &WordFileterChain{}
	chain.AddFilter(&AdWordFilter{})
	chain.AddFilter(&PoWordFilter{})
	fmt.Println(chain.Filter("ad"))
	fmt.Println(chain.Filter("po"))
	fmt.Println(chain.Filter("normal"))
}
