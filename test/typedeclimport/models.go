package typedeclimport

import subpkg "github.com/anyproto/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
