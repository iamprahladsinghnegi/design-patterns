package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var singletonInstance *Singleton

var lock = &sync.Mutex{}

func GetInstance() *Singleton {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			fmt.Println("Creating single instance now.")
			singletonInstance = &Singleton{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singletonInstance
}
