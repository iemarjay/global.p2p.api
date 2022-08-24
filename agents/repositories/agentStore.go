package repositories

import (
	"global.p2p.api/agents/dtos"
	"global.p2p.api/app/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Kyc interface {
	UpdateKyc(agentId string, input *dtos.KycDto) (*dtos.AgentDto, error)
}

type AgentStore struct {
	database database.Database
}

func NewAgentStore(db database.Database) *AgentStore {
	return &AgentStore{
		database: db,
	}
}

func (s AgentStore) AddAgent(data *dtos.AgentMongoDto) (*dtos.AgentMongoDto, error) {
	var agent = &dtos.AgentMongoDto{}
	cursor, err := s.database.Insert(data)
	if err != nil {
		return nil, err
	}

	err = cursor.Decode(agent)
	if err != nil {
		return nil, err
	}

	return agent, err
}

func (s AgentStore) FindAgentByEmailOrPhone(data *dtos.AgentMongoDto) (*dtos.AgentMongoDto, error) {
	var agent = &dtos.AgentMongoDto{}
	filter := data.ToEmailOrPhoneFilter()
	cursor := s.database.FindOne(filter)

	err := cursor.Decode(agent)
	if err != nil {
		return nil, err
	}

	return agent, err
}

func (s AgentStore) Find(id string) (*dtos.AgentMongoDto, error) {
	var agent = &dtos.AgentMongoDto{}
	err := s.database.FindOneByID(id).Decode(agent)
	if err != nil {
		return nil, err
	}

	return agent, nil
}

func (s AgentStore) UpdateKyc(agentId string, input *dtos.KycDto) (*dtos.AgentDto, error) {
	data := input.ToKycMongoDto().ToBson()
	cursor, err := s.database.UpdateOneById(agentId, bson.M{"$set": bson.M{"kyc": data}})
	if err != nil {
		return nil, err
	}

	var agent = &dtos.AgentMongoDto{}
	err = cursor.Decode(agent)
	if err != nil {
		return nil, err
	}

	return agent.ToAgentDto(), nil
}
