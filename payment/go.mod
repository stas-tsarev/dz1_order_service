module github.com/stas-tsarev/dz1_order_service/payment

go 1.26

replace github.com/stas-tsarev/dz1_order_service/shared => ../shared

require (
	github.com/brianvoe/gofakeit/v7 v7.16.0
	github.com/google/uuid v1.6.0
	github.com/stas-tsarev/dz1_order_service/shared v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.12.1
	google.golang.org/grpc v1.83.2
)

require (
	github.com/stretchr/objx v0.5.3 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
