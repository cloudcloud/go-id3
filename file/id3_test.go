package file

import (
	"os"
	"testing"

	"github.com/cloudcloud/go-id3/frames"
)

func TestBaseFile(t *testing.T) {
	if os.Getenv("INT") != "1" {
		t.SkipNow()
	}

	filename := os.Getenv("FILENAME")
	f := &File{Filename: filename, Debug: true}

	defer func() {
		if a := recover(); a != nil {
			t.Error(a)
		}
	}()

	f.Process()
	f.CleanUp()
}

func TestOrphanFuncs(t *testing.T) {
	str := []byte{'a', 'b', 'c'}
	if frames.GetStr(str) != "abc" {
		t.Error("Expected string response for GetStr()")
	}

	str = []byte{6, 3}
	tmp := frames.GetDirectInt(str[0])
	if tmp != 6 {
		t.Errorf("Expected int response for GetDirectInt(), got [%#v]\n", tmp)
	}

	str = []byte{'6', '3'}
	tmp = frames.GetInt(str)
	if tmp != 63 {
		t.Errorf("Expected [54], got [%d]\n", tmp)
	}
}
