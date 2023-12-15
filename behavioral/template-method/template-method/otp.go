package templatemethod

type Otp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}
