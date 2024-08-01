// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2019, The GoGo Authors. All rights reserved.
// http://github.com/anyproto/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package setextensionbytes

import (
	"testing"

	"github.com/anyproto/protobuf/proto"
)

func TestSetextensionBytesTwiceBeforeGet(t *testing.T) {
	myExtendable := &MyExtendable{}

	foo := &Foo{IntFoo: 1}
	err := proto.SetUnsafeExtension(myExtendable, 2, foo)
	if err != nil {
		t.Error(err)
	}
	gotFoo1, err := proto.GetUnsafeExtension(myExtendable, 2)
	if err != nil {
		t.Error(err)
	}

	gotFoo1.(*Foo).IntFoo = 10

	err = proto.SetUnsafeExtension(myExtendable, 2, gotFoo1)
	if err != nil {
		t.Error(err)
	}
	gotFoo2, err := proto.GetUnsafeExtension(myExtendable, 2)
	if err != nil {
		t.Error(err)
	}

	if gotFoo2.(*Foo).IntFoo != 10 {
		t.Fatalf("got %d want %d", gotFoo2.(*Foo).IntFoo, 10)
	}
}
