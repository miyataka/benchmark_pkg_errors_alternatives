package main

import (
	"fmt"
	"testing"

	stderrors "errors"

	cdberrors "github.com/cockroachdb/errors"
	arkerrors "github.com/goark/errs"
	pkgerrors "github.com/pkg/errors"
)

func noErrors(at, depth int) error {
	if at >= depth {
		return stderrors.New("no error")
	}
	return noErrors(at+1, depth)
}

func pkgErrors(at, depth int) error {
	if at >= depth {
		return pkgerrors.New("ye error")
	}
	return pkgErrors(at+1, depth)
}

func cdbErrors(at, depth int) error {
	if at >= depth {
		return cdberrors.New("ye error")
	}
	return cdbErrors(at+1, depth)
}

func arkErrors(at, depth int) error {
	if at >= depth {
		return arkerrors.New("ye error")
	}
	return arkErrors(at+1, depth)
}

// GlobalE is an exported global to store the result of benchmark results,
// preventing the compiler from optimising the benchmark functions away.
var GlobalE interface{}

func BenchmarkErrors(b *testing.B) {
	type run struct {
		name  string
		stack int
	}
	runs := []run{
		{"errors", 10},
		{"errors", 100},
		{"errors", 1000},
		{"pkg/errors", 10},
		{"pkg/errors", 100},
		{"pkg/errors", 1000},
		{"cockroachdb/errors", 10},
		{"cockroachdb/errors", 100},
		{"cockroachdb/errors", 1000},
		{"goark/errs", 10},
		{"goark/errs", 100},
		{"goark/errs", 1000},
	}
	for _, r := range runs {
		part := r.name
		name := fmt.Sprintf("%s-stack-%d", part, r.stack)
		var f (func(a, b int) error)
		switch r.name {
		case "errors":
			f = noErrors
		case "pkg/errors":
			f = pkgErrors
		case "cockroachdb/errors":
			f = cdbErrors
		case "goark/errs":
			f = arkErrors
		}
		b.Run(name, func(b *testing.B) {
			var err error
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				err = f(0, r.stack)
			}
			b.StopTimer()
			GlobalE = err
		})
	}
}

func BenchmarkStackFormatting(b *testing.B) {
	type run struct {
		name   string
		stack  int
		format string
	}
	runs := []run{
		{"errors", 10, "%s"},
		{"errors", 10, "%v"},
		{"errors", 10, "%+v"},
		{"errors", 30, "%s"},
		{"errors", 30, "%v"},
		{"errors", 30, "%+v"},
		{"errors", 60, "%s"},
		{"errors", 60, "%v"},
		{"errors", 60, "%+v"},
		{"pkg/errors", 10, "%s"},
		{"pkg/errors", 10, "%v"},
		{"pkg/errors", 10, "%+v"},
		{"pkg/errors", 30, "%s"},
		{"pkg/errors", 30, "%v"},
		{"pkg/errors", 30, "%+v"},
		{"pkg/errors", 60, "%s"},
		{"pkg/errors", 60, "%v"},
		{"pkg/errors", 60, "%+v"},
		{"cockroachdb/errors", 10, "%s"},
		{"cockroachdb/errors", 10, "%v"},
		{"cockroachdb/errors", 10, "%+v"},
		{"cockroachdb/errors", 30, "%s"},
		{"cockroachdb/errors", 30, "%v"},
		{"cockroachdb/errors", 30, "%+v"},
		{"cockroachdb/errors", 60, "%s"},
		{"cockroachdb/errors", 60, "%v"},
		{"cockroachdb/errors", 60, "%+v"},
		{"goark/errs", 10, "%s"},
		{"goark/errs", 10, "%v"},
		{"goark/errs", 10, "%+v"},
		{"goark/errs", 30, "%s"},
		{"goark/errs", 30, "%v"},
		{"goark/errs", 30, "%+v"},
		{"goark/errs", 60, "%s"},
		{"goark/errs", 60, "%v"},
		{"goark/errs", 60, "%+v"},
	}

	var stackStr string
	for _, r := range runs {
		name := fmt.Sprintf("%s-%s-stack-%d", r.name, r.format, r.stack)
		var f (func(a, b int) error)
		switch r.name {
		case "errors":
			f = noErrors
		case "pkg/errors":
			f = pkgErrors
		case "cockroachdb/errors":
			f = cdbErrors
		case "goark/errs":
			f = arkErrors
		}
		b.Run(name, func(b *testing.B) {
			err := f(0, r.stack)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				stackStr = fmt.Sprintf(r.format, err)
			}
			b.StopTimer()
		})
	}

	GlobalE = stackStr
}

func BenchmarkCockroachDBStackFormatting(b *testing.B) {
	type run struct {
		name   string
		stack  int
		format string
	}
	runs := []run{
		{"cockroachdb/errors", 10, "%s"},
	}

	var stackStr string
	for _, r := range runs {
		name := fmt.Sprintf("%s-%s-stack-%d", r.name, r.format, r.stack)
		b.Run(name, func(b *testing.B) {
			err := cdbErrors(0, r.stack)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				stackStr = fmt.Sprintf(r.format, err)
			}
			b.StopTimer()
		})
	}

	GlobalE = stackStr
}

func BenchmarkCockroachDBStackFormattingWithNew(b *testing.B) {
	type run struct {
		name   string
		stack  int
		format string
	}
	runs := []run{
		{"cockroachdb/errors", 10, "%s"},
	}

	var stackStr string
	for _, r := range runs {
		name := fmt.Sprintf("%s-%s-stack-%d", r.name, r.format, r.stack)
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				err := cdbErrors(0, r.stack)
				stackStr = fmt.Sprintf(r.format, err)
			}
			b.StopTimer()
		})
	}

	GlobalE = stackStr
}

func BenchmarkGoarkStackFormatting(b *testing.B) {
	type run struct {
		name   string
		stack  int
		format string
	}
	runs := []run{
		{"goark/errs", 10, "%+v"},
	}

	var stackStr string
	for _, r := range runs {
		name := fmt.Sprintf("%s-%s-stack-%d", r.name, r.format, r.stack)
		b.Run(name, func(b *testing.B) {
			err := arkErrors(0, r.stack)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				stackStr = fmt.Sprintf(r.format, err)
			}
			b.StopTimer()
		})
	}

	GlobalE = stackStr
}

func BenchmarkGoarkStackFormattingWithNew(b *testing.B) {
	type run struct {
		name   string
		stack  int
		format string
	}
	runs := []run{
		{"goark/errs", 10, "%+v"},
	}

	var stackStr string
	for _, r := range runs {
		name := fmt.Sprintf("%s-%s-stack-%d", r.name, r.format, r.stack)
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				err := arkErrors(0, r.stack)
				stackStr = fmt.Sprintf(r.format, err)
			}
			b.StopTimer()
		})
	}

	GlobalE = stackStr
}
