package core

import (
	"fmt"
	"sort"

	"log"
	"os"

	config "stash.mtvi.com/scm/ms/hls-packager-service/config"
)

// Workflow : contains tasks list of workflow definition.
type Workflow struct {
	steps  []*Step
	status string
	logger *log.Logger
}

// Logger : workflow logger
var Logger = log.New(os.Stdout, "[packager] ", log.Ldate|log.Ltime|log.Lshortfile)

// ByIndex : sort by step index
type ByIndex []*Step

func (a ByIndex) Len() int           { return len(a) }
func (a ByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIndex) Less(i, j int) bool { return a[i].index < a[j].index }

// NewWorkflow : creates a new workflow definition.
func NewWorkflow() *Workflow {
	return &Workflow{
		steps:  make([]*Step, 0),
		status: config.NotStarted,
	}
}

// SetLogger : sets log writer.
func (wf *Workflow) SetLogger(logger *log.Logger) {
	wf.logger = logger
}

// AddStep : add step with name.
func (wf *Workflow) AddStep(step *Step) {
	wf.steps = append(wf.steps, step)
}

// Execute : begin workflow
func (wf *Workflow) Execute() error {
	wf.status = config.InProgress
	return wf.Run()
}

// Run : defined workflow tasks.
func (wf *Workflow) Run() error {
	return wf.run(wf.steps)
}

func (wf *Workflow) run(steps []*Step) error {
	sort.Sort(ByIndex(steps))
	for i, s := range steps {
		wf.logger.Print(fmt.Sprintf("workflow: Start step: %v", steps[i].name))
		if err := s.Run(); err != nil {
			wf.status = config.Errored
			return err
		}
		wf.logger.Print(fmt.Sprintf("workflow: Complete step: %v", steps[i].name))
	}
	wf.status = config.Complete
	return nil
}

// Summary : returns task flow summary.
// func (wf *Workflow) Summary() string {
// 	return buildTaskSummary(wf.tasks, " -> ", true)
// }

// func buildTaskSummary(tasks []*namedTask, delimiter string, showNumber bool) string {
// 	names := make([]string, len(tasks))
// 	for i, t := range tasks {
// 		var number string
// 		if showNumber {
// 			number = fmt.Sprintf("%d.", i+1)
// 		}
// 		if w, ok := t.task.(*Workflow); ok {
// 			names[i] = fmt.Sprintf("%s%s<Workflow>(%s)", number, t.name, w.Summary())
// 		} else if pt, ok := t.task.(*ParallelTask); ok {
// 			names[i] = fmt.Sprintf("%s%s<ParallelTask>(%s)", number, t.name, pt.Summary())
// 		} else {

// 			names[i] = fmt.Sprintf("%s%s<%s>", number, t.name, nameOfTask(t.task))
// 		}
// 	}
// 	return strings.Join(names, delimiter)
// }

// func nameOfTask(task Task) string {
// 	t := reflect.TypeOf(task)
// 	if t.Kind() == reflect.Ptr {
// 		return t.Elem().Name()
// 	}
// 	return t.Name()
// }
