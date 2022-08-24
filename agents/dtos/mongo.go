package dtos

import (
	"github.com/naamancurtis/mongo-go-struct-to-bson/mapper"
	"go.mongodb.org/mongo-driver/bson"
)

type AgentMongoDto struct {
	ID       string      `bson:"_id,omitempty"`
	Name     string      `bson:"name,omitempty"`
	Nickname string      `bson:"nickname,omitempty"`
	Email    string      `bson:"email,omitempty"`
	Phone    string      `bson:"phone,omitempty"`
	Password string      `bson:"password,omitempty"`
	Country  string      `bson:"country,omitempty"`
	Kyc      KycMongoDto `bson:"kyc"`
}

func (asd AgentMongoDto) ToEmailOrPhoneFilter() bson.M {
	return bson.M{"$or": []bson.M{{"email": asd.Email}, {"phone": asd.Phone}}}
}

func (asd AgentMongoDto) ToAgentDto() *AgentDto {
	a := &AgentDto{
		ID:       asd.ID,
		Name:     asd.Name,
		Nickname: asd.Nickname,
		Email:    asd.Email,
		Phone:    asd.Phone,
		Password: asd.Password,
		Country:  asd.Country,
	}

	if asd.Kyc != (KycMongoDto{}) {
		a.Kyc = asd.Kyc.ToKycDto()
	}

	return a
}

type KycMongoDto struct {
	Street           string `bson:"street,omitempty"`
	City             string `bson:"city,omitempty"`
	State            string `bson:"state,omitempty"`
	Country          string `bson:"country,omitempty"`
	AddressProofPath string `bson:"address_proof_path,omitempty"`
	IdCardPath       string `bson:"id_card_path,omitempty"`
}

func (d *KycMongoDto) ToBson() (data bson.M) {
	return mapper.ConvertStructToBSONMap(d, nil)
}

func (d *KycMongoDto) ToKycDto() *KycDto {
	return &KycDto{
		Street:           d.Street,
		City:             d.City,
		State:            d.State,
		Country:          d.Country,
		AddressProofPath: d.AddressProofPath,
		IdCardPath:       d.IdCardPath,
	}
}
