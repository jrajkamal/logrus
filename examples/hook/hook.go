package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/tobi/airbrake-go"
)

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.TextFormatter) // default
	log.Hooks.Add(new(logrus.AirbrakeHook))
}

func main() {
	airbrake.Endpoint = "https://exceptions.shopify.com/notifier_api/v2/notices.xml"
	airbrake.ApiKey = "9615665b35b79abd557eb2be7772a862"
	airbrake.Environment = "production"

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}
