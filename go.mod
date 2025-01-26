module patrickbrosi.de/vdv4522gtfs

go 1.23.3

replace patrickbrosi.de/vdv452parser => ../vdv452parser

replace patrickbrosi.de/vdv452writer => ../vdv452writer

replace patrickbrosi.de/x10parser => ../x10parser

replace patrickbrosi.de/x10writer => ../x10writer

require (
	github.com/patrickbr/gtfsparser v0.0.0-20250109120112-0f616022f79d
	github.com/patrickbr/gtfswriter v0.0.0-20241126214321-b6c6255581e4
	github.com/spf13/pflag v1.0.5
	patrickbrosi.de/vdv452parser v0.0.0-00010101000000-000000000000
)

require (
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/pebbe/go-proj-4 v5.0.0+incompatible // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	patrickbrosi.de/x10parser v0.0.0-00010101000000-000000000000 // indirect
)
