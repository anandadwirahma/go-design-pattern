package facade

import "fmt"

type WalletFacade struct {
	Account      *Account
	Wallet       *Wallet
	SecurityCode *SecurityCode
	Notification *Notification
	Ledger       *Ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting Create Account")
	walletFacade := &WalletFacade{
		Account:      NewAccount(accountID),
		Wallet:       NewWallet(),
		SecurityCode: NewSecurityCode(code),
		Notification: &Notification{},
		Ledger:       &Ledger{},
	}
	fmt.Println("Account Created")
	return walletFacade
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting Add Money to Wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.Wallet.CreditBalance(amount)
	w.Notification.SendWalletCreditNotification()
	w.Ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyForWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.Wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.Notification.SendWalletDebitNotification()
	w.Ledger.MakeEntry(accountID, "debit", amount)
	return nil
}
