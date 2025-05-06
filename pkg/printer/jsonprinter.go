package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"
)

type JSONPrinter[T any] struct{}

func (p *JSONPrinter[T]) PrintNewLine(w io.Writer) error {
	nl := []byte("\n")
	if runtime.GOOS == "windows" {
		nl = []byte("\r\n")
	}
	_, err := w.Write(nl)
	return err
}

func (p *JSONPrinter[T]) Render(s T) ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

func (p *JSONPrinter[T]) RenderAll(slice []T) ([]byte, error) {
	return json.MarshalIndent(slice, "", "  ")
}

func (p *JSONPrinter[T]) Print(s T, w io.Writer) error {
	b, err := p.Render(s)
	if err != nil {
		return fmt.Errorf("error rendering: %w", err)
	}
	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("error writing: %w", err)
	}
	return p.PrintNewLine(w)
}

func (p *JSONPrinter[T]) PrintAll(slice []T, w io.Writer) error {
	b, err := p.RenderAll(slice)
	if err != nil {
		return fmt.Errorf("error rendering: %w", err)
	}
	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("error writing: %w", err)
	}
	return p.PrintNewLine(w)
}
