// Sangat disarankan pakai Zoho, karena pake gmail lebih sulit tai lah ada authentication berlapis, nyusahin cok
package main

import (
    "gopkg.in/gomail.v2"
    "log"
)

const CONFIG_SMTP_HOST = "smtp.zoho.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Datenshi Support <support@troke.id>"
const CONFIG_AUTH_EMAIL = "namaemailkamubos@crot.com"
const CONFIG_AUTH_PASSWORD = "passwordnya_bos_moncrot"

func main() {
    mailer := gomail.NewMessage()
    mailer.SetHeader("From", CONFIG_SENDER_NAME)
    mailer.SetHeader("To", "Emailtujuan_bos_moncrot@gmail")
    mailer.SetHeader("Subject", "Test mail")
    mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")

    dialer := gomail.NewDialer(
        CONFIG_SMTP_HOST,
        CONFIG_SMTP_PORT,
        CONFIG_AUTH_EMAIL,
        CONFIG_AUTH_PASSWORD,
    )

    err := dialer.DialAndSend(mailer)
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Println("Mail sent!")
}
