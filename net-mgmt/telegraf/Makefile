# $FreeBSD$

PORTNAME=	telegraf
PORTVERSION=	0.10.2
CATEGORIES=	net-mgmt

MAINTAINER=	cheffo@freebsd-bg.org
COMMENT=	TIME-SERIES DATA COLLECTION

LICENSE=	MIT

BUILD_DEPENDS=	go>=1.5.0:${PORTSDIR}/lang/go

ONLY_FOR_ARCHS=	i386 amd64

USE_RC_SUBR=	telegraf

USE_GITHUB=	yes

GH_ACCOUNT=	influxdata:DEFAULT \
		Shopify:Shopify_sarama \
		Sirupsen:Sirupsen_logrus \
		amir:amir_raidman \
		armon:armon_go_metrics \
		aws:aws_aws_sdk_go \
		beorn7:beorn7_perks \
		boltdb:boltdb_bolt \
		cenkalti:cenkalti_backoff \
		dancannon:dancannon_gorethink \
		davecgh:davecgh_go_spew \
		eapache:eapache_go_resiliency,eapache_queue \
		fsouza:fsouza_go_dockerclient \
		go-ini:go_ini_ini \
		go-sql-driver:go_sql_driver_mysql \
		gogo:gogo_protobuf \
		golang:golang_protobuf,golang_snappy,golang_crypto,golang_text \
		gonuts:gonuts_go_shellquote \
		gorilla:gorilla_context,gorilla_mux \
		hailocab:hailocab_go_hostpool \
		hashicorp:hashicorp_go_msgpack,hashicorp_raft,hashicorp_raft_boltdb \
		influxdata:influxdata_config,influxdata_influxdb \
		jmespath:jmespath_go_jmespath \
		klauspost:klauspost_crc32 \
		lib:lib_pq \
		matttproud:matttproud_golang_protobuf_extensions \
		mreiferson:mreiferson_go_snappystream \
		naoina:naoina_go_stringutil,naoina_toml \
		nsqio:nsqio_go_nsq \
		pmezard:pmezard_go_difflib \
		prometheus:prometheus_client_golang,prometheus_client_model,prometheus_common,prometheus_procfs \
		samuel:samuel_go_zookeeper \
		shirou:shirou_gopsutil \
		soniah:soniah_gosnmp \
		streadway:streadway_amqp \
		stretchr:stretchr_objx,stretchr_testify \
		wvanbergen:wvanbergen_kafka,wvanbergen_kazoo_go \
		zensqlmonitor:zensqlmonitor_go_mssqldb \
		fatih:fatih_pool \
		go-mgo:go_mgo \
		go-yaml:go_yaml \
		go-mqtt:go_mqtt,go_uuid \
		cesanta:cesanta_goxnet

GH_PROJECT=	amqp:streadway_amqp \
		aws-sdk-go:aws_aws_sdk_go \
		backoff:cenkalti_backoff \
		bolt:boltdb_bolt \
		client_golang:prometheus_client_golang \
		client_model:prometheus_client_model \
		common:prometheus_common \
		config:influxdata_config \
		context:gorilla_context \
		crc32:klauspost_crc32 \
		go-difflib:pmezard_go_difflib \
		go-dockerclient:fsouza_go_dockerclient \
		go-hostpool:hailocab_go_hostpool \
		go-jmespath:jmespath_go_jmespath \
		go-metrics:armon_go_metrics \
		go-msgpack:hashicorp_go_msgpack \
		go-mssqldb:zensqlmonitor_go_mssqldb \
		go-nsq:nsqio_go_nsq \
		go-resiliency:eapache_go_resiliency \
		go-shellquote:gonuts_go_shellquote \
		go-snappystream:mreiferson_go_snappystream \
		go-spew:davecgh_go_spew \
		go-stringutil:naoina_go_stringutil \
		go-zookeeper:samuel_go_zookeeper \
		golang_protobuf_extensions:matttproud_golang_protobuf_extensions \
		gopsutil:shirou_gopsutil \
		gorethink:dancannon_gorethink \
		gosnmp:soniah_gosnmp \
		influxdb:influxdata_influxdb \
		ini:go_ini_ini \
		kafka:wvanbergen_kafka \
		kazoo-go:wvanbergen_kazoo_go \
		logrus:Sirupsen_logrus \
		mux:gorilla_mux \
		mysql:go_sql_driver_mysql \
		objx:stretchr_objx \
		perks:beorn7_perks \
		pq:lib_pq \
		procfs:prometheus_procfs \
		protobuf:gogo_protobuf \
		protobuf:golang_protobuf \
		queue:eapache_queue \
		raft-boltdb:hashicorp_raft_boltdb \
		raft:hashicorp_raft \
		raidman:amir_raidman \
		sarama:Shopify_sarama \
		snappy:golang_snappy \
		testify:stretchr_testify \
		toml:naoina_toml \
		crypto:golang_crypto \
		pool:fatih_pool \
		mgo:go_mgo \
		yaml:go_yaml \
		mqtt:go_mqtt \
		uuid:go_uuid \
		goxnet:cesanta_goxnet \
		text:golang_text

GH_TAGNAME=	4ba9bba:Shopify_sarama \
		be52937:Sirupsen_logrus \
		53c1b96:amir_raidman \
		345426c:armon_go_metrics \
		371c1c6:aws_aws_sdk_go \
		b965b61:beorn7_perks \
		2f846c3:boltdb_bolt \
		4dc7767:cenkalti_backoff \
		7d1af5b:dancannon_gorethink \
		5215b55:davecgh_go_spew \
		b86b1ec:eapache_go_resiliency \
		ded5959:eapache_queue \
		0099401:fsouza_go_dockerclient \
		193d1ec:go_ini_ini \
		267b128:go_sql_driver_mysql \
		82d16f7:gogo_protobuf \
		b982704:golang_protobuf \
		894fd46:golang_snappy \
		e842a11:gonuts_go_shellquote \
		1c83b3e:gorilla_context \
		26a6070:gorilla_mux \
		e80d13c:hailocab_go_hostpool \
		fa3f638:hashicorp_go_msgpack \
		057b893:hashicorp_raft \
		d1e82c1:hashicorp_raft_boltdb \
		dbabd86:influxdata_config \
		af1a22d:influxdata_influxdb \
		0b12d6b:jmespath_go_jmespath \
		81ac418:klauspost_crc32 \
		05333d3:lib_pq \
		d0c3fe8:matttproud_golang_protobuf_extensions \
		028eae7:mreiferson_go_snappystream \
		6b638e9:naoina_go_stringutil \
		7511716:naoina_toml \
		2118015:nsqio_go_nsq \
		792786c:pmezard_go_difflib \
		15006a7:prometheus_client_golang \
		fa8ad6f:prometheus_client_model \
		dba5e39:prometheus_common \
		406e5b7:prometheus_procfs \
		218e9c8:samuel_go_zookeeper \
		59b63de:shirou_gopsutil \
		b1b4f88:soniah_gosnmp \
		b4f3cea:streadway_amqp \
		1a9d0bb:stretchr_objx \
		9f9027f:stretchr_testify \
		1a8639a:wvanbergen_kafka \
		0f76871:wvanbergen_kazoo_go \
		ffe5510:zensqlmonitor_go_mssqldb \
		1f22c01:golang_crypto \
		cba550e:fatih_pool \
		03c9f3e:go_mgo \
		f7716cb:go_yaml \
		e02a4a4:go_mqtt \
		f9b8669:cesanta_goxnet \
		c67034b:golang_text \
		dee7705:go_uuid

SUB_FILES=	telegraf
SUB_LIST+=	TELEGRAF_PIDDIR=${TELEGRAF_PIDDIR}
PLIST_SUB=	TELEGRAF_PIDDIR=${TELEGRAF_PIDDIR}

TELEGRAF_PIDDIR=	/var/run/${PORTNAME}/

STRIP=		# stripping can break go binaries

post-patch:
	@${MKDIR} ${WRKSRC}/src/code.google.com/p
	@${MKDIR} ${WRKSRC}/src/github.com/go-mqtt
	@${MKDIR} ${WRKSRC}/src/github.com/naoina
	@${MKDIR} ${WRKSRC}/src/github.com/armon
	@${MKDIR} ${WRKSRC}/src/github.com/boltdb
	@${MKDIR} ${WRKSRC}/src/github.com/fatih
	@${MKDIR} ${WRKSRC}/src/github.com/gogo
	@${MKDIR} ${WRKSRC}/src/github.com/hashicorp
	@${MKDIR} ${WRKSRC}/src/github.com/jwilder
	@${MKDIR} ${WRKSRC}/src/github.com/influxdata
	@${MKDIR} ${WRKSRC}/src/github.com/influxdb
	@${MKDIR} ${WRKSRC}/src/github.com/stretchr
	@${MKDIR} ${WRKSRC}/src/github.com/davecgh
	@${MKDIR} ${WRKSRC}/src/github.com/pmezard
	@${MKDIR} ${WRKSRC}/src/github.com/Shopify
	@${MKDIR} ${WRKSRC}/src/github.com/amir
	@${MKDIR} ${WRKSRC}/src/github.com/aws
	@${MKDIR} ${WRKSRC}/src/github.com/fsouza
	@${MKDIR} ${WRKSRC}/src/github.com/go-sql-driver
	@${MKDIR} ${WRKSRC}/src/github.com/gorilla
	@${MKDIR} ${WRKSRC}/src/github.com/eapache
	@${MKDIR} ${WRKSRC}/src/github.com/go-ini
	@${MKDIR} ${WRKSRC}/src/github.com/golang
	@${MKDIR} ${WRKSRC}/src/github.com/gonuts
	@${MKDIR} ${WRKSRC}/src/github.com/jmespath
	@${MKDIR} ${WRKSRC}/src/github.com/klauspost
	@${MKDIR} ${WRKSRC}/src/github.com/lib
	@${MKDIR} ${WRKSRC}/src/github.com/nsqio
	@${MKDIR} ${WRKSRC}/src/github.com/prometheus
	@${MKDIR} ${WRKSRC}/src/github.com/beorn7
	@${MKDIR} ${WRKSRC}/src/github.com/matttproud
	@${MKDIR} ${WRKSRC}/src/github.com/mreiferson
	@${MKDIR} ${WRKSRC}/src/github.com/shirou
	@${MKDIR} ${WRKSRC}/src/github.com/streadway
	@${MKDIR} ${WRKSRC}/src/github.com/wvanbergen
	@${MKDIR} ${WRKSRC}/src/github.com/Sirupsen
	@${MKDIR} ${WRKSRC}/src/github.com/cenkalti
	@${MKDIR} ${WRKSRC}/src/github.com/dancannon
	@${MKDIR} ${WRKSRC}/src/github.com/hailocab
	@${MKDIR} ${WRKSRC}/src/github.com/samuel
	@${MKDIR} ${WRKSRC}/src/github.com/soniah
	@${MKDIR} ${WRKSRC}/src/github.com/zensqlmonitor
	@${MKDIR} ${WRKSRC}/src/gopkg.in/dancannon
	@${MKDIR} ${WRKSRC}/src/gopkg.in/fatih
	@${MKDIR} ${WRKSRC}/src/gopkg.in/mqtt.v0
	@${MKDIR} ${WRKSRC}/src/golang.org/x
	@${MKDIR} ${WRKSRC}/src/github.com/${GH_ACCOUNT}/${PORTNAME}
.for src in CHANGELOG.md CONFIGURATION.md CONTRIBUTING.md Godeps Godeps_windows \
	LICENSE LICENSE_OF_DEPENDENCIES.md Makefile README.md accumulator.go \
	agent build.py circle.yml cmd etc input.go internal metric.go \
	metric_test.go output.go plugins scripts testutil
	@${MV} ${WRKSRC}/${src} \
		${WRKSRC}/src/github.com/${GH_ACCOUNT}/${PORTNAME}
.endfor
	@${CP} -r ${WRKSRC_dancannon_gorethink} ${WRKSRC}/src/github.com/dancannon/gorethink
	@${CP} -r ${WRKSRC_go_mgo} ${WRKSRC}/src/gopkg.in/mgo.v2-unstable
	@${CP} -r ${WRKSRC_go_mqtt}/packets ${WRKSRC}/src/gopkg.in/mqtt.v0/packets
	@${CP} -r ${WRKSRC_influxdata_influxdb} ${WRKSRC}/src/github.com/influxdb/influxdb
	@${MV} ${WRKSRC_gorilla_mux} ${WRKSRC}/src/github.com/gorilla/mux
	@${MV} ${WRKSRC_gorilla_context} ${WRKSRC}/src/github.com/gorilla/context
	@${MV} ${WRKSRC_zensqlmonitor_go_mssqldb} ${WRKSRC}/src/github.com/zensqlmonitor/go-mssqldb
	@${MV} ${WRKSRC_soniah_gosnmp} ${WRKSRC}/src/github.com/soniah/gosnmp
	@${MV} ${WRKSRC_boltdb_bolt} ${WRKSRC}/src/github.com/boltdb/bolt
	@${MV} ${WRKSRC_cesanta_goxnet} ${WRKSRC}/src/golang.org/x/net
	@${MV} ${WRKSRC_golang_text} ${WRKSRC}/src/golang.org/x/text
	@${MV} ${WRKSRC_go_mqtt} ${WRKSRC}/src/github.com/go-mqtt/mqtt
	@${MV} ${WRKSRC_go_uuid} ${WRKSRC}/src/github.com/go-mqtt/uuid
	@${MV} ${WRKSRC_golang_crypto} ${WRKSRC}/src/golang.org/x/crypto
	@${MV} ${WRKSRC_gogo_protobuf} ${WRKSRC}/src/github.com/gogo/protobuf
	@${MV} ${WRKSRC_cenkalti_backoff} ${WRKSRC}/src/github.com/cenkalti/backoff
	@${MV} ${WRKSRC_hailocab_go_hostpool} ${WRKSRC}/src/github.com/hailocab/go-hostpool
	@${MV} ${WRKSRC_samuel_go_zookeeper} ${WRKSRC}/src/github.com/samuel/go-zookeeper
	@${MV} ${WRKSRC_influxdata_config} ${WRKSRC}/src/github.com/influxdata/config
	@${MV} ${WRKSRC_influxdata_influxdb} ${WRKSRC}/src/github.com/influxdata/influxdb
	@${MV} ${WRKSRC_hashicorp_raft_boltdb} ${WRKSRC}/src/github.com/hashicorp/raft-boltdb
	@${MV} ${WRKSRC_hashicorp_raft} ${WRKSRC}/src/github.com/hashicorp/raft
	@${MV} ${WRKSRC_naoina_toml} ${WRKSRC}/src/github.com/naoina/toml
	@${MV} ${WRKSRC_naoina_go_stringutil} ${WRKSRC}/src/github.com/naoina/go-stringutil
	@${MV} ${WRKSRC_stretchr_objx} ${WRKSRC}/src/github.com/stretchr/objx
	@${MV} ${WRKSRC_stretchr_testify} ${WRKSRC}/src/github.com/stretchr/testify
	@${MV} ${WRKSRC_davecgh_go_spew} ${WRKSRC}/src/github.com/davecgh/go-spew
	@${MV} ${WRKSRC_pmezard_go_difflib} ${WRKSRC}/src/github.com/pmezard/go-difflib
	@${MV} ${WRKSRC_Shopify_sarama} ${WRKSRC}/src/github.com/Shopify/sarama
	@${MV} ${WRKSRC_amir_raidman} ${WRKSRC}/src/github.com/amir/raidman
	@${MV} ${WRKSRC_armon_go_metrics} ${WRKSRC}/src/github.com/armon/go-metrics
	@${MV} ${WRKSRC_aws_aws_sdk_go} ${WRKSRC}/src/github.com/aws/aws-sdk-go
	@${MV} ${WRKSRC_fsouza_go_dockerclient} ${WRKSRC}/src/github.com/fsouza/go-dockerclient
	@${MV} ${WRKSRC_go_sql_driver_mysql} ${WRKSRC}/src/github.com/go-sql-driver/mysql
	@${MV} ${WRKSRC_eapache_go_resiliency} ${WRKSRC}/src/github.com/eapache/go-resiliency
	@${MV} ${WRKSRC_eapache_queue} ${WRKSRC}/src/github.com/eapache/queue
	@${MV} ${WRKSRC_go_ini_ini} ${WRKSRC}/src/github.com/go-ini/ini
	@${MV} ${WRKSRC_golang_protobuf} ${WRKSRC}/src/github.com/golang/protobuf
	@${MV} ${WRKSRC_golang_snappy} ${WRKSRC}/src/github.com/golang/snappy
	@${MV} ${WRKSRC_gonuts_go_shellquote} ${WRKSRC}/src/github.com/gonuts/go-shellquote
	@${MV} ${WRKSRC_hashicorp_go_msgpack} ${WRKSRC}/src/github.com/hashicorp/go-msgpack
	@${MV} ${WRKSRC_jmespath_go_jmespath} ${WRKSRC}/src/github.com/jmespath/go-jmespath
	@${MV} ${WRKSRC_klauspost_crc32} ${WRKSRC}/src/github.com/klauspost/crc32
	@${MV} ${WRKSRC_lib_pq} ${WRKSRC}/src/github.com/lib/pq
	@${MV} ${WRKSRC_nsqio_go_nsq} ${WRKSRC}/src/github.com/nsqio/go-nsq
	@${MV} ${WRKSRC_prometheus_client_golang} ${WRKSRC}/src/github.com/prometheus/client_golang
	@${MV} ${WRKSRC_prometheus_common} ${WRKSRC}/src/github.com/prometheus/common
	@${MV} ${WRKSRC_prometheus_client_model}  ${WRKSRC}/src/github.com/prometheus/client_model
	@${MV} ${WRKSRC_prometheus_procfs}  ${WRKSRC}/src/github.com/prometheus/procfs
	@${MV} ${WRKSRC_beorn7_perks} ${WRKSRC}/src/github.com/beorn7/perks
	@${MV} ${WRKSRC_matttproud_golang_protobuf_extensions} ${WRKSRC}/src/github.com/matttproud/golang_protobuf_extensions
	@${MV} ${WRKSRC_mreiferson_go_snappystream} ${WRKSRC}/src/github.com/mreiferson/go-snappystream
	@${MV} ${WRKSRC_shirou_gopsutil} ${WRKSRC}/src/github.com/shirou/gopsutil
	@${MV} ${WRKSRC_streadway_amqp} ${WRKSRC}/src/github.com/streadway/amqp
	@${MV} ${WRKSRC_wvanbergen_kafka} ${WRKSRC}/src/github.com/wvanbergen/kafka
	@${MV} ${WRKSRC_wvanbergen_kazoo_go} ${WRKSRC}/src/github.com/wvanbergen/kazoo-go
	@${MV} ${WRKSRC_dancannon_gorethink} ${WRKSRC}/src/gopkg.in/dancannon/gorethink.v1
	@${MV} ${WRKSRC_fatih_pool} ${WRKSRC}/src/gopkg.in/fatih/pool.v2
	@${MV} ${WRKSRC_go_mgo} ${WRKSRC}/src/gopkg.in/mgo.v2
	@${MV} ${WRKSRC_go_yaml} ${WRKSRC}/src/gopkg.in/yaml.v2
	@${MV} ${WRKSRC_Sirupsen_logrus} ${WRKSRC}/src/github.com/Sirupsen/logrus

do-build:
		@cd ${WRKSRC}/src/github.com/influxdata/${PORTNAME} &&  ${SETENV} GOPATH=${WRKSRC} go install ./...

do-install:
		${MKDIR} ${STAGEDIR}${TELEGRAF_PIDDIR}

post-install:
		${INSTALL_PROGRAM} ${WRKSRC}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}
		${INSTALL_DATA} \
			${WRKSRC}/src/github.com/${GH_ACCOUNT}/${GH_PROJECT}/etc/${PORTNAME}.conf  \
			${STAGEDIR}${PREFIX}/etc/${PORTNAME}.conf.sample

.include <bsd.port.mk>
