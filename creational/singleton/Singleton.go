package singlton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type config struct{}

var configInstance *config

func GetInstance() *config {
	if configInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if configInstance == nil {
			configInstance = new(config)
		}
	}

	return configInstance
}

func PrintInstance(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\nInstance %v := %p", n, GetInstance())
}
