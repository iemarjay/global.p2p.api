package data

import (
	"global.p2p.api/gp2p/database"
	"go.mongodb.org/mongo-driver/bson"
)

type AgentStoreData struct {
	ID       string `bson:"_id,omitempty"`
	Name     string `bson:"name,omitempty"`
	Nickname string `bson:"nickname,omitempty"`
	Email    string `bson:"email,omitempty"`
	Phone    string `bson:"phone,omitempty"`
	Password string `bson:"password,omitempty"`
	Country  string `bson:"country,omitempty"`
}

type AgentStore struct {
	database database.Database
}

func NewAgentStore(db database.Database) *AgentStore {
	return &AgentStore{
		database: db,
	}
}

func (s AgentStore) AddAgent(data *AgentStoreData) (*AgentStoreData, error) {
	var agent = &AgentStoreData{}
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

func (s AgentStore) FindAgentByEmailOrPhone(identifier string) (*AgentStoreData, error) {
	var agent = &AgentStoreData{}
	filter := s.whereIdentifierIsEmailOrPhoneFilter(identifier)
	cursor := s.database.FindOne(filter)

	err := cursor.Decode(agent)
	if err != nil {
		return nil, err
	}

	return agent, err
}

func (s AgentStore) whereIdentifierIsEmailOrPhoneFilter(identifier string) bson.M {
	return bson.M{"$or": []bson.M{{"email": identifier}, {"phone": identifier}}}
}
