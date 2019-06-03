package pool
import(
	"testing"
	"fmt"
	"sync"
)

func TestPool(t *testing.T) {
	p := New(10)
	taskCount := 50
	var wg sync.WaitGroup
	wg.Add(taskCount)
	for i := 0; i < taskCount; i++{
		data := i
		t := func() {
				for dd := 0; dd  < data; dd++ {
				fmt.Println(dd)
			}
			wg.Done()
		}
		p.Schedule(t)
	}

	wg.Wait()
	fmt.Println("All DONE")
}
