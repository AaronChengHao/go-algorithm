// 单例模式
package designPattern

import "sync"

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
