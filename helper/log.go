package helper

import (
	"github.com/sirupsen/logrus"
)

func LoggerCreate(messages any) {
	logger := logrus.New()

	logger.Println(messages)
}
