package main

import "fmt"

type SecurityCode struct {
	code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (sc *SecurityCode) checkCode(incomingCode int) error {
	if sc.code == incomingCode {
		fmt.Println("Code is verified")
		return nil
	}
	return fmt.Errorf("security code is incorrect")
}
