package mail

import (
	"github.com/spf13/viper"
	"log"
	"net/smtp"
)

type UData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Text  string `json:"text"`
}

func EmailSender(data UData) {

	from := viper.GetString("email")
	password := viper.GetString("pwd")
	port := viper.GetString("portSMTP")

	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", from, password, viper.GetString("host"))

	// SendMail uses TLS connection to send the mail
	// The email is sent to all address in the toList,
	// the body should be of type bytes, not strings
	// This returns error if any occurred.
	result := data.makeMsg()
	err := smtp.SendMail(viper.GetString("host")+":"+port, auth, from, []string{viper.GetString("adminemail")}, result)

	// handling the errors
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Successfully sent")
}

func (d *UData) makeMsg() []byte {
	res := ([]byte("Пришло сообщение от пользователя: " + string(d.Name) + "\n" +
		"Контактная почта: " + string(d.Email) + "\n" +
		"Сообщение: " + string(d.Text)))
	return res
}
