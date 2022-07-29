package http

import (
	"github.com/go-playground/validator/v10"
	"global.p2p.api/agents/services"
)


type AgentRegisterRequest struct {
	Name string `validate:"required,alphaunicode" query:"name"`
	Nickname string `validate:"required,alphanumunicode" query:"nickname"`
	Email string `validate:"required,email" query:"email"`
	Country string `validate:"required,iso3166_1_alpha2" query:"country"`
	Password string `validate:"required,alphanumunicode" query:"password"`
	Phone string `validate:"required,e164" query:"phone"`
}

func (a *AgentRegisterRequest) toAgentServiceData() *services.AgentServiceData {
	return &services.AgentServiceData{
		Name:     a.Name,
		Nickname: a.Nickname,
		Email:    a.Email,
		Country:  a.Country,
		Password: a.Password,
		Phone:    a.Phone,
	}
}

func (a *AgentRegisterRequest) Validate() error {
	return validator.New().Struct(a)
}

