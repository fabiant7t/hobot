package printer_test

import (
	"bytes"
	"testing"

	"github.com/fabiant7t/hobot/pkg/printer"
	"github.com/google/go-cmp/cmp"
)

type Person struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	YearOfBirth int    `json:"year_of_birth"`
}

func TestPrint(t *testing.T) {
	person := Person{FirstName: "Alice", LastName: "Anderson", YearOfBirth: 1980}
	var p printer.RendererPrinter[Person] = &printer.TablePrinter[Person]{}
	var buf bytes.Buffer
	if err := p.Print(person, &buf); err != nil {
		t.Error(err)
	}
	if got, want := string(buf.Bytes()), "Alice,Anderson,1980\n"; got != want {
		t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
	}
}

func TestPrintAll(t *testing.T) {
	people := []Person{
		{FirstName: "Alice", LastName: "Anderson", YearOfBirth: 1980},
		{FirstName: "Bob", LastName: "Brown", YearOfBirth: 1978},
	}
	var p printer.RendererPrinter[Person] = &printer.TablePrinter[Person]{}
	var buf bytes.Buffer
	if err := p.PrintAll(people, &buf); err != nil {
		t.Error(err)
	}
	if got, want := string(buf.Bytes()), "Alice,Anderson,1980\nBob,Brown,1978\n"; got != want {
		t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
	}
}

func TestPrintWithHeader(t *testing.T) {
	person := Person{FirstName: "Alice", LastName: "Anderson", YearOfBirth: 1980}
	var p printer.RendererPrinter[Person] = &printer.TablePrinter[Person]{WithHeader: true}
	var buf bytes.Buffer
	if err := p.Print(person, &buf); err != nil {
		t.Error(err)
	}
	if got, want := string(buf.Bytes()), "FirstName,LastName,YearOfBirth\nAlice,Anderson,1980\n"; got != want {
		t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
	}
}

func TestPrintAllWithHeader(t *testing.T) {
	people := []Person{
		{FirstName: "Alice", LastName: "Anderson", YearOfBirth: 1980},
		{FirstName: "Bob", LastName: "Brown", YearOfBirth: 1978},
	}
	var p printer.RendererPrinter[Person] = &printer.TablePrinter[Person]{WithHeader: true}
	var buf bytes.Buffer
	if err := p.PrintAll(people, &buf); err != nil {
		t.Error(err)
	}
	if got, want := string(buf.Bytes()), "FirstName,LastName,YearOfBirth\nAlice,Anderson,1980\nBob,Brown,1978\n"; got != want {
		t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
	}
}

func TestPrintAllWithHeaderEmpty(t *testing.T) {
	people := []Person{}
	var p printer.RendererPrinter[Person] = &printer.TablePrinter[Person]{WithHeader: true}
	var buf bytes.Buffer
	if err := p.PrintAll(people, &buf); err != nil {
		t.Error(err)
	}
	if got, want := string(buf.Bytes()), "\n"; got != want {
		t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
	}
}

func TestPrintWithHeaderAndCustomFieldnames(t *testing.T) {
	person := Person{FirstName: "Alice", LastName: "Anderson", YearOfBirth: 1980}
	var tp = &printer.TablePrinter[Person]{WithHeader: true}
	tp.SetFieldNames([]string{"LastName", "Invalid", "FirstName"})
	var p printer.RendererPrinter[Person] = tp
	var buf bytes.Buffer
	if err := p.Print(person, &buf); err != nil {
		t.Error(err)
	}
	if got, want := string(buf.Bytes()), "LastName,Invalid,FirstName\nAnderson,,Alice\n"; got != want {
		t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
	}
}

func TestPrintAllWithHeaderAndCustomFieldnames(t *testing.T) {
	people := []Person{
		{FirstName: "Alice", LastName: "Anderson", YearOfBirth: 1980},
		{FirstName: "Bob", LastName: "Brown", YearOfBirth: 1978},
	}
	var tp = &printer.TablePrinter[Person]{WithHeader: true}
	tp.SetFieldNames([]string{"LastName", "Invalid", "FirstName"})
	var p printer.RendererPrinter[Person] = tp
	var buf bytes.Buffer
	if err := p.PrintAll(people, &buf); err != nil {
		t.Error(err)
	}
	if got, want := string(buf.Bytes()), "LastName,Invalid,FirstName\nAnderson,,Alice\nBrown,,Bob\n"; got != want {
		t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
	}
}
