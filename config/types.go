package config

type BillRenameItem struct {
	Identifyers           []string `json:"identifyers"`
	BillNumberIdentifyers []string `json:"bill_number_identifyers"`
	Rename                string   `json:"rename"`
}

type BillRenameItemList []BillRenameItem
