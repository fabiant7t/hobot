package printer

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type YAMLPrinter[T any] struct{}

func (p *YAMLPrinter[T]) Render(s T) ([]byte, error) {
	return yaml.Marshal(s)
}

func (p *YAMLPrinter[T]) RenderAll(slice []T) ([]byte, error) {
	return yaml.Marshal(slice)
}

func (p *YAMLPrinter[T]) Print(s T, w io.Writer) error {
	b, err := p.Render(s)
	if err != nil {
		return fmt.Errorf("error rendering: %w", err)
	}
	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("error writing: %w", err)
	}
	return nil
}

func (p *YAMLPrinter[T]) PrintAll(slice []T, w io.Writer) error {
	b, err := p.RenderAll(slice)
	if err != nil {
		return fmt.Errorf("error rendering: %w", err)
	}
	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("error writing: %w", err)
	}
	return nil
}
