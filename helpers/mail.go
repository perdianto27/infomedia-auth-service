package helpers

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
)

func SendRegisterSuccessEmail(toEmail string, username string) error {
	fmt.Println("SendRegisterSuccessEmail");
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("MAIL_FROM"))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Registrasi Berhasil")

	body := fmt.Sprintf(
		"Halo %s,\n\nSelamat! Akun Anda berhasil terdaftar di aplikasi kami.\nSilakan login untuk mulai menggunakan layanan.\n\nTerima kasih.",
		username,
	)

	m.SetBody("text/plain", body)

	d := gomail.NewDialer(
		os.Getenv("MAIL_HOST"),
		587,
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
	)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}