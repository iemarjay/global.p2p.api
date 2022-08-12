package dtos

import (
	"global.p2p.api/app/hash"
	"mime/multipart"
)

type (
	AgentDto struct {
		ID       string  `json:"id,omitempty"`
		Name     string  `json:"name"`
		Nickname string  `json:"nickname"`
		Email    string  `json:"email"`
		Phone    string  `json:"phone"`
		Password string  `json:"password,omitempty"`
		Country  string  `json:"country"`
		Kyc      *KycDto `json:"kyc"`
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
	return asd.ID + ".email"
}

func (asd *AgentDto) OtpKeyForPhone() string {
	return asd.ID + ".phone"
}

func (asd *AgentDto) ToAgentStoreData() *AgentMongoDto {
	return &AgentMongoDto{
		Name:     asd.Name,
		Nickname: asd.Nickname,
		Email:    asd.Email,
		Phone:    asd.Phone,
		Password: asd.Password,
		Country:  asd.Country,
	}
}

func (asd *AgentDto) ToFindByEmailOrPhoneFilter() *AgentMongoDto {
	return &AgentMongoDto{
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

func NewAgentServiceDataFromAgentStoreData(storeData *AgentMongoDto) *AgentDto {
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

type KycDto struct {
	Street           string `json:"street"`
	City             string `json:"city"`
	State            string `json:"state"`
	Country          string `json:"country"`
	AddressProofPath string `json:"address_proof_path"`
	AddressProofUrl  string `json:"address_proof_url"`
	IdCardPath       string `json:"id_card_path"`
	IdCardUrl        string `json:"id_card_url"`
	IdCard           *multipart.FileHeader
	AddressProof     *multipart.FileHeader
}

func (d *KycDto) ToKycMongoDto() *KycMongoDto {
	return &KycMongoDto{
		Street:           d.Street,
		City:             d.City,
		State:            d.State,
		Country:          d.Country,
		AddressProofPath: d.AddressProofPath,
		IdCardPath:       d.IdCardPath,
	}
}

func (d *KycDto) UploadFiles(uploadFunc func(file *multipart.FileHeader) string) {
	if d.AddressProofPath == "" && d.AddressProof != nil {
		d.AddressProofPath = uploadFunc(d.AddressProof)
	}

	if d.IdCardPath == "" && d.IdCard != nil {
		d.IdCardPath = uploadFunc(d.IdCard)
	}
}

func (d *KycDto) UpdateUrlForAllFiles(fileUrlFunc func(filePath string) string) {
	if d.AddressProofUrl == "" && d.AddressProofPath != "" {
		d.AddressProofUrl = fileUrlFunc(d.AddressProofPath)
	}

	if d.IdCardUrl == "" && d.IdCardPath != "" {
		d.IdCardUrl = fileUrlFunc(d.IdCardPath)
	}
}
