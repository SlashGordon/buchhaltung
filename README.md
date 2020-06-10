[![Build Status](https://travis-ci.org/SlashGordon/buchhaltung.svg?branch=master)](https://travis-ci.org/SlashGordon/buchhaltung)
[![Coverage Status](https://coveralls.io/repos/github/SlashGordon/buchhaltung/badge.svg?branch=master)](https://coveralls.io/github/SlashGordon/buchhaltung?branch=master)
# buchhaltung

## rename invoice

The renaming function reads the PDF and extracts variables for the renaming.

build binary with:

```shell
go get -d -v ./... && go build buchhaltung.go
```

Move all PDF invoices to a directory. At the moment we only support pure PDFs and not scans.

Create a json config file like:

```json
[
    {
        "outputname": "{number}_{company}.pdf",
        "identifyers":     {
            "number":  "Belegnummer\\s+(WF\\d{11})",
            "company": "(weinfreunde)"
        }
    },
    {
        "outputname": "{number}_{company}.pdf",
        "identifyers":     {
            "number":  "Rechnungsnr.:\\s+(F\\d{11})",
            "company": "(klarmobil)"
        }
    }
]
```

And run buchhaltung:
```shell
buchhaltung invoice -i /Users/test/bills -o /Users/test/bills/output -c /Users/test/
config_example.json
```