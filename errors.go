package main

import "errors"

var (
	// ErrInvalidArgs indicates invalid or empty CLI arguments.
	ErrInvalidArgs = errors.New("invalid arguments")

	// ErrInvalidOverride indicates malformed UVPIP_UV environment variable configuration.
	ErrInvalidOverride = errors.New("invalid uv override")

	// ErrTranslation indicates command translation failure.
	ErrTranslation = errors.New("command translation failed")

	// ErrUVNotFound indicates uv could not be located in PATH or candidate locations.
	ErrUVNotFound = errors.New("uv executable not found")

	// ErrRecursiveExecution prevents uvpip from calling itself in an infinite loop.
	ErrRecursiveExecution = errors.New("uv resolves to uvpip itself; refusing recursive execution")
)
