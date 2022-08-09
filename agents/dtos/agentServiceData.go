package dtos

import (
	"global.p2p.api/agents/repositories"
	"global.p2p.api/app/hash"
)

type (
	AgentDto struct {
		ID       string `json:"id,omitempty"`
		Name     string `json:"name"`
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password,omitempty"`
		Country  string `json:"country"`
	}

	AgentVerificationDTO struct {
		ID    string `json:"id,omitempty"`
		Field string `json:"field"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	}
)

func (asd *AgentDto) RouteForMail() string {
	return asd.Email
}

func (asd *AgentDto) Key() string {
	return asd.ID
}

func (asd *AgentDto) OtpKeyForEmail() string {
	return asd.ID+ ".email"
}

func (asd *AgentDto) OtpKeyForPhone() string {
	return asd.ID+ ".phone"
}

func (asd *AgentDto) ToAgentStoreData() *repositories.AgentStoreData {
	return &repositories.AgentStoreData{
		Name:     asd.Name,
		Nickname: asd.Nickname,
		Email:    asd.Email,
		Phone:    asd.Phone,
		Password: asd.Password,
		Country:  asd.Country,
	}
}

func (asd *AgentDto) ToFindByEmailOrPhoneFilter() *repositories.AgentStoreData {
	return &repositories.AgentStoreData{
		Email: asd.Email,
		Phone: asd.Phone,
	}
}

func (asd *AgentDto) ComparePassword(password string) bool {
	return hash.Compare(password, asd.Password)
}

func (asd *AgentDto) HashPassword() error {
	var err error
	asd.Password, err = hash.Generate(asd.Password)

	return err
}

func NewAgentServiceDataFromAgentStoreData(storeData *repositories.AgentStoreData) *AgentDto {
	return &AgentDto{
		ID:       storeData.ID,
		Name:     storeData.Name,
		Nickname: storeData.Nickname,
		Email:    storeData.Email,
		Phone:    storeData.Phone,
		Password: storeData.Password,
		Country:  storeData.Country,
	}
}
