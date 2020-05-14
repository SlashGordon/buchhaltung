package pdf

import (
	"time"
	_ "unsafe" // for license

	"github.com/unidoc/unipdf/v3/common/license"
)

//go:linkname licenseKey github.com/unidoc/unipdf/v3/common/license.licenseKey
var licenseKey *license.LicenseKey

//InitLicense use community license
func InitLicense() {
	lk := license.LicenseKey{}
	lk.CustomerName = "community"
	lk.Tier = license.LicenseTierCommunity
	lk.CreatedAt = time.Now().UTC()
	lk.CreatedAtInt = lk.CreatedAt.Unix()
	licenseKey = &lk
}
