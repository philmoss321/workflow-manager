package model

import workflow "stash.mtvi.com/scm/ms/hls-packager-service/core"

// Task : represents task interface of workflow.
type Task interface {
	Execute() error
}

type db interface {
	Ping() error
	Insert(workflow *workflow.Workflow, collection string)
	Update(namespace string, workflow *workflow.Workflow, collection string)
	UpdateStatus(namespace string, status string, collection string)
	FormStatusReturn(DBData DBMultiReturn) StatusReturn
}

type queue interface {
	PollQueue()
	GetQueueAttributes() error
}

type status interface {
	CheckQueueHealth(qModel *QueueModel) (int64, error)
	CheckDbHealth(dbModel *DbModel) (int64, error)
	FormStatusJSON(cacheError string, sqsError string, mongoError string)
}
