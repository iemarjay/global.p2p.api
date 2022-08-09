package services

import (
	"global.p2p.api/agents/adapters/messages"
	"global.p2p.api/agents/dtos"
	"global.p2p.api/agents/repositories"
	"global.p2p.api/app/notification"
)

type AgentService struct {
	database       *repositories.AgentStore
	messageSender  notification.ContractMessageSender
	welcomeMessage *messages.Welcome
}

func NewAgentService(ds *repositories.AgentStore, sender notification.ContractMessageSender, wm *messages.Welcome) *AgentService {
	return &AgentService{
		database:       ds,
		messageSender:  sender,
		welcomeMessage: wm,
	}
}


func (as *AgentService) RegisterAgent(agentData *dtos.AgentDto) (*dtos.AgentDto, error) {
	err := agentData.HashPassword()
	if err != nil {
		return nil, err
	}

	var agent *dtos.AgentDto
	agent, err = as.addAgent(agentData)
	if err != nil {
		return nil, err
	}

	as.welcomeMessage.SetTo(agent)
	as.messageSender.Send(as.welcomeMessage)
	return agent, nil
}

func (as *AgentService) addAgent(agentData *dtos.AgentDto) (*dtos.AgentDto, error) {
	storeData := agentData.ToAgentStoreData()
	storeAgent, err := as.database.AddAgent(storeData)
	if err != nil {
		return nil, err
	}

	return dtos.NewAgentServiceDataFromAgentStoreData(storeAgent), nil
}
