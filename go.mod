module github.com/autotest-plan/controller

go 1.22.0

require github.com/autotest-plan/rpcdefine v0.0.0

require github.com/autotest-plan/log v0.0.0

require github.com/autotest-plan/errors v0.0.0

require (
	github.com/golang/protobuf v1.5.3 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240205150955-31a09d347014 // indirect
	google.golang.org/grpc v1.61.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace github.com/autotest-plan/rpcdefine => ../rpc-define

replace github.com/autotest-plan/log => ../log

replace github.com/autotest-plan/errors => ../errors