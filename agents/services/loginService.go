package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"global.p2p.api/agents/data"
)

type JwtClaim struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type LoginServiceData struct {
	Token string            `json:"token"`
	User  *AgentServiceData `json:"user"`
}

type LoginService struct {
	store *data.AgentStore
}

func NewLoginService(database *data.AgentStore) *LoginService {
	return &LoginService{store: database}
}

func (s LoginService) Login(data *AgentServiceData) (*LoginServiceData, error) {
	filter := data.toFindByEmailOrPhoneFilter()

	sdAgent, err := s.store.FindAgentByEmailOrPhone(filter.Email)
	if err != nil {
		return nil, err
	}

	agent := newAgentServiceDataFromAgentStoreData(sdAgent)
	if agent.ComparePassword(data.Password) == false {
		return nil, errors.New("credential doesn't match record")
	}

	claims := &JwtClaim{
		Id: agent.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	var t string
	t, err = token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	l := &LoginServiceData{
		Token: t,
		User:  agent,
	}

	return l, err
}