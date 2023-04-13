package domain

type Workflow struct {
	Name         string
	WorkflowType string
	Processes    []Process
}

type Process struct {
	Name        string
	ProcessType string
	Image       string
	Src         string
}
