package services

import (
	"global.p2p.api/agents/dtos"
	"global.p2p.api/agents/repositories"
	"global.p2p.api/app/fileStorage"
)

type AgentKyc struct {
	fileStorage *fileStorage.FileStorage
	repository  repositories.Kyc
}

func NewAgentKyc(uploader *fileStorage.FileStorage, r repositories.Kyc) *AgentKyc {
	return &AgentKyc{
		fileStorage: uploader,
		repository:  r,
	}
}

func (a *AgentKyc) UpdateKyc(agentId string, input *dtos.KycDto) (*dtos.AgentDto, error) {
	uploadFunc := a.fileStorage.Store
	input.UploadFiles(uploadFunc)

	agent, err := a.repository.UpdateKyc(agentId, input)
	if err != nil {
		return nil, err
	}

	fileUrlFunc := a.fileStorage.Url
	agent.Kyc.UpdateUrlForAllFiles(fileUrlFunc)

	return agent, nil
}

