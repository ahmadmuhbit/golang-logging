package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "muhbit").Info("Hello World")

	logger.WithField("username", "ahmad").
		WithField("name", "Ahmad Muhbit").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "muhbit",
		"name":     "Ahmad Muhbit",
	}).Info("Hello World")
}
