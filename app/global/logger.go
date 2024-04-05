package G

import (
	"github.com/charmbracelet/log"
	"os"
	"time"
)

var (
	Logger *log.Logger
)

func init() {
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          AppName,
		Level:           log.DebugLevel,
	})
}
