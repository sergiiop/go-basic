package clase1

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}

func go_routines() {

	var wg sync.WaitGroup

	
	
	
	fmt.Println("Hello")
	wg.Add(1)
	go say("World", &wg)

	wg.Wait()
}
