package factory

import (
	core "github.com/philmoss321/workflow-manager/core"
)

// PremadeWorkflow : constants for premade workflow names
type PremadeWorkflow string

const (
	// PackageHLS : premade hls workflow
	PackageHLS PremadeWorkflow = "PackageHLS"
	// EncodeMP4 : premade encode workflow
	EncodeMP4 PremadeWorkflow = "EncodeMP4"
)

// CreatePremadeWorkflow : generage premade workflows
func CreatePremadeWorkflow(wfType PremadeWorkflow) *core.Workflow {
	workflow := core.NewWorkflow()
	steps := make(map[int]*core.Step)
	for _, workflowList := range PremadeWorkflows {
		if workflowList.WorkflowType == wfType {
			var step *core.Step
			for _, task := range workflowList.Tasks {
				if steps[task.StepIndex] == nil {
					step = core.NewStep(task.StepIndex)
					step.AddTask(task.Name, task.Task, task.Adapter)
					steps[task.StepIndex] = step
				} else {
					steps[task.StepIndex].AddTask(task.Name, task.Task, task.Adapter)
				}
			}
		} else {
			return nil
		}
	}

	for _, v := range steps {
		workflow.AddStep(v)
	}

	return workflow
}
