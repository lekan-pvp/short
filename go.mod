module github.com/lekan-pvp/short

go 1.17

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/go-chi/chi v1.5.4
	github.com/google/uuid v1.3.0
	github.com/itchyny/base58-go v0.2.0
	github.com/jackc/pgerrcode v0.0.0-20201024163028-a0d42d470451
	github.com/lib/pq v1.10.4
	github.com/stretchr/testify v1.7.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/tools v0.1.10
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	honnef.co/go/tools v0.2.2
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20220420153159-1850ba15e1be // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20220421151946-72621c1f0bd3 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.45.0
