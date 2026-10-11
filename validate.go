package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// errUsage marks input rejected before uv starts; runCLI maps it to exit status 2.
// It wraps ErrInvalidArgs so callers can inspect failure types programmatically.
var errUsage = fmt.Errorf("%w: usage error", ErrInvalidArgs)

// CLIInputs specifies all user-supplied inputs evaluated at CLI boundaries:
// command arguments and environment configuration overrides.
type CLIInputs struct {
	Args       []string
	UVOverride string
}

// ValidateInputs runs boundary validation on all CLI entry point inputs.
// It returns a typed error wrapping ErrInvalidArgs if arguments are invalid,
// or ErrInvalidOverride if environment configuration overrides are malformed.
func ValidateInputs(inputs CLIInputs) error {
	if len(inputs.Args) > 0 || inputs.UVOverride == "" {
		if err := validateArgs(inputs.Args); err != nil {
			return err
		}
	}
	if inputs.UVOverride != "" {
		if err := validateUVOverride(inputs.UVOverride); err != nil {
			return err
		}
	}
	return nil
}

// validateArgs rejects a blank pip command, which uv would only report as an
// unrecognized empty subcommand. Values are never echoed: arguments can embed
// index credentials.
func validateArgs(args []string) error {
	if len(args) == 0 || strings.TrimSpace(args[0]) == "" {
		return fmt.Errorf("%w: pip command is empty", errUsage)
	}
	return nil
}

// validateUVOverride checks a non-empty UVPIP_UV before path lookup, so common
// quoting mistakes get a precise message instead of "executable not found".
func validateUVOverride(path string) error {
	switch {
	case strings.TrimSpace(path) == "":
		return fmt.Errorf("%w: UVPIP_UV is set but blank; unset it or give an absolute path", ErrInvalidOverride)
	case strings.TrimSpace(path) != path:
		return fmt.Errorf("%w: UVPIP_UV has leading or trailing whitespace; check its quoting", ErrInvalidOverride)
	case !filepath.IsAbs(path):
		return fmt.Errorf("%w: UVPIP_UV must be an absolute executable path", ErrInvalidOverride)
	}
	return nil
}
