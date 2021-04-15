module github.com/Kubernetes-extension/preset-pod

go 1.14

require (
	github.com/Kubernetes-extension/utils v0.0.0-20210414024104-ec4fbf13c2da
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-gonic/gin v1.7.1
	go.uber.org/zap v1.16.0 // indirect
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v0.18.2
)

replace (
	k8s.io/api => k8s.io/api v0.17.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.3
	k8s.io/client-go => k8s.io/client-go v0.17.3
)
