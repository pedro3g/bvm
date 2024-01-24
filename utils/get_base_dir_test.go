package utils_test

import (
	"testing"

	"github.com/pedro3g/bvm/utils"
)

func TestGetBaseDir(t *testing.T) {
	_, err := utils.GetBaseDir()

	if err != nil {
		t.Errorf("Expect for the BVM_DIR variable to exist. %f", err)
	}
}
