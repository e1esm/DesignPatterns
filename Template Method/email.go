package main

import "fmt"

type Email struct {
	Otp
}

func (e *Email) getRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random OTP: %s", randomOTP)
	return randomOTP
}

func (e *Email) saveOTPCache(otp string) {
	fmt.Printf("Saving otp: %s to the cache", otp)
}

func (e *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending sms: %s", message)
	return nil
}
func (e *Email) getMessage(otp string) string {
	return "EMAIL: OTP for login is " + otp
}
