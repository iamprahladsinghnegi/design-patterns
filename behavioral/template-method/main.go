package main

import (
	"fmt"

	tm "github.com/iamprahladsinghnegi/design-patterns/behavioral/template-method/template-method"
)

func main() {
	smsOTP := &tm.Sms{}
	o := tm.BaseOtp{
		Otp: smsOTP,
	}
	o.GenAndSendOTP(4)

	fmt.Println("")
	emailOTP := &tm.Email{}
	o = tm.BaseOtp{
		Otp: emailOTP,
	}
	o.GenAndSendOTP(4)

}
