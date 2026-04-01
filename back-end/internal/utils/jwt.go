package utils

import "fmt"

func GenerateJWT(email string) string {
	return fmt.Sprintf("token-%s", email)
}
