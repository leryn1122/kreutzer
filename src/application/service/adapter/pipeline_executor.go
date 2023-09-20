package adapter

import (
	"github.com/leryn1122/kreutzer/v2/application/service"
	"github.com/leryn1122/kreutzer/v2/infra/actions"
)

var pipelineExecutorService *PipelineExecutorServiceImpl

type PipelineExecutorServiceImpl struct {
}

func NewPipelineExecutorService() service.PipelineExecutorService {
	if pipelineExecutorService != nil {
		return pipelineExecutorService
	}
	pipelineExecutorService = &PipelineExecutorServiceImpl{}
	return pipelineExecutorService
}

func (s PipelineExecutorServiceImpl) InterpretWorkFlow(workflow *actions.Workflow) error {
	//pipeline := tekton.Pipeline{
	//  ObjectMeta: v1.ObjectMeta{
	//    Name: "",
	//    Labels: map[string]string{
	//      "tekton.kreutzer.io/origin-repo":     "leryn1122/vsp",
	//      "tekton.kreutzer.io/origin-workflow": workflow.Name,
	//    },
	//  },
	//  Spec: tekton.PipelineSpec{},
	//}
	//
	//for name, job := range workflow.Jobs {
	//
	//}
	return nil
}
