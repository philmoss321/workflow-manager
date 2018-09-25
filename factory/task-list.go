package factory

import core "stash.mtvi.com/scm/ms/hls-packager-service/core"
import module "stash.mtvi.com/scm/ms/hls-packager-service/module"

// TaskList : define our task list
type TaskList struct {
	WorkflowType PremadeWorkflow
	Tasks        []*core.NamedTask
}

// PremadeWorkflows : list of premade workflows
var PremadeWorkflows = []TaskList{
	{
		"PackageHLS",
		[]*core.NamedTask{
			{
				Name:      "CreateAssetDirectories",
				Task:      &module.CreateAssetDirectories{Name: "CreateAssetDirectories"},
				StepIndex: 1,
			},
			{
				Name:      "DownloadAsset",
				Task:      &module.DownloadAsset{Name: "DownloadAsset"},
				StepIndex: 2,
			},
		},
	},
}
