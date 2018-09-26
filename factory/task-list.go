package factory

import (
	adapter "github.com/philmoss321/workflow-manager/adapter"
	core "github.com/philmoss321/workflow-manager/core"
	module "github.com/philmoss321/workflow-manager/module"
)

// TaskList : define our task list
type TaskList struct {
	WorkflowType PremadeWorkflow
	Tasks        []*core.NamedTask
}

var downloadAdapter = &adapter.HTTPRequest{Name: "HTTP"}

// PremadeWorkflows : list of premade workflows
var PremadeWorkflows = []TaskList{
	{
		"PackageHLS",
		[]*core.NamedTask{
			{
				Name:      "CreateAssetDirectories",
				Task:      &module.CreateAssetDirectories{Name: "CreateAssetDirectories"},
				StepIndex: 1,
				Adapter:   downloadAdapter,
			},
			{
				Name:      "DownloadAsset",
				Task:      &module.DownloadAsset{Name: "DownloadAsset"},
				StepIndex: 1,
				Adapter:   downloadAdapter,
			},
		},
	},
}
