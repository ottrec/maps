package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/go-json-experiment/json/jsontext"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	for _, name := range os.Args[1:] {
		buf, err := os.ReadFile(name)
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
		out, err := Format(buf, "  ")
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
		err = os.WriteFile(name, append(out, '\n'), 0644)
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
	}
	return nil
}

// Format reformats raw JSON as indented output, except that arrays whose
// elements are all numbers or all strings are emitted on a single line.
func Format(raw []byte, indent string) ([]byte, error) {
	f := &formatter{indent: indent}
	if err := f.value(jsontext.NewDecoder(bytes.NewReader(raw)), 0); err != nil {
		return nil, err
	}
	return f.out.Bytes(), nil
}

type formatter struct {
	out    bytes.Buffer
	indent string
}

func (f *formatter) value(dec *jsontext.Decoder, depth int) error {
	switch dec.PeekKind() {
	case '{':
		return f.object(dec, depth)
	case '[':
		return f.array(dec, depth)
	default:
		v, err := dec.ReadValue()
		if err != nil {
			return err
		}
		return f.writeCompact(v)
	}
}

func (f *formatter) object(dec *jsontext.Decoder, depth int) error {
	if _, err := dec.ReadToken(); err != nil { // '{'
		return err
	}
	f.out.WriteByte('{')
	first := true
	for {
		k := dec.PeekKind()
		if k == '}' {
			break
		}
		if !first {
			f.out.WriteByte(',')
		}
		first = false
		f.newline(depth + 1)
		key, err := dec.ReadValue() // key
		if err != nil {
			return err
		}
		if err := f.writeCompact(key); err != nil {
			return err
		}
		f.out.WriteString(": ")
		if err := f.value(dec, depth+1); err != nil {
			return err
		}
	}
	if _, err := dec.ReadToken(); err != nil { // '}'
		return err
	}
	if !first {
		f.newline(depth)
	}
	f.out.WriteByte('}')
	return nil
}

func (f *formatter) array(dec *jsontext.Decoder, depth int) error {
	raw, err := dec.ReadValue()
	if err != nil {
		return err
	}
	if isScalarArray(raw) {
		return f.scalarArray(raw)
	}

	sub := jsontext.NewDecoder(bytes.NewReader(raw))
	if _, err := sub.ReadToken(); err != nil { // '['
		return err
	}
	f.out.WriteByte('[')
	first := true
	for {
		k := sub.PeekKind()
		if k == ']' {
			break
		}
		if !first {
			f.out.WriteByte(',')
		}
		first = false
		f.newline(depth + 1)
		if err := f.value(sub, depth+1); err != nil {
			return err
		}
	}
	if !first {
		f.newline(depth)
	}
	f.out.WriteByte(']')
	return nil
}

func (f *formatter) scalarArray(raw jsontext.Value) error {
	dec := jsontext.NewDecoder(bytes.NewReader(raw))
	if _, err := dec.ReadToken(); err != nil { // '['
		return err
	}
	f.out.WriteByte('[')
	first := true
	for {
		k := dec.PeekKind()
		if k == ']' {
			break
		}
		if !first {
			f.out.WriteString(", ")
		}
		first = false
		v, err := dec.ReadValue()
		if err != nil {
			return err
		}
		if err := f.writeCompact(v); err != nil {
			return err
		}
	}
	f.out.WriteByte(']')
	return nil
}

func (f *formatter) writeCompact(v jsontext.Value) error {
	if err := v.Compact(); err != nil {
		return err
	}
	f.out.Write(v)
	return nil
}

func (f *formatter) newline(depth int) {
	f.out.WriteByte('\n')
	f.out.WriteString(strings.Repeat(f.indent, depth))
}

func isScalarArray(raw jsontext.Value) bool {
	dec := jsontext.NewDecoder(bytes.NewReader(raw))
	if k, err := dec.ReadToken(); err != nil || k.Kind() != '[' {
		return false
	}
	for {
		k := dec.PeekKind()
		if k == ']' {
			break
		}
		if k != '"' && k != '0' {
			return false
		}
		if _, err := dec.ReadValue(); err != nil {
			return false
		}
	}
	return true
}
