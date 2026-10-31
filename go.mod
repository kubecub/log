module github.com/kubecub/log

go 1.20

require (
	github.com/json-iterator/go v1.1.12
	github.com/kubecub/log/logrus v0.0.0
	github.com/sirupsen/logrus v1.9.4
	go.uber.org/zap v1.27.1
)

require (
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/spf13/pflag v1.0.5
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	k8s.io/klog v1.0.0
)

replace github.com/kubecub/log/logrus => ./logrus
