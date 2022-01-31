package config

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	// DEFAULT_SECTION specifies the name of a section if no name provided
	DEFAULT_SECTION = "default"
	// DEFAULT_COMMENT defines what character(s) indicate a comment `#`
	DEFAULT_COMMENT = []byte{'#'}
	// DEFAULT_COMMENT_SEM defines what alternate character(s) indicate a comment `;`
	DEFAULT_COMMENT_SEM = []byte{';'}
	// DEFAULT_MULTI_LINE_SEPARATOR defines what character indicates a multi-line content
	DEFAULT_MULTI_LINE_SEPARATOR = []byte{'\\'}
)

// ConfigInterface defines the behavior of a Config implementation
type ConfigInterface interface {
	String(key string) string
	Strings(key string) []string
	Bool(key string) (bool, error)
	Int(key string) (int, error)
	Int64(key string) (int64, error)
	Float64(key string) (float64, error)
	Set(key string, value string) error
}

// Config represents an implementation of the ConfigInterface
type Config struct {
	// Section:key=value
	data map[string]map[string]string
}

// NewConfig create an empty configuration representation from file.
func NewConfig(confName string) (ConfigInterface, error) {
	c := &Config{
		data: make(map[string]map[string]string),
	}
	err := c.parse(confName)
	return c, err
}

// NewConfigFromText create an empty configuration representation from text.
func NewConfigFromText(text string) (ConfigInterface, error) {
	c := &Config{
		data: make(map[string]map[string]string),
	}
	err := c.parseBuffer(bufio.NewReader(strings.NewReader(text)))
	return c, err
}

// AddConfig adds a new section->key:value to the configuration.
func (c *Config) AddConfig(section string, option string, value string) bool {
	if section == "" {
		section = DEFAULT_SECTION
	}

	if _, ok := c.data[section]; !ok {
		c.data[section] = make(map[string]string)
	}

	_, ok := c.data[section][option]
	c.data[section][option] = value

	return !ok
}

func (c *Config) parse(fname string) (err error) {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	return c.parseBuffer(buf)
}

func (c *Config) parseBuffer(buf *bufio.Reader) error {
	var section string
	var lineNum int
	var buffer bytes.Buffer
	var canWrite bool
	for {
		if canWrite {
			if err := c.write(section, lineNum, &buffer); err != nil {
				return err
			} else {
				canWrite = false
			}
		}
		lineNum++
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			// force write when buffer is not flushed yet
			if buffer.Len() > 0 {
				if err := c.write(section, lineNum, &buffer); err != nil {
					return err
				}
			}
			break
		} else if err != nil {
			return err
		}

		line = bytes.TrimSpace(line)
		switch {
		case bytes.Equal(line, []byte{}), bytes.HasPrefix(line, DEFAULT_COMMENT_SEM),
			bytes.HasPrefix(line, DEFAULT_COMMENT):
			canWrite = true
			continue
		case bytes.HasPrefix(line, []byte{'['}) && bytes.HasSuffix(line, []byte{']'}):
			// force write when buffer is not flushed yet
			if buffer.Len() > 0 {
				if err := c.write(section, lineNum, &buffer); err != nil {
					return err
				}
				canWrite = false
			}
			section = string(line[1 : len(line)-1])
		default:
			var p []byte
			if bytes.HasSuffix(line, DEFAULT_MULTI_LINE_SEPARATOR) {
				p = bytes.TrimSpace(line[:len(line)-1])
				p = append(p, " "...)
			} else {
				p = line
				canWrite = true
			}

			end := len(p)
			for i, value := range p {
				if value == DEFAULT_COMMENT[0] || value == DEFAULT_COMMENT_SEM[0] {
					end = i
					break
				}
			}
			if _, err := buffer.Write(p[:end]); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Config) write(section string, lineNum int, b *bytes.Buffer) error {
	if b.Len() <= 0 {
		return nil
	}

	optionVal := bytes.SplitN(b.Bytes(), []byte{'='}, 2)
	if len(optionVal) != 2 {
		return fmt.Errorf("parse the content error : line %d , %s = ? ", lineNum, optionVal[0])
	}
	option := bytes.TrimSpace(optionVal[0])
	value := bytes.TrimSpace(optionVal[1])
	c.AddConfig(section, string(option), string(value))

	// flush buffer after adding
	b.Reset()

	return nil
}

// Bool lookups up the value using the provided key and converts the value to a bool
func (c *Config) Bool(key string) (bool, error) {
	return strconv.ParseBool(c.get(key))
}

// Int lookups up the value using the provided key and converts the value to a int
func (c *Config) Int(key string) (int, error) {
	return strconv.Atoi(c.get(key))
}

// Int64 lookups up the value using the provided key and converts the value to a int64
func (c *Config) Int64(key string) (int64, error) {
	return strconv.ParseInt(c.get(key), 10, 64)
}

// Float64 lookups up the value using the provided key and converts the value to a float64
func (c *Config) Float64(key string) (float64, error) {
	return strconv.ParseFloat(c.get(key), 64)
}

// String lookups up the value using the provided key and converts the value to a string
func (c *Config) String(key string) string {
	return c.get(key)
}

// Strings lookups up the value using the provided key and converts the value to an array of string
// by splitting the string by comma
func (c *Config) Strings(key string) []string {
	v := c.get(key)
	if v == "" {
		return nil
	}
	return strings.Split(v, ",")
}

// Set sets the value for the specific key in the Config
func (c *Config) Set(key string, value string) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}

	var (
		section string
		option  string
	)

	keys := strings.Split(strings.ToLower(key), "::")
	if len(keys) >= 2 {
		section = keys[0]
		option = keys[1]
	} else {
		option = keys[0]
	}

	c.AddConfig(section, option, value)
	return nil
}

// section.key or key
func (c *Config) get(key string) string {
	var (
		section string
		option  string
	)

	keys := strings.Split(strings.ToLower(key), "::")
	if len(keys) >= 2 {
		section = keys[0]
		option = keys[1]
	} else {
		section = DEFAULT_SECTION
		option = keys[0]
	}

	if value, ok := c.data[section][option]; ok {
		return value
	}

	return ""
}
