module github.com/Bar-Nik/back-template

go 1.22

toolchain go1.22.2

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.34.2-20240920164238-5a7b106cbb87.2
	github.com/Masterminds/squirrel v1.5.4
	github.com/felixge/httpsnoop v1.0.4
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/gorilla/mux v1.8.1
	github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus v1.0.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.1
	github.com/jmoiron/sqlx v1.3.5
	github.com/laher/mergefs v0.1.1
	github.com/lib/pq v1.10.9
	github.com/minio/minio-go/v7 v7.0.63
	github.com/mvrilo/go-redoc v0.1.4
	github.com/nats-io/nats.go v1.31.0
	github.com/o1egl/paseto/v2 v2.1.1
	github.com/ory/dockertest/v3 v3.10.0
	github.com/prometheus/client_golang v1.17.0
	github.com/rs/cors v1.10.1
	github.com/rs/xid v1.5.0
	github.com/samber/lo v1.38.1
	github.com/sipki-tech/database v0.2.11
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.56.0
	go.uber.org/mock v0.3.0
	golang.org/x/crypto v0.28.0
	golang.org/x/sync v0.8.0
	google.golang.org/genproto/googleapis/api v0.0.0-20241015192408-796eee8c2d53
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.35.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	cloud.google.com/go v0.116.0 // indirect
	cloud.google.com/go/accessapproval v1.8.1 // indirect
	cloud.google.com/go/accesscontextmanager v1.9.1 // indirect
	cloud.google.com/go/aiplatform v1.68.0 // indirect
	cloud.google.com/go/analytics v0.25.1 // indirect
	cloud.google.com/go/apigateway v1.7.1 // indirect
	cloud.google.com/go/apigeeconnect v1.7.1 // indirect
	cloud.google.com/go/apigeeregistry v0.9.1 // indirect
	cloud.google.com/go/appengine v1.9.1 // indirect
	cloud.google.com/go/area120 v0.9.1 // indirect
	cloud.google.com/go/artifactregistry v1.15.1 // indirect
	cloud.google.com/go/asset v1.20.2 // indirect
	cloud.google.com/go/assuredworkloads v1.12.1 // indirect
	cloud.google.com/go/automl v1.14.1 // indirect
	cloud.google.com/go/baremetalsolution v1.3.1 // indirect
	cloud.google.com/go/batch v1.11.1 // indirect
	cloud.google.com/go/beyondcorp v1.1.1 // indirect
	cloud.google.com/go/bigquery v1.63.1 // indirect
	cloud.google.com/go/bigtable v1.33.0 // indirect
	cloud.google.com/go/billing v1.19.1 // indirect
	cloud.google.com/go/binaryauthorization v1.9.1 // indirect
	cloud.google.com/go/certificatemanager v1.9.1 // indirect
	cloud.google.com/go/channel v1.19.0 // indirect
	cloud.google.com/go/cloudbuild v1.18.0 // indirect
	cloud.google.com/go/clouddms v1.8.1 // indirect
	cloud.google.com/go/cloudtasks v1.13.1 // indirect
	cloud.google.com/go/compute v1.28.1 // indirect
	cloud.google.com/go/contactcenterinsights v1.15.0 // indirect
	cloud.google.com/go/container v1.40.0 // indirect
	cloud.google.com/go/containeranalysis v0.13.1 // indirect
	cloud.google.com/go/datacatalog v1.22.1 // indirect
	cloud.google.com/go/dataflow v0.10.1 // indirect
	cloud.google.com/go/dataform v0.10.1 // indirect
	cloud.google.com/go/datafusion v1.8.1 // indirect
	cloud.google.com/go/datalabeling v0.9.1 // indirect
	cloud.google.com/go/dataplex v1.19.1 // indirect
	cloud.google.com/go/dataproc/v2 v2.9.0 // indirect
	cloud.google.com/go/dataqna v0.9.1 // indirect
	cloud.google.com/go/datastore v1.19.0 // indirect
	cloud.google.com/go/datastream v1.11.1 // indirect
	cloud.google.com/go/deploy v1.23.0 // indirect
	cloud.google.com/go/dialogflow v1.58.0 // indirect
	cloud.google.com/go/dlp v1.19.0 // indirect
	cloud.google.com/go/documentai v1.34.0 // indirect
	cloud.google.com/go/domains v0.10.1 // indirect
	cloud.google.com/go/edgecontainer v1.3.1 // indirect
	cloud.google.com/go/errorreporting v0.3.1 // indirect
	cloud.google.com/go/essentialcontacts v1.7.1 // indirect
	cloud.google.com/go/eventarc v1.14.1 // indirect
	cloud.google.com/go/filestore v1.9.1 // indirect
	cloud.google.com/go/firestore v1.17.0 // indirect
	cloud.google.com/go/functions v1.19.1 // indirect
	cloud.google.com/go/gkebackup v1.6.1 // indirect
	cloud.google.com/go/gkeconnect v0.11.1 // indirect
	cloud.google.com/go/gkehub v0.15.1 // indirect
	cloud.google.com/go/gkemulticloud v1.4.0 // indirect
	cloud.google.com/go/gsuiteaddons v1.7.1 // indirect
	cloud.google.com/go/iam v1.2.1 // indirect
	cloud.google.com/go/iap v1.10.1 // indirect
	cloud.google.com/go/ids v1.5.1 // indirect
	cloud.google.com/go/iot v1.8.1 // indirect
	cloud.google.com/go/kms v1.20.0 // indirect
	cloud.google.com/go/language v1.14.1 // indirect
	cloud.google.com/go/lifesciences v0.10.1 // indirect
	cloud.google.com/go/logging v1.11.0 // indirect
	cloud.google.com/go/longrunning v0.6.1 // indirect
	cloud.google.com/go/managedidentities v1.7.1 // indirect
	cloud.google.com/go/maps v1.14.0 // indirect
	cloud.google.com/go/mediatranslation v0.9.1 // indirect
	cloud.google.com/go/memcache v1.11.1 // indirect
	cloud.google.com/go/metastore v1.14.1 // indirect
	cloud.google.com/go/monitoring v1.21.1 // indirect
	cloud.google.com/go/networkconnectivity v1.15.1 // indirect
	cloud.google.com/go/networkmanagement v1.14.1 // indirect
	cloud.google.com/go/networksecurity v0.10.1 // indirect
	cloud.google.com/go/notebooks v1.12.1 // indirect
	cloud.google.com/go/optimization v1.7.1 // indirect
	cloud.google.com/go/orchestration v1.11.0 // indirect
	cloud.google.com/go/orgpolicy v1.14.0 // indirect
	cloud.google.com/go/osconfig v1.14.1 // indirect
	cloud.google.com/go/oslogin v1.14.1 // indirect
	cloud.google.com/go/phishingprotection v0.9.1 // indirect
	cloud.google.com/go/policytroubleshooter v1.11.1 // indirect
	cloud.google.com/go/privatecatalog v0.10.1 // indirect
	cloud.google.com/go/pubsub v1.44.0 // indirect
	cloud.google.com/go/pubsublite v1.8.2 // indirect
	cloud.google.com/go/recaptchaenterprise/v2 v2.17.2 // indirect
	cloud.google.com/go/recommendationengine v0.9.1 // indirect
	cloud.google.com/go/recommender v1.13.1 // indirect
	cloud.google.com/go/redis v1.17.1 // indirect
	cloud.google.com/go/resourcemanager v1.10.1 // indirect
	cloud.google.com/go/resourcesettings v1.8.1 // indirect
	cloud.google.com/go/retail v1.19.0 // indirect
	cloud.google.com/go/run v1.6.0 // indirect
	cloud.google.com/go/scheduler v1.11.1 // indirect
	cloud.google.com/go/secretmanager v1.14.1 // indirect
	cloud.google.com/go/security v1.18.1 // indirect
	cloud.google.com/go/securitycenter v1.35.1 // indirect
	cloud.google.com/go/servicedirectory v1.12.1 // indirect
	cloud.google.com/go/shell v1.8.1 // indirect
	cloud.google.com/go/spanner v1.70.0 // indirect
	cloud.google.com/go/speech v1.25.1 // indirect
	cloud.google.com/go/storagetransfer v1.11.1 // indirect
	cloud.google.com/go/talent v1.7.1 // indirect
	cloud.google.com/go/texttospeech v1.8.1 // indirect
	cloud.google.com/go/tpu v1.7.1 // indirect
	cloud.google.com/go/trace v1.11.1 // indirect
	cloud.google.com/go/translate v1.12.1 // indirect
	cloud.google.com/go/video v1.23.1 // indirect
	cloud.google.com/go/videointelligence v1.12.1 // indirect
	cloud.google.com/go/vision/v2 v2.9.1 // indirect
	cloud.google.com/go/vmmigration v1.8.1 // indirect
	cloud.google.com/go/vmwareengine v1.3.1 // indirect
	cloud.google.com/go/vpcaccess v1.8.1 // indirect
	cloud.google.com/go/webrisk v1.10.1 // indirect
	cloud.google.com/go/websecurityscanner v1.7.1 // indirect
	cloud.google.com/go/workflows v1.13.1 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bufbuild/protovalidate-go v0.7.2 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/containerd/continuity v0.4.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/cli v24.0.7+incompatible // indirect
	github.com/docker/docker v24.0.7+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/cel-go v0.21.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/lyft/protoc-gen-star/v2 v2.0.4-0.20230330145011-496ad1ac90a4 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/nkeys v0.4.6 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runc v1.1.10 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/afero v1.10.0 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	go.opentelemetry.io/otel v1.31.0 // indirect
	go.opentelemetry.io/otel/metric v1.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.31.0 // indirect
	golang.org/x/exp v0.0.0-20240325151524-a685a6edb6d8 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da // indirect
	google.golang.org/genproto v0.0.0-20241015192408-796eee8c2d53 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
