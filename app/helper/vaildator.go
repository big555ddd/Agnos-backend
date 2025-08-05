package helper

import "regexp"

func IsValidEmail(email string) bool {
	// ใช้ regex มาตรฐานสำหรับ email
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
