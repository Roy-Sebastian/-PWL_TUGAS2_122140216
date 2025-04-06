module orchestrator

go 1.24.1

require (
	google.golang.org/grpc v1.71.1
	google.golang.org/protobuf v1.36.6
	order-service v0.0.0-00010101000000-000000000000
	payment-service v0.0.0-00010101000000-000000000000
	shipping-service v0.0.0-00010101000000-000000000000
)

replace order-service => ../order-service

replace payment-service => ../payment-service

replace shipping-service => ../shipping-service

require (
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
)
