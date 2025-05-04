package printer

import "io"

type Renderer[T any] interface {
	Render(T) ([]byte, error)
	RenderAll([]T) ([]byte, error)
}

type Printer[T any] interface {
	Print(T, io.Writer) error
	PrintAll([]T, io.Writer) error
}

type RendererPrinter[T any] interface {
	Renderer[T]
	Printer[T]
}
