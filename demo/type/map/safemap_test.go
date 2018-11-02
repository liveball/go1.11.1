package mapx

import (
	"fmt"
	"sync"
	"testing"
)

type SafeMap struct {
	sync.RWMutex
	Map map[int]int
}

func Test_safeMap(t *testing.T) {
	const workers = 3
	var wg sync.WaitGroup
	wg.Add(workers)
	safeMap := new(SafeMap)
	safeMap.Map = make(map[int]int)
	for i := 0; i < workers; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < i; j++ {
				safeMap.writeMap(i, i)
				safeMap.readMap(i)
			}
		}(i)
	}
	wg.Wait()
}

func (sm *SafeMap) readMap(key int) int {
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()
	fmt.Println("key=", key, "value=", value)
	return value
}

func (sm *SafeMap) writeMap(key int, value int) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}
