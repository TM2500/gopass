package out

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/gopasspw/gopass/pkg/ctxutil"
)

var (
	// Stdout is exported for tests
	Stdout io.Writer = os.Stdout
	// Stderr is exported for tests
	Stderr io.Writer = os.Stderr
)

func newline(ctx context.Context) string {
	if HasNewline(ctx) {
		return "\n"
	}
	return ""
}

// Print prints the given string
func Print(ctx context.Context, arg string) {
	Printf(ctx, "%s", arg)
}

// Printf formats and prints the given string
func Printf(ctx context.Context, format string, args ...interface{}) {
	if ctxutil.IsHidden(ctx) {
		return
	}
	fmt.Fprintf(Stdout, Prefix(ctx)+format+newline(ctx), args...)
}

// Notice prints the string with an exclamation mark
func Notice(ctx context.Context, arg string) {
	Noticef(ctx, "%s", arg)
}

// Noticef prints the string with an exclamation mark in front
func Noticef(ctx context.Context, format string, args ...interface{}) {
	if ctxutil.IsHidden(ctx) {
		return
	}
	fmt.Fprintf(Stdout, Prefix(ctx)+"⚠ "+format+newline(ctx), args...)
}

// Error prints the string with a red cross in front
func Error(ctx context.Context, arg string) {
	Errorf(ctx, "%s", arg)
}

// Errorf prints the string in red to stderr
func Errorf(ctx context.Context, format string, args ...interface{}) {
	if ctxutil.IsHidden(ctx) {
		return
	}
	fmt.Fprint(Stderr, color.RedString(Prefix(ctx)+"❌ "+format+newline(ctx), args...))
}

// OK prints the string with a green checkmark in front
func OK(ctx context.Context, arg string) {
	OKf(ctx, "%s", arg)
}

// OKf prints the string in with an OK checkmark in front
func OKf(ctx context.Context, format string, args ...interface{}) {
	if ctxutil.IsHidden(ctx) {
		return
	}
	fmt.Fprintf(Stdout, Prefix(ctx)+"✅ "+format+newline(ctx), args...)
}

// Warning prints the string with a warning sign in front
func Warning(ctx context.Context, arg string) {
	Warningf(ctx, "%s", arg)
}

// Warningf prints the string in yellow to stderr and prepends a warning sign
func Warningf(ctx context.Context, format string, args ...interface{}) {
	if ctxutil.IsHidden(ctx) {
		return
	}
	fmt.Fprint(Stderr, color.YellowString(Prefix(ctx)+"⚠ "+format+newline(ctx), args...))
}
