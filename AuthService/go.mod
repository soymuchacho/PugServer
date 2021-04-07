module AuthService

go 1.15

require (
	AuthService/handle v0.0.0-00010101000000-000000000000
	github.com/Unknwon/goconfig v0.0.0-20200908083735-df7de6a44db8
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/gin-gonic/gin v1.6.3
)

replace AuthService/handle => ./handle
