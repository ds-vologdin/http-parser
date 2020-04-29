package counter

import "sync"

type TaskCounter struct {
	task chan struct{}
	wg   sync.WaitGroup
}

func (c *TaskCounter) Inc() {
	<-c.task
	c.wg.Add(1)
}
func (c *TaskCounter) Done() {
	c.wg.Done()
	c.task <- struct{}{}
}

func (c *TaskCounter) Wait() {
	c.wg.Wait()
}

func NewTaskCounter(max int) TaskCounter {
	c := make(chan struct{}, max)
	for i := 0; i < max; i++ {
		c <- struct{}{}
	}
	return TaskCounter{task: c}
}
