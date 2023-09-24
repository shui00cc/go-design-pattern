package main

import "fmt"

// 抽象目标
type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

// 抽象观察者
type IObserver interface {
	Update(msg string)
}

// 实现目标的接口
type Subject struct {
	observers []IObserver
}

// 为目标注册观察者
func (sub *Subject) Register(observer IObserver) {
	sub.observers = append(sub.observers, observer)
}

// 移除观察者
func (sub *Subject) Remove(observer IObserver) {
	for i, ob := range sub.observers {
		if ob == observer {
			sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
		}
		break
	}
}

// 通知
func (sub *Subject) Notify(msg string) {
	for _, o := range sub.observers {
		o.Update(msg)
	}
}

// 实例化观察者1
type Observer1 struct {
}

func (Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s\n", msg)
}

// 实例化观察者2
type Observer2 struct {
}

func (Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s\n", msg)
}

func main() {
	// 实例化目标
	sub := &Subject{}
	// 注册观察者
	sub.Register(&Observer1{})
	sub.Notify("hello 1")
	sub.Register(&Observer2{})
	sub.Notify("hello 1and2")
}
