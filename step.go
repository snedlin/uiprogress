package uiprogress

const (
	// Waiting is an indicator for a step that has no outcome yet
	Waiting StepStatus = iota

	// Success is an indicator for a successful step
	Success

	// Failure is an indicator for a failed step
	Failure
)

// StepStatus is an enumeration of statuses for steps
type StepStatus int

// Step represents a single step that makes up a unit of progress
type Step struct {
	Name   string
	Status StepStatus
}
