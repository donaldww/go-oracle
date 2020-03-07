// Copyright © 2020 by Donald Wilson donaldww@icloud.com. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package oracle

import (
	"bytes"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// ProgramName returns the name of the current running program.
func ProgramName() string {
	return filepath.Base(os.Args[0])
}

// ThisIPAddr returns the IP of the current machine.
func ThisIPAddr() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer func() {
		_ = conn.Close()
	}()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// MacAddr returns the hardware address of the current computer.
func MacAddr() string {
	// Retrieve a list of network devices.
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			// Search succeeds when interface is up and interface is not empty.
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				return i.HardwareAddr.String()
			}
		}
	}
	return ""
}

// NewViper returns an instance of a spf13/viper file.
func NewViper(cfile, cpath string) (*viper.Viper, error) {
	c := viper.New()
	c.SetConfigName(cfile)

	// AddConfigPath may be called multiple times to add multiple directories.
	c.AddConfigPath(cpath)

	// The config file does not include the file extension.
	if err := c.ReadInConfig(); err != nil {
		return nil, err
	}
	return c, nil
}

const (
	dbuser     = "dbuser"
	dbname     = "dbname"
	dbport     = "dbport"
	dbhost     = "dbhost"
	dbpassword = "dbpassword"
	sslmode    = "sslmode"
)

// PgConnect builds a PostgreSQL connection string using
// values found in `path`/`yamlFile`.yaml.
func PgConnect(yamlFile, path string) string {
	v, err := NewViper(yamlFile, path)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	var connect strings.Builder
	addParam := func(item string) {
		connect.WriteString(item + "=" + v.GetString(item) + " ")
	}
	lastParam := func(item string) {
		connect.WriteString(item + "=" + v.GetString(item))
	}

	addParam(dbuser)
	addParam(dbname)
	addParam(dbport)
	addParam(dbhost)
	addParam(dbpassword)
	lastParam(sslmode)

	return connect.String()
}
