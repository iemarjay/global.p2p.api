package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"global.p2p.api/agents/dtos"
	"global.p2p.api/agents/repositories"
	"global.p2p.api/app/http"
	"time"
)

type LoginServiceData struct {
	Token string         `json:"token"`
	User  *dtos.AgentDto `json:"user"`
}

type LoginService struct {
	store *repositories.AgentStore
}

func NewLoginService(database *repositories.AgentStore) *LoginService {
	return &LoginService{store: database}
}

func (s LoginService) Login(data *dtos.AgentDto) (*LoginServiceData, error) {
	filter := data.ToFindByEmailOrPhoneFilter()

	sdAgent, err := s.store.FindAgentByEmailOrPhone(filter)
	if err != nil {
		return nil, err
	}

	agent := dtos.NewAgentServiceDataFromAgentStoreData(sdAgent)
	if agent.ComparePassword(data.Password) == false {
		return nil, errors.New("credential doesn't match record")
	}

	c := &http.JwtClaim{
		Id: agent.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	t, err := c.ToSignJwtString()

	if err != nil {
		return nil, err
	}

	l := &LoginServiceData{
		Token: t,
		User:  agent,
	}

	return l, err
}