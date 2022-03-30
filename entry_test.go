package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.WithField("username", "muhbit").Info("Hello Logging")

	// Sebenearnya dia melakukan ini
	// Ini tidak perlu dibuat karena sudah di urus logrusnya
	// Kode di bawah ini hanya untuk tahu saja
	entry := logrus.NewEntry(logger)
	entry.WithField("username", "muhbit")
	entry.Info("Hello Entry")
}
