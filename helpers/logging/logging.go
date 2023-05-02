package logging

import (
	"fmt"
	"net/http"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func newLoggrus() {
	log = logrus.New()
	// Will log anything that is info or above (warn, error, fatal, panic)
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
		DisableQuote:  true,
	})
}

func Infof(infoMsg string, req *http.Request) {
	fields := logrus.Fields{}
	if req.Method != "" {
		fields["Method"] = req.Method
	}

	newLoggrus()
	log.WithFields(fields).Info(infoMsg)
}

func Errorf(err error, req *http.Request) {
	fields := logrus.Fields{}
	if req.Method != "" {
		fields["Method"] = req.Method
	}
	if req.Body != nil {
		fields["Params"] = req.Body
	}

	newLoggrus()
	log.WithFields(fields).Errorf(fmt.Sprintf("Error : %s", err.Error()))
}
