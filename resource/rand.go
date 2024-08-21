package resource

import (
	"math/rand"
	"sync"
	"time"
)

var once sync.Once

func GetRand() (val int) {
	//第一次调用的时候启动随机数种子
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})
	val = rand.Int() % 16
	return val
}
