package database

import "github.com/adityasunny1189/FitnFine/models"

func CreateWallet(userId int) {
	wallet := models.Wallet{
		UserId: userId,
		Balance: 100,
	}
	DB.Create(&wallet)
}

func findWallet(userId int) *models.Wallet {
	var wallet models.Wallet
	DB.Find(&wallet, "user_id = ?", userId)
	return &wallet
}

func TopupWallet(userId int, amount float64) {
	wallet := findWallet(userId)
	newBalance := wallet.Balance + amount
	DB.Model(&models.Wallet{}).Where("user_id = ?", userId).Update("balance", newBalance)
}

func CreditWallet(userId int, reward float64) {
	wallet := findWallet(userId)
	newBalance := wallet.Balance + reward
	DB.Model(&models.Wallet{}).Where("user_id = ?", userId).Update("balance", newBalance)
}

func DebitWallet(userId int, penalty float64) {
	wallet := findWallet(userId)
	newBalance := wallet.Balance - penalty
	DB.Model(&models.Wallet{}).Where("user_id = ?", userId).Update("balance", newBalance)
}

func CheckWalletBalance(userId int, totalPenalty float64) bool {
	wallet := findWallet(userId)
	return wallet.Balance > totalPenalty
}

func GetWalletBalance(userId int) float64 {
	wallet := findWallet(userId)
	return wallet.Balance
}