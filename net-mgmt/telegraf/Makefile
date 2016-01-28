# $FreeBSD$

PORTNAME=	telegraf
PORTVERSION=	0.10.0
DISTVERSIONPREFIX=	v
CATEGORIES=	net-mgmt

MAINTAINER=	cheffo@freebsd-bg.org
COMMENT=	TIME-SERIES DATA COLLECTION

LICENSE=	MIT

BUILD_DEPENDS=	go>=1.5.0:${PORTSDIR}/lang/go

ONLY_FOR_ARCHS=	i386 amd64

USE_RC_SUBR=	telegraf

USE_GITHUB=	yes
GH_ACCOUNT=	influxdb:DEFAULT \
		naoina:toml \
		naoina:go_stringutil \
		boltdb:bolt \
		cheffo:crypto \
		fatih:pool \
		gogo:gogo_protobuf \
		hashicorp:raft,raft_boltdb,go_msgpack \
		influxdb:influxdb \
		stretchr:objx,testify davecgh:go_spew pmezard:go_difflib Shopify:sarama amir:raidman armon:go_metrics aws:aws_sdk_go fsouza:go_dockerclient \
		go-sql-driver:mysql eapache:go_resiliency,queue go-ini:ini golang:golang_protobuf,golang_snappy gonuts:go_shellquote \
		jmespath:go_jmespath klauspost:crc32 lib:pq nsqio:go_nsq prometheus:client_golang,common,client_model,procfs beorn7:perks \
		matttproud:golang_protobuf_extensions \
		mreiferson:go_snappystream shirou:gopsutil streadway:amqp wvanbergen:kafka,kazoo_go dancannon:gorethink fatih:pool Sirupsen:logrus \
		cenkalti:backoff hailocab:go_hostpool samuel:go_zookeeper go-mgo:mgo go-yaml:yaml go-mqtt:mqtt,uuid cesanta:goxnet
GH_PROJECT=	bolt:bolt crypto:crypto \
		pool:pool \
		protobuf:gogo_protobuf \
		raft-boltdb:raft_boltdb raft:raft \
		toml:toml \
		influxdb:influxdb go-stringutil:go_stringutil \
		testify:testify go-spew:go_spew go-difflib:go_difflib objx:objx sarama:sarama raidman:raidman go-metrics:go_metrics aws-sdk-go:aws_sdk_go \
		go-dockerclient:go_dockerclient mysql:mysql go-resiliency:go_resiliency queue:queue ini:ini protobuf:golang_protobuf snappy:golang_snappy \
		go-shellquote:go_shellquote go-msgpack:go_msgpack go-jmespath:go_jmespath crc32:crc32 pq:pq go-nsq:go_nsq client_golang:client_golang \
		common:common perks:perks golang_protobuf_extensions:golang_protobuf_extensions go-snappystream:go_snappystream client_model:client_model \
		procfs:procfs gopsutil:gopsutil amqp:amqp kafka:kafka gorethink:gorethink pool:pool kazoo-go:kazoo_go logrus:logrus backoff:backoff \
		go-hostpool:go_hostpool go-zookeeper:go_zookeeper mgo:mgo yaml:yaml mqtt:mqtt uuid:uuid goxnet:goxnet
GH_TAGNAME=	b34b35e:bolt 7511716:toml 1351f93:crypto \
		e492fd3:gogo_protobuf \
		cba550e:pool d1e82c1:raft_boltdb \
		d136cd1:raft \
		6d3a860:influxdb 6b638e9:go_stringutil \
		f390dcf:testify 5215b55:go_spew 792786c:go_difflib 1a9d0bb:objx d37c73f:sarama 6a8e089:raidman 345426c:go_metrics 3ad0b07:aws_sdk_go \
		6fb38e6:go_dockerclient 809154b:mysql b86b1ec:go_resiliency ded5959:queue afbd495:ini 2402d76:golang_protobuf 723cc1e:golang_snappy \
		e842a11:go_shellquote fa3f638:go_msgpack c01cf91:go_jmespath 999f312:crc32 8ad2b29:pq 2118015:go_nsq 67994f1:client_golang 0a3005b:common \
		b965b61:perks d0c3fe8:golang_protobuf_extensions 028eae7:go_snappystream fa8ad6f:client_model 406e5b7:procfs 61dd483:gopsutil b4f3cea:amqp \
		1a8639a:kafka ff457ca:gorethink 0f76871:kazoo_go f7f79f7:logrus 4dc7767:backoff 50839ee:go_hostpool 218e9c8:go_zookeeper 03c9f3e:mgo \
		f7716cb:yaml e02a4a4:mqtt ca53cad:uuid f9b8669:goxnet

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
	@${MKDIR} ${WRKSRC}/src/github.com/influxdb
	@${MKDIR} ${WRKSRC}/src/github.com/stretchr
	@${MKDIR} ${WRKSRC}/src/github.com/davecgh
	@${MKDIR} ${WRKSRC}/src/github.com/pmezard
	@${MKDIR} ${WRKSRC}/src/github.com/Shopify
	@${MKDIR} ${WRKSRC}/src/github.com/amir
	@${MKDIR} ${WRKSRC}/src/github.com/aws
	@${MKDIR} ${WRKSRC}/src/github.com/fsouza
	@${MKDIR} ${WRKSRC}/src/github.com/go-sql-driver
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
	@${MKDIR} ${WRKSRC}/src/gopkg.in/dancannon
	@${MKDIR} ${WRKSRC}/src/gopkg.in/fatih
	@${MKDIR} ${WRKSRC}/src/gopkg.in/mqtt.v0
	@${MKDIR} ${WRKSRC}/src/golang.org/x
	@${MKDIR} ${WRKSRC}/src/github.com/${GH_ACCOUNT}/${PORTNAME}
.for src in CHANGELOG.md CONFIGURATION.md CONTRIBUTING.md Godeps LICENSE \
	README.md accumulator.go agent.go agent_test.go build.py circle.yml \
	cmd etc internal plugins scripts testutil
	@${MV} ${WRKSRC}/${src} \
		${WRKSRC}/src/github.com/${GH_ACCOUNT}/${PORTNAME}
.endfor
#Mongdb input doos not compile - kill it
#	@${RM} -rf  ${WRKSRC}/src/github.com/${GH_ACCOUNT}/${PORTNAME}/plugins/inputs/mongodb
	@${CP} -r ${WRKSRC_gorethink} ${WRKSRC}/src/github.com/dancannon/gorethink
	@${CP} -r ${WRKSRC_mgo} ${WRKSRC}/src/gopkg.in/mgo.v2-unstable
	@${CP} -r ${WRKSRC_mqtt}/packets ${WRKSRC}/src/gopkg.in/mqtt.v0/packets
	@${MV} ${WRKSRC_bolt} ${WRKSRC}/src/github.com/boltdb/bolt
	@${MV} ${WRKSRC_goxnet} ${WRKSRC}/src/golang.org/x/net
	@${MV} ${WRKSRC_mqtt} ${WRKSRC}/src/github.com/go-mqtt/mqtt
	@${MV} ${WRKSRC_uuid} ${WRKSRC}/src/github.com/go-mqtt/uuid
	@${MV} ${WRKSRC_crypto} ${WRKSRC}/src/golang.org/x/crypto
	@${MV} ${WRKSRC_gogo_protobuf} ${WRKSRC}/src/github.com/gogo/protobuf
	@${MV} ${WRKSRC_backoff} ${WRKSRC}/src/github.com/cenkalti/backoff
	@${MV} ${WRKSRC_go_hostpool} ${WRKSRC}/src/github.com/hailocab/go-hostpool
	@${MV} ${WRKSRC_go_zookeeper} ${WRKSRC}/src/github.com/samuel/go-zookeeper
	@${MV} ${WRKSRC_influxdb} ${WRKSRC}/src/github.com/influxdb/influxdb
	@${MV} ${WRKSRC_raft_boltdb} ${WRKSRC}/src/github.com/hashicorp/raft-boltdb
	@${MV} ${WRKSRC_raft} ${WRKSRC}/src/github.com/hashicorp/raft
	@${MV} ${WRKSRC_toml} ${WRKSRC}/src/github.com/naoina/toml
	@${MV} ${WRKSRC_go_stringutil} ${WRKSRC}/src/github.com/naoina/go-stringutil
	@${MV} ${WRKSRC_objx} ${WRKSRC}/src/github.com/stretchr/objx
	@${MV} ${WRKSRC_testify} ${WRKSRC}/src/github.com/stretchr/testify
	@${MV} ${WRKSRC_go_spew} ${WRKSRC}/src/github.com/davecgh/go-spew
	@${MV} ${WRKSRC_go_difflib} ${WRKSRC}/src/github.com/pmezard/go-difflib
	@${MV} ${WRKSRC_sarama} ${WRKSRC}/src/github.com/Shopify/sarama
	@${MV} ${WRKSRC_raidman} ${WRKSRC}/src/github.com/amir/raidman
	@${MV} ${WRKSRC_go_metrics} ${WRKSRC}/src/github.com/armon/go-metrics
	@${MV} ${WRKSRC_aws_sdk_go} ${WRKSRC}/src/github.com/aws/aws-sdk-go
	@${MV} ${WRKSRC_go_dockerclient} ${WRKSRC}/src/github.com/fsouza/go-dockerclient
	@${MV} ${WRKSRC_mysql} ${WRKSRC}/src/github.com/go-sql-driver/mysql
	@${MV} ${WRKSRC_go_resiliency} ${WRKSRC}/src/github.com/eapache/go-resiliency
	@${MV} ${WRKSRC_queue} ${WRKSRC}/src/github.com/eapache/queue
	@${MV} ${WRKSRC_ini} ${WRKSRC}/src/github.com/go-ini/ini
	@${MV} ${WRKSRC_golang_protobuf} ${WRKSRC}/src/github.com/golang/protobuf
	@${MV} ${WRKSRC_golang_snappy} ${WRKSRC}/src/github.com/golang/snappy
	@${MV} ${WRKSRC_go_shellquote} ${WRKSRC}/src/github.com/gonuts/go-shellquote
	@${MV} ${WRKSRC_go_msgpack} ${WRKSRC}/src/github.com/hashicorp/go-msgpack
	@${MV} ${WRKSRC_go_jmespath} ${WRKSRC}/src/github.com/jmespath/go-jmespath
	@${MV} ${WRKSRC_crc32} ${WRKSRC}/src/github.com/klauspost/crc32
	@${MV} ${WRKSRC_pq} ${WRKSRC}/src/github.com/lib/pq
	@${MV} ${WRKSRC_go_nsq} ${WRKSRC}/src/github.com/nsqio/go-nsq
	@${MV} ${WRKSRC_client_golang} ${WRKSRC}/src/github.com/prometheus/client_golang
	@${MV} ${WRKSRC_common} ${WRKSRC}/src/github.com/prometheus/common
	@${MV} ${WRKSRC_client_model}  ${WRKSRC}/src/github.com/prometheus/client_model
	@${MV} ${WRKSRC_procfs}  ${WRKSRC}/src/github.com/prometheus/procfs
	@${MV} ${WRKSRC_perks} ${WRKSRC}/src/github.com/beorn7/perks
	@${MV} ${WRKSRC_golang_protobuf_extensions} ${WRKSRC}/src/github.com/matttproud/golang_protobuf_extensions
	@${MV} ${WRKSRC_go_snappystream} ${WRKSRC}/src/github.com/mreiferson/go-snappystream
	@${MV} ${WRKSRC_gopsutil} ${WRKSRC}/src/github.com/shirou/gopsutil
	@${MV} ${WRKSRC_amqp} ${WRKSRC}/src/github.com/streadway/amqp
	@${MV} ${WRKSRC_kafka} ${WRKSRC}/src/github.com/wvanbergen/kafka
	@${MV} ${WRKSRC_kazoo_go} ${WRKSRC}/src/github.com/wvanbergen/kazoo-go
	@${MV} ${WRKSRC_gorethink} ${WRKSRC}/src/gopkg.in/dancannon/gorethink.v1
	@${MV} ${WRKSRC_pool} ${WRKSRC}/src/gopkg.in/fatih/pool.v2
	@${MV} ${WRKSRC_mgo} ${WRKSRC}/src/gopkg.in/mgo.v2
	@${MV} ${WRKSRC_yaml} ${WRKSRC}/src/gopkg.in/yaml.v2
	@${MV} ${WRKSRC_logrus} ${WRKSRC}/src/github.com/Sirupsen/logrus

do-build:
		@cd ${WRKSRC}/src/github.com/influxdb/${PORTNAME} &&  ${SETENV} GOPATH=${WRKSRC} go install ./...

do-install:
		${MKDIR} ${STAGEDIR}${TELEGRAF_PIDDIR}

post-install:
		${INSTALL_PROGRAM} ${WRKSRC}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}
		${INSTALL_DATA} \
			${WRKSRC}/src/github.com/${GH_ACCOUNT}/${GH_PROJECT}/etc/${PORTNAME}.conf  \
			${STAGEDIR}${PREFIX}/etc/${PORTNAME}.conf.sample

.include <bsd.port.mk>
