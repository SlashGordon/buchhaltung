package types

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBillItem(t *testing.T) {
	jsonBlob := []byte(`
	[{"identifyers": {"name": "Peter", "age": "22"}, "outputname": "name_age.pdf"}, {"identifyers": {"name": "Klaus", "age": "55"}, "outputname": "name_age.pdf"}]`)
	testWriteItems := BillRenameItemList{}
	err := json.Unmarshal(jsonBlob, &testWriteItems)
	assert.NoError(t, err)
	testJSON, _ := json.Marshal(testWriteItems)
	err = ioutil.WriteFile("output.json", testJSON, 0644)
	assert.NoError(t, err)
	testReadItems := BillRenameItemList{}
	err = testReadItems.Unmarshal("")
	assert.Error(t, err)
	err = testReadItems.Unmarshal("output.json")
	assert.NoError(t, err)
	assert.Equal(t, testWriteItems, testReadItems)
}
