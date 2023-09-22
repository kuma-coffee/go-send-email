package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/gomail.v2"
)

var (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

func sendMailSimple(subject, body, emailSender, passwordSender, emailReceiver string) {
	auth := smtp.PlainAuth(
		"",
		emailSender,
		passwordSender,
		smtpAuthAddress,
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail(
		smtpServerAddress,
		auth,
		emailSender,
		[]string{emailReceiver},
		[]byte(msg),
	)

	if err != nil {
		log.Fatal(err)
	}
}

func sendMailSimpleHTML(subject, html, templatePath, emailSender, passwordSender, emailReceiver, nameSender string) {
	// get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(&body, struct{ Name string }{Name: "Kuma"})

	auth := smtp.PlainAuth(
		"",
		emailSender,
		passwordSender,
		smtpAuthAddress,
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	// msg := "Subject: " + subject + "\n" + headers + "\n\n" + html
	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		smtpServerAddress,
		auth,
		emailSender,
		[]string{emailReceiver},
		[]byte(msg),
	)

	if err != nil {
		log.Fatal(err)
	}
}

func sendGomail(subject, html, templatePath, emailSender, passwordSender, emailReceiver, nameSender string) {
	// get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(&body, struct{ Name string }{Name: "Kuma"})

	// send with gomail
	m := gomail.NewMessage()
	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailReceiver)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())
	m.Attach("img.jpeg")

	d := gomail.NewDialer(smtpAuthAddress, 587, emailSender, passwordSender)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func sendSendGrip(emailSender, passwordSender, emailReceiver string) {
	from := mail.NewEmail("Example User", emailSender)
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Example User", emailReceiver)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func main() {
	// load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error laoding .env file")
	}

	var (
		// nameSender    = os.Getenv("EMAIL_SENDER_NAME")
		emailSender   = os.Getenv("EMAIL_SENDER_ADDRESS")
		emailPassword = os.Getenv("EMAIL_SENDER_PASSWORD")
		emailReceiver = os.Getenv("EMAIL_RECEIVER_ADDRESS")
		// subject       = "Test Golang Send Email"
		// body          = "Test Golang Send Email from " + emailSender
		// html         = "<h1>Test Golang Send Email</h1>" + "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla euismod, odio et pretium vehicula, ante sapien tristique sapien, venenatis rutrum enim eros id metus. Donec vestibulum, libero et cursus aliquet, erat neque laoreet velit, at faucibus orci sem sed ex. Vivamus sodales, libero non fermentum congue, mi ante placerat erat, non tincidunt orci risus a est. Mauris ipsum ligula, vehicula fringilla metus in, malesuada tincidunt augue. Praesent maximus facilisis justo in convallis. Sed eget orci commodo, euismod odio non, tempor lacus. Donec nulla nibh, pulvinar nec eros ac, ornare malesuada enim. Phasellus id volutpat odio. Nullam in eros risus.</p>"
		// templatePath = "./test.html"
	)

	// sendMailSimple(subject, body, emailSender, emailPassword, emailReceiver)

	// sendMailSimpleHTML(subject, html, templatePath, emailSender, emailPassword, emailReceiver, nameSender)

	// sendGomail(subject, html, templatePath, emailSender, emailPassword, emailReceiver, nameSender)

	sendSendGrip(emailSender, emailPassword, emailReceiver)
}
