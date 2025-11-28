// Package config loads and provides the application configuration.
package config

import (
	"time"
)

type (
	// LogLevel represents the severity or verbosity of log records.
	LogLevel string
)

const (
	// LogLevelDebug captures detailed information, typically useful for development
	// and debugging.
	LogLevelDebug LogLevel = "debug"

	// LogLevelInfo captures general information about the application's operation.
	LogLevelInfo LogLevel = "info"

	// LogLevelWarn captures non-critical events or potentially harmful situations.
	LogLevelWarn LogLevel = "warn"

	// LogLevelError captures critical events or errors that require immediate
	// attention.
	LogLevelError LogLevel = "error"
)

type (
	// LogFormat represents the encoding style of log records.
	LogFormat string
)

const (
	// LogFormatText renders log records as human-readable text.
	LogFormatText LogFormat = "text"

	// LogFormatJSON renders log records as structured JSON objects.
	LogFormatJSON LogFormat = "json"
)

type (
	// LogOutput represents the destination stream of log records.
	LogOutput string
)

const (
	// LogOutputStdout writes log records to the standard output stream (stdout).
	LogOutputStdout LogOutput = "stdout"

	// LogOutputStderr writes log records to the standard error stream (stderr).
	LogOutputStderr LogOutput = "stderr"
)

const (
	// EnvLogLevel specifies the environment variable name for configuring the
	// [LogLevel].
	EnvLogLevel = "LOG_LEVEL"

	// EnvLogFormat specifies the environment variable name for configuring the
	// [LogFormat].
	EnvLogFormat = "LOG_FORMAT"

	// EnvLogOutput specifies the environment variable name for configuring the
	// [LogOutput].
	EnvLogOutput = "LOG_OUTPUT"

	// EnvServerAddress specifies the environment variable name for configuring the
	// server's address.
	EnvServerAddress = "SERVER_ADDRESS"

	// EnvServerReadTimeout specifies the environment variable name for configuring the
	// server's read timeout.
	EnvServerReadTimeout = "SERVER_READ_TIMEOUT"

	// EnvServerReadHeaderTimeout specifies the environment variable name for
	// configuring the server's read header timeout.
	EnvServerReadHeaderTimeout = "SERVER_READ_HEADER_TIMEOUT"

	// EnvServerWriteTimeout specifies the environment variable name for configuring
	// the server's write timeout.
	EnvServerWriteTimeout = "SERVER_WRITE_TIMEOUT"

	// EnvServerIdleTimeout specifies the environment variable name for configuring the
	// server's idle timeout.
	EnvServerIdleTimeout = "SERVER_IDLE_TIMEOUT"

	// EnvServerShutdownTimeout specifies the environment variable name for configuring
	// the server's shutdown timeout.
	EnvServerShutdownTimeout = "SERVER_SHUTDOWN_TIMEOUT"
)

const (
	DefaultLogLevel                = LogLevelInfo
	DefaultLogFormat               = LogFormatText
	DefaultLogOutput               = LogOutputStdout
	DefaultServerAddress           = "localhost:8080"
	DefaultServerReadTimeout       = 5 * time.Second
	DefaultServerReadHeaderTimeout = 2 * time.Second
	DefaultServerWriteTimeout      = 10 * time.Second
	DefaultServerIdleTimeout       = 1 * time.Minute
	DefaultServerShutdownTimeout   = 15 * time.Second
)
