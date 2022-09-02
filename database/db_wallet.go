package database

import (
	"errors"

	"github.com/adityasunny1189/FitnFine/models"
)

func CreateWallet(userId int) error {
	wallet := models.Wallet{
		UserId:  userId,
		Balance: 100,
	}
	if err := DB.Create(&wallet).Error; err != nil {
		return err
	}
	return nil
}

func findWallet(userId int) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := DB.Find(&wallet, "user_id = ?", userId).Error; err != nil {
		return &models.Wallet{}, err
	}
	return &wallet, nil
}

func updateWallet(userId int, balance float64) error {
	err := DB.Model(&models.Wallet{}).Where("user_id = ?", userId).Update("balance", balance).Error
	if err != nil {
		return err
	}
	return nil
}

func TopupWallet(userId int, amount float64) error {
	wallet, err := findWallet(userId)
	if err != nil {
		return errors.New("find waller error ")
	}
	newBalance := wallet.Balance + amount
	updateWallet(userId, newBalance)
	return nil
}

func CreditWallet(userId int, reward float64) error {
	wallet, err := findWallet(userId)
	if err != nil {
		return errors.New("find waller error ")
	}
	newBalance := wallet.Balance + reward
	updateWallet(userId, newBalance)
	return nil
}

func DebitWallet(userId int, penalty float64) error {
	wallet, err := findWallet(userId)
	if err != nil {
		return errors.New("find waller error ")
	}
	newBalance := wallet.Balance - penalty
	updateWallet(userId, newBalance)
	return nil
}

func CheckWalletBalance(userId int, totalPenalty float64) (bool, error) {
	wallet, err := findWallet(userId)
	if err != nil {
		return false, errors.New("find wallet error ")
	}
	return wallet.Balance > totalPenalty, err
}

func GetWalletBalance(userId int) (float64, error) {
	wallet, err := findWallet(userId)
	if err != nil {
		return 0, errors.New("find wallet error ")
	}
	return wallet.Balance, nil
}
