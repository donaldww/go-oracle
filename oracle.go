// Copyright 2020 by Donald Wilson. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.
package go_oracle

import (
	"log"
	"net"
	"os"
	"path/filepath"
)

// ProgramName returns the name of the current running program.
func ProgramName() string {
	return filepath.Base(os.Args[0])
}

// ThisIPAddr returns the IP of the current machine.
func ThisIPAddr() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
