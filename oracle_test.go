// Copyright 2020 by Donald Wilson. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.
package oracle

import (
	"fmt"
	"os"
)

func ExampleProgramName() {
	fmt.Println(ProgramName())
	// Output:
	// ___oracle_test_go
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

func ExampleNewViper() {
	v := NewViper("oracle", os.Getenv("HOME")+"/.config/oracle")
	fmt.Println(v.GetString("hello"))
	// Output:
	// Hello World!
}
