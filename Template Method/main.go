package main

func main() {
	smsOTP := &SMS{}
	o := &Otp{
		iOtp: smsOTP,
	}
	o.getAndSendOTP(4)

	emailOTP := &Email{}
	o = &Otp{
		iOtp: emailOTP,
	}
	o.getAndSendOTP(4)
}
