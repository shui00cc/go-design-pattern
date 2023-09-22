package main

import "fmt"

// 饿汉式单例
type Singleton struct{}

var singleton *Singleton = new(Singleton)

func GetInstance() *Singleton {
	return singleton
}

func main() {
	single := GetInstance()
	fmt.Println(single)
}
