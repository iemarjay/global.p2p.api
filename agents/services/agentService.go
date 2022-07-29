package services

import (
	"global.p2p.api/agents/adapters/messages"
	"global.p2p.api/agents/data"
	"global.p2p.api/gp2p/notification"
)

type AgentService struct {
	database       *data.AgentStore
	messageSender  notification.ContractMessageSender
	welcomeMessage *messages.WelcomeMessage
}

func NewAgentService(ds *data.AgentStore, sender notification.ContractMessageSender, wm *messages.WelcomeMessage) *AgentService {
	return &AgentService{
		database:       ds,
		messageSender:  sender,
		welcomeMessage: wm,
	}
}

func (as *AgentService) RegisterAgent(agentData *AgentServiceData) (*AgentServiceData, error) {
	err := agentData.HashPassword()
	if err != nil {
		return nil, err
	}

	var agent *AgentServiceData
	agent, err = as.addAgent(agentData)
	if err != nil {
		return nil, err
	}

	as.welcomeMessage.SetTo(agent)
	as.messageSender.Send(as.welcomeMessage)
	return agent, nil
}

func (as *AgentService) addAgent(agentData *AgentServiceData) (*AgentServiceData, error) {
	storeData := agentData.toAgentStoreData()
	storeAgent, err := as.database.AddAgent(storeData)
	if err != nil {
		return nil, err
	}

	return newAgentServiceDataFromAgentStoreData(storeAgent), nil
}
