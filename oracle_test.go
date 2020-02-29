// Copyright 2020 by Donald Wilson. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.
package oracle

import (
	"fmt"
)

func ExampleProgramName() {
	fmt.Println(ProgramName())
	// Output:
	// ___go_test_github_com_donaldww_go_oracle
}

func ExampleThisIPAddr() {
	fmt.Println(ThisIPAddr())
	// Output:
	// 192.168.0.11
}
