module gitlab.unanet.io/devops/eve-sch

go 1.14

replace gitlab.unanet.io/devops/eve => /Users/centzi/code/devops/eve

require (
	github.com/aws/aws-sdk-go v1.25.41
	github.com/golang/protobuf v1.4.0 // indirect
	github.com/hashicorp/vault v1.4.0 // indirect
	github.com/hashicorp/vault/api v1.0.5-0.20200317185738-82f498082f02
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.4.0
	gitlab.unanet.io/devops/eve v0.0.0-20200503233403-af7f7671826f
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20200429183012-4b2356b1ed79 // indirect
	golang.org/x/net v0.0.0-20200501053045-e0ff5e5a1de5 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v0.18.2
)
