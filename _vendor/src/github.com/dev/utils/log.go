package utils

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

func LogInfo(s string, r *http.Request) {
	logrus.Infoln(s, r.RemoteAddr, r.Proto, r.Method, r.RequestURI, r.Host)
}

func LogWarn(s string, r *http.Request) {
	logrus.Warningln(s, r.RemoteAddr, r.Proto, r.Method, r.RequestURI, r.Host)
}
