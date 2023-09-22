package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var lazySingleton *Singleton
var once sync.Once // 只会执行一次

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}

func main() {
	s := GetLazyInstance()
	fmt.Println(s)
}
