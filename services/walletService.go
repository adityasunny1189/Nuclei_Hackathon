package services

import "github.com/adityasunny1189/FitnFine/database"

func AddRewardService(userId int, amount float64) error {
	return database.CreditWallet(userId, amount)
}

func AddPenaltyService(userId int, amount float64) error {
	return database.DebitWallet(userId, amount)
}

func TopupWalletService(userId int, amount float64) error {
	return database.TopupWallet(userId, amount)
}