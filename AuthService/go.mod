module AuthService

go 1.15

require (
	PugCommon v0.0.0-00010101000000-000000000000
	github.com/Unknwon/goconfig v0.0.0-20200908083735-df7de6a44db8
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/gin-gonic/gin v1.6.3
	github.com/smartystreets/goconvey v1.6.4 // indirect
)

replace PugCommon => ../PugCommon

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace AuthService/handle => ./handle

replace AuthService/service => ./service
