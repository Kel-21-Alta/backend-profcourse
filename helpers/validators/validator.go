package validators

import (
	"github.com/spf13/viper"
	"regexp"
)

func CheckEmail(email string) bool {
	var emailCompany string
	if viper.GetString("email.regex") == "" {
		emailCompany = "gmail.com"
	} else {
		emailCompany = viper.GetString("email.regex")
	}

	regexEmail := `^[a-z0-9._%+\-]+@` + emailCompany + `$`
	emailRegex := regexp.MustCompile(regexEmail)
	return emailRegex.MatchString(email)
}
