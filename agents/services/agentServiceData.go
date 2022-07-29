package services

import (
	"global.p2p.api/agents/data"
	"global.p2p.api/gp2p/hash"
)

type AgentServiceData struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password,omitempty"`
	Country  string `json:"country"`
}

func (asd *AgentServiceData) RouteForMail() string {
	return asd.Email
}

func (asd *AgentServiceData) toAgentStoreData() *data.AgentStoreData {
	return &data.AgentStoreData{
		Name:     asd.Name,
		Nickname: asd.Nickname,
		Email:    asd.Email,
		Phone:    asd.Phone,
		Password: asd.Password,
		Country:  asd.Country,
	}
}

func (asd *AgentServiceData) ComparePassword(password string) bool {
	return hash.Compare(password, asd.Password)
}

func (asd *AgentServiceData) HashPassword() error {
	var err error
	asd.Password, err = hash.Generate(asd.Password)

	return err
}

func newAgentServiceDataFromAgentStoreData(storeData*data.AgentStoreData) *AgentServiceData {
	return &AgentServiceData{
		ID:       storeData.ID,
		Name:     storeData.Name,
		Nickname: storeData.Nickname,
		Email:    storeData.Email,
		Phone:    storeData.Phone,
		Password: storeData.Password,
		Country:  storeData.Country,
	}
}
