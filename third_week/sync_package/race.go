package syncpackage

//go test -race
import (
	"fmt"
	"sync"
	"time"
)

//Wait waitgroup
var Wait sync.WaitGroup

//Counter counter
var Counter int = 0

// var test interface{}

func race() {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go routineFunc(routine)
	}
	Wait.Wait()
	fmt.Printf("Final Counter:%d", Counter)
}

func routineFunc(id int) {
	for count := 0; count < 2; count++ {
		value := Counter
		time.Sleep(1 * time.Nanosecond)
		value++
		Counter = value
	}
	Wait.Done()
}
