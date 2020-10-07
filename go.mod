module github.com/taishi8117/go-bitflyer

go 1.14

//replace github.com/kkohtaka/go-bitflyer => github.com/taishi8117/go-bitflyer v0.0.0-20200929053129-a29a7866c2de
replace github.com/kkohtaka/go-bitflyer => ./

require (
	github.com/google/go-querystring v0.0.0-20170111101155-53e6ce116135
	github.com/gorilla/rpc v1.1.0
	github.com/gorilla/websocket v1.2.0
	github.com/json-iterator/go v1.1.10
	github.com/kkohtaka/go-bitflyer v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.8.0
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.9.0
)
