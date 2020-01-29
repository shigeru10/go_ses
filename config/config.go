package config

// Email is a constructor
type Email struct {
	Sender    string `envconfig:"SENDER" default:""`
	Recipient string `envconfig:"RECIPIENT" default:""`
}
