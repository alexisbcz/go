package exit

import (
	"log/slog"
	"os"
)

func WithErr(err error) {
	slog.Error("some fatal error occurred", "error", err)
	os.Exit(1)
}
