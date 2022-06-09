module github.com/l-angel/tunnel

go 1.16

replace github.com/siddontang/go-mysql => github.com/go-mysql-org/go-mysql v1.3.0

require github.com/go-mysql-org/go-mysql v1.3.0

require github.com/apache/rocketmq-client-go/v2 v2.1.0-rc3

require (
	github.com/samuel/go-zookeeper v0.0.0-20201211165307-7117e9ea2414
	github.com/siddontang/go-log v0.0.0-20190221022429-1e957dd83bed // indirect
	go.uber.org/atomic v1.7.0
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	gopkg.in/yaml.v3 v3.0.1
)
