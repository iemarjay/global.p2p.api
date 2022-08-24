package dtos

import (
	"github.com/go-playground/validator/v10"
	"mime/multipart"
)

type AgentRegisterRequest struct {
	Name     string `validate:"required,alphaunicode" query:"name"`
	Nickname string `validate:"required,alphanumunicode" query:"nickname"`
	Email    string `validate:"required,email" query:"email"`
	Country  string `validate:"required,iso3166_1_alpha2" query:"country"`
	Password string `validate:"required,alphanumunicode" query:"password"`
	Phone    string `validate:"required,e164" query:"phone"`
}

type AgentLoginRequest struct {
	Identifier string `validate:"required" query:"identifier"`
	Password   string `validate:"required" query:"password"`
}

type KycRequest struct {
	Street       string                `json:"street" form:"street"`
	City         string                `json:"city" form:"city"`
	State        string                `json:"state" form:"state"`
	Country      string                `json:"country" form:"country"`
	IdCard       *multipart.FileHeader `json:"id_card"`
	AddressProof *multipart.FileHeader `json:"address_proof"`
}

func (a *AgentRegisterRequest) ToAgentServiceData() *AgentDto {
	return &AgentDto{
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

func (r *AgentLoginRequest) ToAgentServiceData() *AgentDto {
	return &AgentDto{
		Email:    r.Identifier,
		Phone:    r.Identifier,
		Password: r.Password,
	}
}

func (r KycRequest) ToKycDto() *KycDto {
	return &KycDto{
		Street: r.Street,
		City: r.City,
		State: r.State,
		Country: r.Country,
		IdCard:       r.IdCard,
		AddressProof: r.AddressProof,
	}
}
