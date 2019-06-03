package pool
import(
	"testing"
	"fmt"
	"sync"
)

func TestDispatcherPool(t *testing.T) {
	d := NewDispatcher(10)
	d.Run()
	taskCount := 50
	var wg sync.WaitGroup
	wg.Add(taskCount)
	for i := 0; i < taskCount; i++{
		data := i
		task := Task {
			Proc:func() {
				for dd := 0; dd  < data; dd++ {
				fmt.Println(dd)
			}
			wg.Done()
		}}
		d.Schedule(task)
	}

	wg.Wait()
	fmt.Println("All DONE")
}
