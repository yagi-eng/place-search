package errmsg

import "github.com/sirupsen/logrus"

// LogFatal errがnilでなければエラー出力する
func LogFatal(err error) {
	if err != nil {
		logrus.Fatal("fatal error: %s", err)
	}
}
