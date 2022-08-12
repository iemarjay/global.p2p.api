package services

import (
	"global.p2p.api/agents/adapters/messages"
	"global.p2p.api/agents/dtos"
	"global.p2p.api/agents/helpers"
	"global.p2p.api/agents/repositories"
)

type (
	AgentVerification struct {
		database            *repositories.AgentStore
		verificationMessage  *messages.Verification
		otp                 *helpers.OtpGenerator
	}

	verificationError struct {
		Email string `json:"email,omitempty"`
		Phone string `json:"phone,omitempty"`
	}
)

func NewAgentVerificationService(ds *repositories.AgentStore, vm *messages.Verification,
	otp *helpers.OtpGenerator) *AgentVerification {
	return &AgentVerification{
		database:            ds,
		verificationMessage:  vm,
		otp:                 otp,
	}
}

func (avs *AgentVerification) SendAgentVerificationCode(input *dtos.AgentVerificationDTO) error {
	sa, err := avs.database.Find(input.ID)
	if err != nil {
		return err
	}

	agent := dtos.NewAgentServiceDataFromAgentStoreData(sa)

	avs.verificationMessage.
		SetTo(agent).
		SendVia(input.Field)

	return nil
}

func (avs AgentVerification) VerifyAgent(input *dtos.AgentVerificationDTO) (error, *verificationError) {
	sa, err := avs.database.Find(input.ID)
	if err != nil {
		return err, nil
	}

	agent := dtos.NewAgentServiceDataFromAgentStoreData(sa)
	ve := &verificationError{
		Email: avs.validateMailToken(agent, input.Email),
		Phone: avs.validatePhoneToken(input.Phone, agent),
	}

	if ve.Phone == "failed" || ve.Email == "failed" {
		return nil, ve
	}

	return nil, nil
}

func (avs AgentVerification) validatePhoneToken(phone string, agent *dtos.AgentDto) string {
	var msg string
	if phone == "" {
		msg = ""
	}
	phonePassed := avs.otp.Validate(agent.OtpKeyForPhone(), phone)
	if phonePassed {
		msg = "success"
	} else {
		msg = "failed"
	}
	return msg
}

func (avs AgentVerification) validateMailToken(agent *dtos.AgentDto, email string) string {
	if email == "" {
		return ""
	}

	emailPassed := avs.otp.Validate(agent.OtpKeyForEmail(), email)
	var msg string
	if emailPassed {
		msg = "success"
	} else {
		msg = "failed"
	}

	return msg
}

func (ve verificationError) Error() string {
	var msg string
	if ve.Email != "" && ve.Phone != "" {
		msg = "Email and Phone verification failed"
	} else if ve.Phone != "" {
		msg = "Phone verification failed"
	} else if ve.Email != "" {
		msg = "Email verification failed"
	}

	return msg
}
