package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NRIC          string `json:"nric" gorm:"primaryKey"`
	WalletAddress string `json:"walletaddress"`
}
