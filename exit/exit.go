package exit

import (
	"log/slog"
	"os"
)

func WithError(err error) {
	slog.Error("some fatal error occurred", "error", err)
	os.Exit(1)
}
