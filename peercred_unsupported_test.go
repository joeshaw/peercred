// +build !linux

package peercred

import "testing"

func TestUnsupportedRead(t *testing.T) {
	cred, err := Read(nil)
	if cred != nil {
		t.Errorf("want nil, have %v", cred)
	}
	if err != errUnsupported {
		t.Errorf("want %v, have %v", errUnsupported, err)
	}
}
