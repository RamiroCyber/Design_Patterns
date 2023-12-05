package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Email)\n", message)
}

type SmsNotifier struct{}

func (SmsNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: SMS)\n", message)
}

type NotifierFactory interface {
	CreateNotifier() Notifier
}

type EmailNotifierFactory struct{}

func (EmailNotifierFactory) CreateNotifier() Notifier {
	return EmailNotifier{}
}

type SmsNotifierFactory struct{}

func (SmsNotifierFactory) CreateNotifier() Notifier {
	return SmsNotifier{}
}

func CreateNotifierEmail(factory NotifierFactory) {
	factory = EmailNotifierFactory{}

	notifier := factory.CreateNotifier()

	notifier.Send("Olá, Bem-vindo ao estudo de Factory Method")
}

func CreateNotifierSms(factory NotifierFactory) {
	factory = SmsNotifierFactory{}

	notifier := factory.CreateNotifier()

	notifier.Send("Olá, Bem-vindo ao estudo de Factory Method")
}

func main() {
	var factory NotifierFactory
	CreateNotifierEmail(factory)
	CreateNotifierSms(factory)
}
