package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

// Tidak disarankan, lebih disarankan manual (hook_test)
func TestSingleton(t *testing.T) {
	logrus.Info("Hello World")
	logrus.Warn("Hello World")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello World")
	logrus.Warn("Hello World")
}
