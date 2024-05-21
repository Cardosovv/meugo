// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p

func _[P int]() {
	type A = P
	_ = A(0) // don't crash with this conversion
}
