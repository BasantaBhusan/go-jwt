package models

import "gorm.io/gorm"

type Kyc struct {
	gorm.Model
	UserID         uint `gorm:"not null"`
	FullName       string
	MobileNumber   string
	FirmRegistered bool
	Address        Address     `gorm:"foreignKey:KycID"`
	WorkingArea    WorkingArea `gorm:"foreignKey:KycID"`
	Service        Service     `gorm:"foreignKey:KycID"`
}

type Address struct {
	gorm.Model
	UserID       uint
	KycID        uint
	Province     string
	District     string
	Municipality string
	WardNumber   string
	Latitude     string
	Longitude    string
}

type WorkingArea struct {
	gorm.Model
	UserID     uint `gorm:"not null"`
	KycID      uint
	AreaName   string
	Activities []Activity
}

type Activity struct {
	gorm.Model
	WorkingAreaID uint
	KycID         uint
	ActivityName  string
	Items         []ActivityItem
}

type ActivityItem struct {
	gorm.Model
	ActivityID uint
	Name       string
}

type ServiceType string

const (
	ExpertAdvice          ServiceType = "Expert Advice"
	BusinessPartnership   ServiceType = "Business Partnership"
	BankLoanFacilitation  ServiceType = "Bank Loan Facilitation"
	TrainingAndCoaching   ServiceType = "Training and Coaching"
	ColdStoreConstruction ServiceType = "Cold Store Construction"
	AssistanceInMarketing ServiceType = "Assistance in Marketing"
	InvestmentService     ServiceType = "Investment"
)

type Service struct {
	gorm.Model
	UserID      uint `gorm:"not null"`
	KycID       uint
	ServiceName ServiceType
	Investment  InvestmentOption
}

type InvestmentOption string

const (
	UpTo5LAKHS  InvestmentOption = "up to 5 Lakhs"
	UpTo10LAKHS InvestmentOption = "up to 10 Lakhs"
	UpTo25LAKHS InvestmentOption = "up to 25 Lakhs"
	UpTo50LAKHS InvestmentOption = "up to 50 Lakhs"
	UpTo1CRORE  InvestmentOption = "up to 1 Crore"
	Above1CRORE InvestmentOption = "above 1 Crore"
)
