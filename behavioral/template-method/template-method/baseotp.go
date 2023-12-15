package templatemethod

type BaseOtp struct {
	Otp Otp
}

func (o *BaseOtp) GenAndSendOTP(otpLength int) error {
	otp := o.Otp.genRandomOTP(otpLength)
	o.Otp.saveOTPCache(otp)
	message := o.Otp.getMessage(otp)
	err := o.Otp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
