package model

// DbModel : implements db functions
type DbModel struct {
	db
}

// QueueModel : implements queue functions
type QueueModel struct {
	queue
}

// StatusModel : implements status functions
type StatusModel struct {
	status
}

// NewDb : creates new DB Model
func NewDb(db db) *DbModel {
	return &DbModel{
		db: db,
	}
}

// NewQueue : creates new queue Model
func NewQueue(queue queue) *QueueModel {
	return &QueueModel{
		queue: queue,
	}
}

// NewStatus : creates status model
func NewStatus(status status) *StatusModel {
	return &StatusModel{
		status: status,
	}
}
