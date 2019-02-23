package staticUtil

import (
	"testing"
	"time"
)

func mockTimeNow() {
	timeNow = func() time.Time {
		loc, _ := time.LoadLocation("Europe/Berlin")
		return time.Date(2018, 3, 30, 12, 20, 20, 20, loc)
	}
}

func TestGenerateDatePath(t *testing.T) {
	mockTimeNow()
	actual := GenerateDatePath()
	expected := "2018/3/30/"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestGetDate(t *testing.T) {
	mockTimeNow()
	actual := GetDate()
	expected := "2018-3-30"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestDirExists(t *testing.T) {
	actual := DirExists("path/to/nowhere")
	expected := false

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}

	actual = DirExists("testResources")
	expected = true

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}

	actual = DirExists("")
	expected = false

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
