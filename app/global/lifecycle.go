package G

import "os"

func ErrExit(msg string, err error) {
	Logger.Error(msg, "err", err)
	Exit()
}

func Exit() {
	Cleanup()
	os.Exit(1)
}

func Cleanup() {
	StopBot(true)
}
