package main

import "fmt"

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	SecurityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Creating new account")
	walletFacade := &WalletFacade{
		account:      NewAccount(accountID),
		wallet:       NewWallet(),
		SecurityCode: NewSecurityCode(code),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
	fmt.Println("Created new account")
	return walletFacade
}

func (wf *WalletFacade) addMoneyToWallet(accountID string, securityCode, amount int) error {
	fmt.Println("Starting to add money to the account")
	err := wf.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = wf.SecurityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	wf.wallet.creditBalance(amount)
	wf.notification.sendWalletCreditNotification()
	wf.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (wf *WalletFacade) deductMoneyFromWallet(accountID string, securityCode, amount int) error {
	fmt.Println("Starting to debit money from account")
	if err := wf.account.checkAccount(accountID); err != nil {
		return err
	}
	if err := wf.SecurityCode.checkCode(securityCode); err != nil {
		return err
	}
	if err := wf.wallet.debitBalance(amount); err != nil {
		return err
	}
	wf.notification.sendWalletDebitNotification()
	wf.ledger.makeEntry(accountID, "debit", amount)
	return nil
}
