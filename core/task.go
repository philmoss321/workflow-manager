package core

import (
	config "stash.mtvi.com/scm/ms/hls-packager-service/config"
)

// Task : represents task interface of workflow.
type Task interface {
	Execute(adapter interface{}) error
}

// NamedTask : Individual task
type NamedTask struct {
	Name      string
	Status    string
	Task      Task
	StepIndex int
	Adapter   interface{}
}

// Execute : implement Task.Execute.
func (nt *NamedTask) Execute() error {
	nt.Status = config.InProgress
	if err := nt.Task.Execute(nt.Adapter); err != nil {
		nt.Status = config.Errored
		return err
	}
	nt.Status = config.Complete
	return nil
}
