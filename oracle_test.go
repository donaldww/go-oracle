// Copyright 2020 by Donald Wilson. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.
package oracle

import (
	"fmt"
	"testing"
)

func TestProgramName(t *testing.T) {
	progName := ProgramName()
	n1 := "go-oracle.test"                                      // when 'go test' is run on the command line
	n2 := "___go_test_github_com_donaldww_go_oracle"            // when run in GoLand as part of global run
	n3 := "___TestProgramName_in_github_com_donaldww_go_oracle" // when run as individual test in GoLand

	if !(progName == n1 || progName == n2 || progName == n3) {
		t.Errorf("error: got=%s, exected either %s, %s, or %s\n", progName, n1, n2, n3)
	}
}

func ExampleThisIPAddr() {
	fmt.Println(ThisIPAddr())
	// Output:
	// 192.168.0.11
}

func ExampleMacAddr() {
	fmt.Println(MacAddr())
	// Output:
	// 78:7b:8a:af:4a:4c
}

func testDataHome() string {
	return "./testdata/oracle/"
}

func ExampleNewViper() {
	v, err := NewViper("oracle", testDataHome())
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v.GetString("hello"))
	}
	// Output:
	// Hello World!
}

func ExamplePgConnect() {
	c := PgConnect("db", testDataHome())
	fmt.Println(c)
	// Output:
	// dbuser=postgres dbname=orc dbport=5432 dbhost=db.donaldww.com dbpassword=orakular sslmode=disable
}
