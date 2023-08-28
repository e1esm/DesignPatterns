package main

import "fmt"

type SMS struct {
	Otp
}

func (s *SMS) getRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS generating random OTP: %s", randomOTP)
	return randomOTP
}

func (s *SMS) saveOTPCache(otp string) {
	fmt.Printf("Saving otp: %s to the cache", otp)
}

func (s *SMS) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s", message)
	return nil
}
func (s *SMS) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}
