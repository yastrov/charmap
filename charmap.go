// Package charmap implements functions for character encodings conversion.
// A number of 8bit encodings are supported. The package provides Encode and
// Decode functions to convert a string from and to UTF-8 respectively.
package charmap

import (
	"bytes"
	"errors"
	"strings"
	"unicode/utf8"
)

var aliasesMap = make(map[string]string)
var codecsMap = make(map[string]codec)

type codec interface {
	Encode(data string) (string, error)
	Decode(data string) (string, error)
	EncodeToBuffer(s []byte) (*bytes.Buffer, error)
	DecodeToBuffer(s []byte) (*bytes.Buffer, error)
}

func register(c codec, name string, aliases ...string) {
	codecsMap[name] = c
	for _, alias := range aliases {
		aliasesMap[alias] = name
	}
}

var ErrUnknownEncoding error = errors.New("encoding is not supported")
var ErrInvalidCodepoint error = errors.New("cannot convert one or more codepoints")

// List returns a list of all supported encodings as a slice of strings
func List() []string {
	list := make([]string, 0)
	for name, _ := range codecsMap {
		list = append(list, name)
	}
	return list
}

func getCodecForEncoding(encoding string) string {
	encoding = strings.ToUpper(encoding)
	encoding = strings.Replace(encoding, "_", "-", -1)

	if name, ok := aliasesMap[encoding]; ok {
		encoding = name
	}

	return encoding
}

// Encode converts a string from UTF-8 to the specified encoding. Returns converted string.
// If the input string contains illegal characters for the specified encoding,
// these characters will be replaced with a substitute character ('?') and
// ErrInvalidCodepoint will be returned in error value.
// If the specified encoding is unknown, it will return the input string and ErrUnknownEncoding
func Encode(data string, encoding string) (string, error) {
	encoding = getCodecForEncoding(encoding)

	if codec, ok := codecsMap[encoding]; ok {
		result, err := codec.Encode(data)
		return result, err
	}

	return data, ErrUnknownEncoding
}

// Decode converts a string from the specified encoding to UTF-8.  Returns converted string.
// If the input string contains illegal characters for the specified encoding,
// these characters will be replaced with a substitute character (utf8.RuneError) and
// ErrInvalidCodepoint will be returned in error value.
// If the specified encoding is unknown, it will return the input string and ErrUnknownEncoding
func Decode(data string, encoding string) (string, error) {
	encoding = getCodecForEncoding(encoding)

	if codec, ok := codecsMap[encoding]; ok {
		result, err := codec.Decode(data)
		return result, err
	}

	return data, ErrUnknownEncoding
}

// simple 8bit codecs definition support

func reverseByteRuneMap(m map[byte]rune) (newmap map[rune]byte) {
	newmap = make(map[rune]byte, len(m))
	for k, v := range m {
		newmap[v] = k
	}
	return
}

func mapBytesToRunes(cm map[byte]rune, data string) (result string, err error) {
	size := len(data)
	buf := bytes.NewBuffer(make([]byte, 0, size))

	for i := 0; i < size; i++ {
		c := data[i]
		if r, ok := cm[c]; ok {
			buf.WriteRune(r)
		} else {
			err = ErrInvalidCodepoint
			buf.WriteRune(utf8.RuneError)
		}
	}

	result = buf.String()
	return result, err
}

func mapRunesToBytes(cm map[rune]byte, data string) (result string, err error) {
	size := len(data)
	buf := bytes.NewBuffer(make([]byte, 0, size/2))

	for _, r := range data {
		if c, ok := cm[r]; ok {
			buf.WriteByte(c)
		} else {
			err = ErrInvalidCodepoint
			buf.WriteByte('?')
		}
	}

	result = buf.String()
	return result, err
}

/*------------------------------------------*/
/*Section for buffer support*/
func EncodeToBuffer(data []byte, encoding string) (*bytes.Buffer, error) {
	encoding = getCodecForEncoding(encoding)

	if codec, ok := codecsMap[encoding]; ok {
		result, err := codec.EncodeToBuffer(data)
		return result, err
	}

	return bytes.NewBuffer(data), ErrUnknownEncoding
}

func DecodeToBuffer(data []byte, encoding string) (*bytes.Buffer, error) {
	encoding = getCodecForEncoding(encoding)

	if codec, ok := codecsMap[encoding]; ok {
		result, err := codec.DecodeToBuffer(data)
		return result, err
	}

	return bytes.NewBuffer(data), ErrUnknownEncoding
}

func mapBytesToRunesBuffer(cm map[byte]rune, data []byte) (result *bytes.Buffer, err error) {
	size := len(data)
	result = bytes.NewBuffer(make([]byte, 0, size))

	for i := 0; i < size; i++ {
		c := data[i]
		if r, ok := cm[c]; ok {
			result.WriteRune(r)
		} else {
			err = ErrInvalidCodepoint
			result.WriteRune(utf8.RuneError)
		}
	}

	return result, err
}

func mapRunesToBytesBuffer(cm map[rune]byte, data []byte) (result *bytes.Buffer, err error) {
	size := len(data)
	result = bytes.NewBuffer(make([]byte, 0, size/2))

	for _, r := range data {
		if c, ok := cm[rune(r)]; ok {
			result.WriteByte(c)
		} else {
			err = ErrInvalidCodepoint
			result.WriteByte('?')
		}
	}

	return result, err
}

/*End buffer support section*/

type codecMap8Bit struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c *codecMap8Bit) Encode(s string) (string, error) {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c *codecMap8Bit) Decode(s string) (string, error) {
	return mapBytesToRunes(c.DecodeMap, s)
}

func (c *codecMap8Bit) EncodeToBuffer(s []byte) (*bytes.Buffer, error) {
	return mapRunesToBytesBuffer(c.EncodeMap, s)
}

func (c *codecMap8Bit) DecodeToBuffer(s []byte) (*bytes.Buffer, error) {
	return mapBytesToRunesBuffer(c.DecodeMap, s)
}
