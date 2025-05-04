package printer

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"sync"
)

type TablePrinter[T any] struct {
	sync.Mutex
	WithHeader    bool
	headerWritten bool
	fieldNames    []string
}

// SetFieldNames allows you to customize which field names are printed
// and in what order, instead of automatically deriving and printing
// all fields in their definition order.
func (p *TablePrinter[T]) SetFieldNames(fieldNames []string) {
	p.Lock()
	defer p.Unlock()
	p.fieldNames = fieldNames
}

// SetWithHeader is a flag that determines whether the table is rendered
// with a header.
func (p *TablePrinter[T]) SetWithHeader(withHeader bool) {
	p.Lock()
	defer p.Unlock()
	p.WithHeader = withHeader
}

func (p *TablePrinter[T]) Render(s T) ([]byte, error) {
	p.Lock()
	defer p.Unlock()
	if len(p.fieldNames) == 0 {
		fieldNames, err := FieldNames(s)
		if err != nil {
			return nil, fmt.Errorf("error getting field names: %w", err)
		}
		p.fieldNames = fieldNames
	}
	buf := bytes.NewBuffer(nil)
	w := csv.NewWriter(buf)
	if p.WithHeader && !p.headerWritten {
		if err := w.Write(p.fieldNames); err != nil {
			return nil, fmt.Errorf("error writing headers: %w", err)
		}
		p.headerWritten = true
	}
	values := make([]string, len(p.fieldNames))
	m, err := FieldMap(s)
	if err != nil {
		return nil, fmt.Errorf("error getting field map: %w", err)
	}
	for i, k := range p.fieldNames {
		v, ok := m[k]
		if ok {
			values[i] = fmt.Sprintf("%v", v)
		} else {
			values[i] = ""
		}
	}
	if err := w.Write(values); err != nil {
		return nil, fmt.Errorf("error writing values: %w", err)
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

func (p *TablePrinter[T]) RenderAll(slice []T) ([]byte, error) {
	p.Lock()
	defer p.Unlock()
	if len(p.fieldNames) == 0 && len(slice) > 0 {
		fieldNames, err := FieldNames(slice[0])
		if err != nil {
			return nil, fmt.Errorf("error getting field names: %w", err)
		}
		p.fieldNames = fieldNames
	}
	buf := bytes.NewBuffer(nil)
	w := csv.NewWriter(buf)
	if p.WithHeader && !p.headerWritten {
		if err := w.Write(p.fieldNames); err != nil {
			return nil, fmt.Errorf("error writing headers: %w", err)
		}
		p.headerWritten = true
	}
	for _, s := range slice {
		values := make([]string, len(p.fieldNames))
		m, err := FieldMap(s)
		if err != nil {
			return nil, fmt.Errorf("error getting field map: %w", err)
		}
		for i, k := range p.fieldNames {
			v, ok := m[k]
			if ok {
				values[i] = fmt.Sprintf("%v", v)
			} else {
				values[i] = ""
			}
		}
		if err := w.Write(values); err != nil {
			return nil, fmt.Errorf("error writing values: %w", err)
		}
	}
	w.Flush()
	return buf.Bytes(), w.Error()
}

func (p *TablePrinter[T]) Print(s T, w io.Writer) error {
	b, err := p.Render(s)
	if err != nil {
		return fmt.Errorf("error rendering: %w", err)
	}
	_, err = w.Write(b)
	return err
}

func (p *TablePrinter[T]) PrintAll(slice []T, w io.Writer) error {
	b, err := p.RenderAll(slice)
	if err != nil {
		return fmt.Errorf("error rendering all: %w", err)
	}
	if _, err = w.Write(b); err != nil {
		return fmt.Errorf("error printing: %w", err)
	}
	return nil
}

func (p *TablePrinter[T]) Reset() {
	p.Lock()
	defer p.Unlock()
	p.headerWritten = false
}
