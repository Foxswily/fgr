package template

import (
	"os"
	"testing"
)

func Test_Read(t *testing.T) {
	data, err := Read("..\\test\\notexist.yaml")
	if err != nil && !os.IsNotExist(err) {
		t.Errorf("something error %s", err)
	}
	data, err = Read("..\\test\\test.yaml")
	if err != nil {
		t.Errorf("something error %s", err)
	}
	t.Logf("data %v", data)
}

func Test_Write(t *testing.T) {
	data, err := Read("..\\test\\d1.yaml")
	if err != nil {
		t.Errorf("something error %s", err)
	}
	t.Logf("data %v", data)
	err = Write(data)
	if err != nil {
		t.Errorf("write error %s", err)
	}
}
