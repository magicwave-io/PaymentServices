module paidpiper.com/payment-gateway

go 1.20

require (
	github.com/bobesa/go-domain-util v0.0.0-20190911083921-4033b5f7dd89
	github.com/deepmap/oapi-codegen v1.6.1
	github.com/go-chi/chi/v5 v5.0.2
	github.com/go-errors/errors v0.0.0-20150906023321-a41850380601
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.1.1
	github.com/gorilla/handlers v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/ipfs/go-cid v0.4.1
	github.com/mattn/go-sqlite3 v1.14.7
	github.com/miekg/dns v1.1.61
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/xid v1.2.1
	github.com/sirupsen/logrus v1.6.0
	github.com/stellar/go v0.0.0-20210324164845-827227e3edd3
	github.com/stretchr/testify v1.5.1
	github.com/thoas/go-funk v0.9.3
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	go.opentelemetry.io/otel v0.4.2
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.4.2
	golang.org/x/net v0.26.0
	google.golang.org/grpc v1.27.1
	paidpiper.com/provider-service v0.0.0
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/apache/thrift v0.13.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gin-contrib/multitemplate v0.0.0-20210428235909-8a2f6dd269a0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.7.1 // indirect
	github.com/go-chi/chi v4.0.3+incompatible // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/ilyakaznacheev/cleanenv v1.2.5 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/klauspost/cpuid/v2 v2.0.4 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/manucorporat/sse v0.0.0-20160126180136-ee05b128a739 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiformats/go-base32 v0.0.3 // indirect
	github.com/multiformats/go-base36 v0.1.0 // indirect
	github.com/multiformats/go-multibase v0.0.3 // indirect
	github.com/multiformats/go-multihash v0.0.15 // indirect
	github.com/multiformats/go-varint v0.0.6 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/segmentio/go-loggly v0.5.1-0.20171222203950-eb91657e62b2 // indirect
	github.com/stellar/go-xdr v0.0.0-20201028102745-f80a23dac78a // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
	google.golang.org/api v0.20.0 // indirect
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03 // indirect
	google.golang.org/protobuf v1.26.0-rc.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	olympos.io/encoding/edn v0.0.0-20200308123125-93e3b8dd0e24 // indirect
)

replace paidpiper.com/provider-service => ./../payment-gateway-webui
