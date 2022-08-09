package repositories

import (
	"global.p2p.api/app/database"
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

func (asd AgentStoreData) toEmailOrPhoneFilter() bson.M {
	return bson.M{"$or": []bson.M{{"email": asd.Email}, {"phone": asd.Phone}}}
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

func (s AgentStore) FindAgentByEmailOrPhone(data *AgentStoreData) (*AgentStoreData, error) {
	var agent = &AgentStoreData{}
	filter := data.toEmailOrPhoneFilter()
	cursor := s.database.FindOne(filter)

	err := cursor.Decode(agent)
	if err != nil {
		return nil, err
	}

	return agent, err
}

func (s AgentStore) Find(id string) (*AgentStoreData, error) {
	var agent = &AgentStoreData{}
	err := s.database.FindOneByID(id).Decode(agent)
	if err != nil {
		return nil, err
	}

	return agent, nil
}
