package data

import (
	"global.p2p.api/gp2p/database"
)


type AgentStoreData struct {
	ID       string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Nickname string `bson:"nickname"`
	Email    string `bson:"email"`
	Phone    string `bson:"phone"`
	Password string `bson:"password"`
	Country string `bson:"country"`
}

type AgentStore struct {
	database database.Database
}

func NewAgentStore(db database.Database) *AgentStore {
	return &AgentStore{
		database: db,
	}
}

func (store AgentStore) AddAgent(data *AgentStoreData) (*AgentStoreData, error) {
	var agent = &AgentStoreData{}
	cursor, err := store.database.Insert(data)
	if err != nil {
		return nil, err
	}

	err = cursor.Decode(agent)
	if err != nil {
		return nil, err
	}

	return agent, err
}

