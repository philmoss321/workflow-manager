package core

import (
	"fmt"
	"sync"

	multierror "github.com/hashicorp/go-multierror"
	config "stash.mtvi.com/scm/ms/hls-packager-service/config"
)

// Step : represents parallel tasks on workflow.
type Step struct {
	name   string
	Tasks  []*NamedTask
	status string
	index  int
}

// StepFactory : Make new step
func StepFactory(index int) func(tasks []*NamedTask) *Step {
	return func(tasks []*NamedTask) *Step {
		step := NewStep(index)

		for _, task := range tasks {
			step.AddTask(task.Name, task.Task)
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
func (step *Step) AddTask(name string, task Task) {
	step.Tasks = append(step.Tasks, &NamedTask{Name: name, Task: task, Status: config.NotStarted})
}

// Run : Run the step!
func (step *Step) Run() error {
	step.status = config.InProgress
	return step.run(step.Tasks)
}

func (step *Step) run(nt []*NamedTask) error {
	var wg sync.WaitGroup
	errChan := make(chan error)
	for i, t := range nt {
		wg.Add(1)
		Logger.Print(fmt.Sprintf("workflow: Start step: %v", nt[i].Name))
		go func(namedTask *NamedTask) {
			if err := namedTask.Execute(); err != nil {
				errChan <- err
			}
			wg.Done()
		}(t)
		Logger.Print(fmt.Sprintf("workflow: Complete task: %v", nt[i].Name))
	}

	resultChan := make(chan error)
	go func() {
		var result *multierror.Error
		for err := range errChan {
			result = multierror.Append(result, err)
		}
		resultChan <- result.ErrorOrNil()
	}()

	wg.Wait()
	if <-resultChan == nil {
		step.status = config.Complete
	} else {
		step.status = config.Errored
	}
	close(errChan)
	return <-resultChan
}
