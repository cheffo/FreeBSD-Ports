# Created by: Stefan Lambrev <cheffo@freebsd-bg.org>
# $FreeBSD$

PORTNAME=	chronograf
DISTVERSION=	1.7.15
CATEGORIES=	net-mgmt

MAINTAINER=	cheffo@FreeBSD.org
COMMENT=	Time-series data collection

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

ONLY_FOR_ARCHS=	amd64 armv6 armv7 i386

USES=		gmake go python:build

BUILD_DEPENDS=	yarn:www/yarn-node10 \
		npm:www/npm-node10

USE_RC_SUBR=	chronograf

USE_GITHUB=	yes

GH_ACCOUNT=	influxdata
GH_SUBDIR=	src/github.com/${GH_ACCOUNT_DEFAULT}/${GH_PROJECT_DEFAULT}

GH_TUPLE=	\
		BurntSushi:toml:v0.3.1:burntsushi_toml/vendor/github.com/BurntSushi/toml \
		Masterminds:semver:v1.4.2:masterminds_semver/vendor/github.com/Masterminds/semver \
		NYTimes:gziphandler:6710af535839:nytimes_gziphandler/vendor/github.com/NYTimes/gziphandler \
		Sirupsen:logrus:3ec0642a7fb6:sirupsen_logrus/vendor/github.com/Sirupsen/logrus \
		alecthomas:kingpin:v2.2.6:alecthomas_kingpin/vendor/github.com/alecthomas/kingpin \
		alecthomas:template:a0175ee3bccc:alecthomas_template/vendor/github.com/alecthomas/template \
		alecthomas:units:2efee857e7cf:alecthomas_units/vendor/github.com/alecthomas/units \
		apache:arrow:af6fa24be0db:apache_arrow/vendor/github.com/apache/arrow \
		apex:log:v1.1.0:apex_log/vendor/github.com/apex/log \
		aws:aws-sdk-go:v1.16.18:aws_aws_sdk_go/vendor/github.com/aws/aws-sdk-go \
		blakesmith:ar:8bd4349a67f2:blakesmith_ar/vendor/github.com/blakesmith/ar \
		boltdb:bolt:5cc10bbbc5c1:boltdb_bolt/vendor/github.com/boltdb/bolt \
		bouk:httprouter:ee8b3818a7f5:bouk_httprouter/vendor/github.com/bouk/httprouter \
		caarlos0:ctrlc:v1.0.0:caarlos0_ctrlc/vendor/github.com/caarlos0/ctrlc \
		campoy:unique:88950e537e7e:campoy_unique/vendor/github.com/campoy/unique \
		census-instrumentation:opencensus-go:v0.22.0:census_instrumentation_opencensus_go/vendor/go.opencensus.io \
		cespare:xxhash:v1.1.0:cespare_xxhash/vendor/github.com/cespare/xxhash \
		dgrijalva:jwt-go:24c63f56522a:dgrijalva_jwt_go/vendor/github.com/dgrijalva/jwt-go \
		dustin:go-humanize:bb3d318650d4:dustin_go_humanize/vendor/github.com/dustin/go-humanize \
		eclipse:paho.mqtt.golang:v1.2.0:eclipse_paho_mqtt_golang/vendor/github.com/eclipse/paho.mqtt.golang \
		elazarl:go-bindata-assetfs:9a6736ed45b4:elazarl_go_bindata_assetfs/vendor/github.com/elazarl/go-bindata-assetfs \
		emirpasic:gods:v1.9.0:emirpasic_gods/vendor/github.com/emirpasic/gods \
		fatih:color:v1.7.0:fatih_color/vendor/github.com/fatih/color \
		go-sql-driver:mysql:v1.4.1:go_sql_driver_mysql/vendor/github.com/go-sql-driver/mysql \
		go-warnings:warnings:v0.1.2:go_warnings_warnings/vendor/gopkg.in/warnings.v0 \
		go-yaml:yaml:v2.2.2:go_yaml_yaml/vendor/gopkg.in/yaml.v2 \
		gogo:protobuf:v1.1.1:gogo_protobuf/vendor/github.com/gogo/protobuf \
		golang:appengine:v1.6.1:golang_appengine/vendor/google.golang.org/appengine \
		golang:crypto:f99c8df09eb5:golang_crypto/vendor/golang.org/x/crypto \
		golang:net:3b0461eec859:golang_net/vendor/golang.org/x/net \
		golang:oauth2:0f29369cfe45:golang_oauth2/vendor/golang.org/x/oauth2 \
		golang:protobuf:v1.3.2:golang_protobuf/vendor/github.com/golang/protobuf \
		golang:sync:112230192c58:golang_sync/vendor/golang.org/x/sync \
		golang:sys:c5567b49c5d0:golang_sys/vendor/golang.org/x/sys \
		golang:text:v0.3.2:golang_text/vendor/golang.org/x/text \
		golang:tools:6cdbf07be9d0:golang_tools/vendor/golang.org/x/tools \
		google:go-cmp:v0.3.0:google_go_cmp/vendor/github.com/google/go-cmp \
		google:go-genproto:c506a9f90610:google_go_genproto/vendor/google.golang.org/genproto \
		google:go-github:v17.0.0:google_go_github/vendor/github.com/google/go-github \
		google:go-querystring:v1.0.0:google_go_querystring/vendor/github.com/google/go-querystring \
		googleapis:gax-go:v2.0.5:googleapis_gax_go/vendor/github.com/googleapis/gax-go \
		googleapis:google-api-go-client:v0.7.0:googleapis_google_api_go_client/vendor/google.golang.org/api \
		googleapis:google-cloud-go:v0.43.0:googleapis_google_cloud_go/vendor/cloud.google.com/go \
		goreleaser:goreleaser:v0.97.0:goreleaser_goreleaser/vendor/github.com/goreleaser/goreleaser \
		goreleaser:nfpm:v0.9.7:goreleaser_nfpm/vendor/github.com/goreleaser/nfpm \
		grpc:grpc-go:v1.21.1:grpc_grpc_go/vendor/google.golang.org/grpc \
		hashicorp:golang-lru:v0.5.1:hashicorp_golang_lru/vendor/github.com/hashicorp/golang-lru \
		imdario:mergo:v0.3.6:imdario_mergo/vendor/github.com/imdario/mergo \
		inconshreveable:mousetrap:v1.0.0:inconshreveable_mousetrap/vendor/github.com/inconshreveable/mousetrap \
		influxdata:changelog:v1.1.0:influxdata_changelog/vendor/github.com/influxdata/changelog \
		influxdata:flux:v0.50.2:influxdata_flux/vendor/github.com/influxdata/flux \
		influxdata:influxdb:v1.1.5:influxdata_influxdb/vendor/github.com/influxdata/influxdb \
		influxdata:kapacitor:v1.5.0:influxdata_kapacitor/vendor/github.com/influxdata/kapacitor \
		influxdata:line-protocol:32c6aa80de5e:influxdata_line_protocol/vendor/github.com/influxdata/line-protocol \
		influxdata:tdigest:bf2b5ad3c0a9:influxdata_tdigest/vendor/github.com/influxdata/tdigest \
		influxdata:usage-client:6d3895376368:influxdata_usage_client/vendor/github.com/influxdata/usage-client \
		jbenet:go-context:d14ea06fba99:jbenet_go_context/vendor/github.com/jbenet/go-context \
		jessevdk:go-flags:v1.4.0:jessevdk_go_flags/vendor/github.com/jessevdk/go-flags \
		jmespath:go-jmespath:c2b33e8439af:jmespath_go_jmespath/vendor/github.com/jmespath/go-jmespath \
		kamilsk:retry:495c1d672c93:kamilsk_retry/vendor/github.com/kamilsk/retry \
		kevinburke:ssh_config:81db2a75821e:kevinburke_ssh_config/vendor/github.com/kevinburke/ssh_config \
		lestrrat-go:jwx:v0.9.0:lestrrat_go_jwx/vendor/github.com/lestrrat-go/jwx \
		lib:pq:v1.0.0:lib_pq/vendor/github.com/lib/pq \
		mattn:go-colorable:v0.0.9:mattn_go_colorable/vendor/github.com/mattn/go-colorable \
		mattn:go-isatty:v0.0.4:mattn_go_isatty/vendor/github.com/mattn/go-isatty \
		mattn:go-zglob:v0.0.1:mattn_go_zglob/vendor/github.com/mattn/go-zglob \
		matttproud:golang_protobuf_extensions:v1.0.1:matttproud_golang_protobuf_extensions/vendor/github.com/matttproud/golang_protobuf_extensions \
		microcosm-cc:bluemonday:v1.0.1:microcosm_cc_bluemonday/vendor/github.com/microcosm-cc/bluemonday \
		mitchellh:go-homedir:v1.1.0:mitchellh_go_homedir/vendor/github.com/mitchellh/go-homedir \
		opentracing:opentracing-go:v1.0.2:opentracing_opentracing_go/vendor/github.com/opentracing/opentracing-go \
		pelletier:go-buffruneio:v0.2.0:pelletier_go_buffruneio/vendor/github.com/pelletier/go-buffruneio \
		pkg:errors:v0.8.1:pkg_errors/vendor/github.com/pkg/errors \
		prometheus:client_model:fd36f4220a90:prometheus_client_model/vendor/github.com/prometheus/client_model \
		prometheus:common:v0.6.0:prometheus_common/vendor/github.com/prometheus/common \
		satori:go.uuid:v1.2.0:satori_go_uuid/vendor/github.com/satori/go.uuid \
		segmentio:kafka-go:v0.2.0:segmentio_kafka_go/vendor/github.com/segmentio/kafka-go \
		sergi:go-diff:v1.0.0:sergi_go_diff/vendor/github.com/sergi/go-diff \
		spf13:cobra:v0.0.3:spf13_cobra/vendor/github.com/spf13/cobra \
		spf13:pflag:v1.0.3:spf13_pflag/vendor/github.com/spf13/pflag \
		src-d:gcfg:v1.4.0:src_d_gcfg/vendor/github.com/src-d/gcfg \
		src-d:go-billy:v4.2.1:src_d_go_billy/vendor/gopkg.in/src-d/go-billy.v4 \
		src-d:go-git:v4.8.1:src_d_go_git/vendor/gopkg.in/src-d/go-git.v4 \
		tylerb:graceful:v1.2.15:tylerb_graceful/vendor/github.com/tylerb/graceful \
		uber-go:atomic:v1.3.2:uber_go_atomic/vendor/go.uber.org/atomic \
		uber-go:multierr:v1.1.0:uber_go_multierr/vendor/go.uber.org/multierr \
		uber-go:zap:v1.9.1:uber_go_zap/vendor/go.uber.org/zap \
		xanzy:ssh-agent:v0.2.0:xanzy_ssh_agent/vendor/github.com/xanzy/ssh-agent \
		parcel-bundler:parcel:e2ea843:parcel_bundler_parcel/vendor/github.com/parcel-bundler/parcel

		# Mirrors for the following packages are not currently known, please look them up and handle these tuples manually:
		#	::e561f6794a2a:group_name/vendor/honnef.co/go/tools

GO_BUILDFLAGS=	-ldflags "-X main.version=${DISTVERSION}"

USERS=		${CHRONOGRAF_USER}
GROUPS=		${CHRONOGRAF_GROUP}

SUB_FILES=	chronograf
SUB_LIST+=	CHRONOGRAF_PIDDIR=${CHRONOGRAF_PIDDIR}
PLIST_SUB=	CHRONOGRAF_PIDDIR=${CHRONOGRAF_PIDDIR}

#Use influxd user as it's aleary present in ports system
CHRONOGRAF_USER=	influxd
CHRONOGRAF_GROUP=	influxd

.include <bsd.port.pre.mk>

.if ${OPSYS} == "FreeBSD" && (${OSVERSION} < 1200000 || ${OSVERSION} >= 1200015)
SUB_LIST+=	CHRONOGRAF_LOGCMD="daemon"
.else
SUB_LIST+=	CHRONOGRAF_LOGCMD="logger"
.endif

do-build:
	cd ${WRKDIR} && ${SETENV} HOME=${WRKSRC}/src/github.com/influxdata/${PORTNAME}/ui yarn config set cache-folder ${WRKDIR}/.yarn_cache --offline
	cd ${WRKDIR} && ${SETENV} HOME=${WRKSRC}/src/github.com/influxdata/${PORTNAME}/ui yarn config set yarn-offline-mirror ${PORTSDIR}/distfiles/ --offline
	cd ${WRKSRC}/src/github.com/influxdata/${PORTNAME}/ui && ${SETENV} ${MAKE_ENV} PYTHON=${PYTHON_CMD} HOME=${WRKSRC}/src/github.com/influxdata/${PORTNAME}/ui \
		yarn --no-progress --no-emoji
	cd ${WRKSRC}/src/github.com/influxdata/${PORTNAME} && \
		${SETENV} ${MAKE_ENV} GOPATH=${WRKSRC} ${GMAKE} build
	cd ${WRKSRC}/src/github.com/influxdata/${PORTNAME} && \
		${SETENV} ${MAKE_ENV} GOPATH=${WRKSRC} GO111MODULE=off go install ./...
	${LN} -s ${WRKSRC}/bin ${WRKDIR}/bin

post-install:
	@${SED} -i .bak -e \
		"s|/var/lib/chronograf|/var/db/chronograf|g" \
		${WRKSRC}/etc/config.sample.toml
	${INSTALL_DATA} \
		${WRKSRC}/etc/config.sample.toml  \
		${STAGEDIR}${PREFIX}/etc/chronograf.conf.sample

.include <bsd.port.post.mk>
