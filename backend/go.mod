module github.com/lyft/clutch/backend

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Masterminds/squirrel v1.5.0
	github.com/aws/aws-sdk-go-v2 v1.9.0
	github.com/aws/aws-sdk-go-v2/config v1.7.0
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.12.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.5.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.15.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.6.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.14.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.7.0
	github.com/aws/smithy-go v1.8.0
	github.com/bradleyfalzon/ghinstallation v1.1.1
	github.com/bufbuild/buf v0.37.0
	github.com/cactus/go-statsd-client/statsd v0.0.0-20200623234511-94959e3146b2
	github.com/coreos/go-oidc/v3 v3.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/envoyproxy/go-control-plane v0.9.9
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/fullstorydev/grpcurl v1.8.2
	github.com/go-git/go-billy/v5 v5.3.1
	github.com/go-git/go-git/v5 v5.4.2
	github.com/gobwas/glob v0.2.3
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/golang/protobuf v1.5.2
	github.com/google/go-github/v29 v29.0.3 // indirect
	github.com/google/go-github/v37 v37.0.0
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/iancoleman/strcase v0.2.0
	github.com/jhump/protoreflect v1.9.0
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.10.2
	github.com/m3db/prometheus_client_golang v0.8.1 // indirect
	github.com/m3db/prometheus_client_model v0.1.0 // indirect
	github.com/m3db/prometheus_common v0.1.0 // indirect
	github.com/m3db/prometheus_procfs v0.8.1 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/shurcooL/githubv4 v0.0.0-20210725200734-83ba7b4c9228
	github.com/shurcooL/graphql v0.0.0-20200928012149-18c5c3165e3a
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546
	github.com/slack-go/slack v0.9.4
	github.com/stretchr/testify v1.7.0
	github.com/twmb/murmur3 v1.1.5 // indirect
	github.com/uber-go/tally v3.4.2+incompatible
	go.uber.org/zap v1.19.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf
	golang.org/x/oauth2 v0.0.0-20210628180205-a41e5a781914
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210921065528-437939a70204 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.4 // indirect
	google.golang.org/genproto v0.0.0-20210921142501-181ce0d877f6
	google.golang.org/grpc v1.40.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/square/go-jose.v2 v2.6.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
)
