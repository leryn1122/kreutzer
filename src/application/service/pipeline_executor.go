package service

import "github.com/leryn1122/kreutzer/v2/infra/actions"

type PipelineExecutorService interface {
	InterpretWorkFlow(workflow *actions.Workflow) error
}
