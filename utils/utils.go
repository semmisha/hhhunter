package utils

import (
	"github.com/sirupsen/logrus"
	"strings"
)

func MapToString(config map[string]string, logger *logrus.Logger) string {
	bld := strings.Builder{}
	for i, j := range config {
		if len(i) > 0 && len(j) > 0 {
			bld.WriteString(i + "=" + j + " ")

		}
	}
	if len(bld.String()) == 0 {

		logger.Error("SQL config len = 0 ")

	}
	return bld.String()

}
