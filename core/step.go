package core

import (
	"fmt"
	"sync"

	multierror "github.com/hashicorp/go-multierror"
	config "github.com/philmoss321/workflow-manager/config"
)

// Step : represents parallel tasks on workflow.
type Step struct {
	Tasks  []*NamedTask
	status string
	index  int
}

// StepFactory : Make new step
func StepFactory(index int) func(tasks []*NamedTask) *Step {
	return func(tasks []*NamedTask) *Step {
		step := NewStep(index)

		for _, task := range tasks {
			step.AddTask(task.Name, task.Task, task.Adapter)
		}

		return step
	}
}

// NewStep : creates a parallel task by task list.
func NewStep(index int) *Step {
	return &Step{
		Tasks:  make([]*NamedTask, 0),
		status: config.NotStarted,
		index:  index,
	}
}

// AddTask add parallel task with name
func (step *Step) AddTask(name string, task Task, adapter interface{}) {
	step.Tasks = append(step.Tasks, &NamedTask{Name: name, Task: task, Status: config.NotStarted, Adapter: adapter})
}

// Run : Run the step!
func (step *Step) Run() error {
	step.status = config.InProgress
	return step.run(step.Tasks)
}

func (step *Step) run(nt []*NamedTask) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(nt))
	for i, t := range nt {
		wg.Add(1)
		Logger.Print(fmt.Sprintf("workflow: Start task: %v", nt[i].Name))
		go func(namedTask *NamedTask) {
			if err := namedTask.Execute(); err != nil {
				Logger.Print(err)
				errChan <- err
				Logger.Print(fmt.Sprintf("workflow: Errored task: %v", namedTask.Name))
			} else {
				Logger.Print(fmt.Sprintf("workflow: Completed task: %v", namedTask.Name))
			}
			wg.Done()
		}(t)
	}
	resultChan := make(chan error)
	go func() {
		var result *multierror.Error
		for err := range errChan {
			step.status = config.Errored
			result = multierror.Append(result, err)
		}
		resultChan <- result.ErrorOrNil()
	}()
	wg.Wait()
	close(errChan)
	err := <-resultChan
	if err == nil {
		step.status = config.Complete
	} else {
		step.status = config.Errored
	}
	return err
}
