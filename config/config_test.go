package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBillItem(t *testing.T) {
	testItems := BillRenameItemList{}
	err := testItems.Unmarshal("")
	assert.Error(t, err)
}
