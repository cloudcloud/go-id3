package frames

import "testing"

func TestOrphanFuncs(t *testing.T) {
	str := []byte{'a', 'b', 'c'}
	if GetStr(str) != "abc" {
		t.Error("Expected string response for GetStr()")
	}

	str = []byte{6, 3}
	tmp := GetDirectInt(str[0])
	if tmp != 6 {
		t.Errorf("Expected int response for GetDirectInt(), got [%#v]\n", tmp)
	}

	str = []byte{'6', '3'}
	tmp = GetInt(str)
	if tmp != 63 {
		t.Errorf("Expected [54], got [%d]\n", tmp)
	}

	str = []byte{'\x04', '\x23'}
	tmp = GetSize(str, 8)
	if tmp != 1059 {
		t.Errorf("Expected [1059], got [%d]\n", tmp)
	}

	str = []byte{'\x01'}
	bo := GetBoolBit(str[0], 0)
	if !bo {
		t.Errorf("Expected [true], got [%v]\n", bo)
	}
}

func TestAllVersion2FramesMap(t *testing.T) {
	for k, v := range Version22Frames {
		if v() == nil {
			t.Errorf("Expected [%s] to be Mapped", k)
		}
	}
}

func TestAllVersion3FramesMap(t *testing.T) {
	for k, v := range Version23Frames {
		if v() == nil {
			t.Errorf("Expected [%s] to be Mapped", k)
		}
	}
}

func TestAllVersion4FramesMap(t *testing.T) {
	for k, v := range Version24Frames {
		if v() == nil {
			t.Errorf("Expected [%s] to be Mapped", k)
		}
	}
}

func TestUnknownFrame(t *testing.T) {
	f := NewFrame("BOBB", "", Version3)
	if f != nil {
		t.Fatalf("Got %#v, expected <nil>", f)
	}
}

func TestBustedString(t *testing.T) {
	b := []byte("This shouldn\xf0\xf1t be valid")

	expected := "This shouldnt be valid"
	found := GetStr(b)
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestBaseGenFunc(t *testing.T) {
	f := Gen("TPE1", "", Version4)()

	expected := "TPE1"
	found := f.GetName()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
