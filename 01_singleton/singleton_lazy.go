package _1_singleton

import "sync"

var (
	lazySingleton *Singleton
	once          = &sync.Once{} // 只会执行一次
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
