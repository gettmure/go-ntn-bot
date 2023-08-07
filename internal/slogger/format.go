package slogger

import (
	"golang.org/x/exp/slog"
)

func FormatError(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func FormatBody(body []byte) slog.Attr {
	return slog.Attr{
		Key:   "body",
		Value: slog.StringValue(string(body)),
	}
}
