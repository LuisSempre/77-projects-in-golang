// goroutines example
package main
import (
	"fmt"
	"time"
)
func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Main:", i)
		time.Sleep(150 * time.Millisecond)
	}

	time.Sleep(time.Second)
}