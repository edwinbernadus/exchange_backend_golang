package tests_data

import (
	"fiff_golang_draft/module_student"
	"strconv"
	"testing"
)

func TestGetTotal(t *testing.T) {
	var b = module_student.GetTotal(3, 4)
	var target = 7
	if b != target {

		var msg = "SALAH! harusnya " + strconv.FormatInt(int64(target), 10)
		t.Errorf(msg)
	}

}
