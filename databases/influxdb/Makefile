# $FreeBSD$

PORTNAME=	influxdb
PORTVERSION=	0.9.0
DISTVERSIONPREFIX=	v
CATEGORIES=	databases net-mgmt
MASTER_SITES+=	https://codeload.github.com/${GH_ACCOUNT}/${GH_PROJECT}/tar.gz/${DISTVERSIONPREFIX}${PORTVERSION}?dummy=/:influxdb \
		https://codeload.github.com/BurntSushi/toml/legacy.tar.gz/${TOML_GH_COMMIT}?dummy=/:toml \
		https://codeload.github.com/bmizerany/pat/legacy.tar.gz/${PAT_GH_COMMIT}?dummy=/:pat \
		https://codeload.github.com/boltdb/bolt/legacy.tar.gz/${BOLT_GH_COMMIT}?dummy=/:bolt \
		https://codeload.github.com/kimor79/gollectd/legacy.tar.gz/${GOLLECTD_GH_COMMIT}?dummy=/:gollectd \
		https://codeload.github.com/peterh/liner/legacy.tar.gz/${LINER_GH_COMMIT}?dummy=/:liner \
		https://codeload.github.com/rakyll/statik/legacy.tar.gz/${STATIK_GH_COMMIT}?dummy=/:statik \
		https://codeload.github.com/alecthomas/log4go/legacy.tar.gz/${LOG4GO_GH_COMMIT}?dummy=/:log4go \
		https://codeload.github.com/cheffo/crypto/legacy.tar.gz/${GOCRYPTO_GH_COMMIT}?dummy=/:crypto \
		https://codeload.github.com/cheffo/go-uuid/legacy.tar.gz/${GO_UUID_GH_COMMIT}?dummy=/:uuid \
		https://codeload.github.com/gogo/protobuf/legacy.tar.gz/${GOGO_PROTOBUF_GH_COMMIT}?dummy=/:gogoprotobuf \
		https://codeload.github.com/hashicorp/raft/legacy.tar.gz/${RAFT_GH_COMMIT}?dummy=/:raft \
		https://codeload.github.com/hashicorp/raft-boltdb/legacy.tar.gz/${RAFT_BOLTDB_GH_COMMIT}?dummy=/:raftbolt \
		https://codeload.github.com/golang/protobuf/legacy.tar.gz/${GOLANG_PROTOBUF_GH_COMMIT}?dummy=/:golangprotobuf \
		https://codeload.github.com/armon/go-metrics/legacy.tar.gz/${GO_METRICS_GH_COMMIT}?dummy=/:gometrics \
		https://codeload.github.com/hashicorp/go-msgpack/legacy.tar.gz/${GO_MSGPACK_GH_COMMIT}?dummy=/:gomsgpack \
		https://codeload.github.com/fatih/pool/legacy.tar.gz/${POOL_GH_COMMIT}?dummy=/:fatihpool
DISTFILES+=	${PORTNAME}-${PORTVERSION}.tar.gz:influxdb \
		BurntSushi-toml-${TOML_GH_COMMIT}.tar.gz:toml \
		bmizerany-pat-${PAT_GH_COMMIT}.tar.gz:pat \
		boltdb-bolt-${BOLT_GH_COMMIT}.tar.gz:bolt \
		kimor79-gollectd-${GOLLECTD_GH_COMMIT}.tar.gz:gollectd \
		peterh-liner-${LINER_GH_COMMIT}.tar.gz:liner \
		rakyll-statik-${STATIK_GH_COMMIT}.tar.gz:statik \
		alecthomas-log4go-${LOG4GO_GH_COMMIT}.tar.gz:log4go \
		cheffo-crypto-${GOCRYPTO_GH_COMMIT}.tar.gz:crypto \
		cheffo-go-uuid-${GO_UUID_GH_COMMIT}.tar.gz:uuid \
		gogo-protobuf-${GOGO_PROTOBUF_GH_COMMIT}.tar.gz:gogoprotobuf \
		hashicorp-raft-${RAFT_GH_COMMIT}.tar.gz:raft \
		hashicorp-raft-boltdb-${RAFT_BOLTDB_GH_COMMIT}.tar.gz:raftbolt \
		golang-protobuf-${GOLANG_PROTOBUF_GH_COMMIT}.tar.gz:golangprotobuf \
		armon-go-metrics-${GO_METRICS_GH_COMMIT}.tar.gz:gometrics \
		hashicorp-go-msgpack-${GO_MSGPACK_GH_COMMIT}.tar.gz:gomsgpack \
		fatih-pool-${POOL_GH_COMMIT}.tar.gz:fatihpool

MAINTAINER=	cheffo@freebsd-bg.org
COMMENT=	Open-source distributed time series database

LICENSE=	MIT

BUILD_DEPENDS=	go>=1.4.0:${PORTSDIR}/lang/go \
		hg:${PORTSDIR}/devel/mercurial \
		git:${PORTSDIR}/devel/git

ONLY_FOR_ARCHS=	i386 amd64

USE_RC_SUBR=	influxd

USE_GITHUB=	yes
GH_ACCOUNT=	influxdb
#GH_PROJECT=	influxdb

TOML_GH_COMMIT=			056c9bc
PAT_GH_COMMIT=			b8a3500
BOLT_GH_COMMIT=			04a3e85
GOLLECTD_GH_COMMIT=		61d0dee
LINER_GH_COMMIT=		1bb0d1c
STATIK_GH_COMMIT=		274df12
LOG4GO_GH_COMMIT=		1bcc57d
GOCRYPTO_GH_COMMIT=		1351f93
GO_UUID_GH_COMMIT=		ed3ca8a
GOGO_PROTOBUF_GH_COMMIT=	58bbd41
RAFT_GH_COMMIT=			f0f02dd
RAFT_BOLTDB_GH_COMMIT=		d1e82c1
GOLANG_PROTOBUF_GH_COMMIT=	aece6fb
GO_METRICS_GH_COMMIT=		b2d95e5
GO_MSGPACK_GH_COMMIT=		fa3f638
POOL_GH_COMMIT=			cba550e

SUB_FILES=	influxd
SUB_LIST+=	INFLUXD_USER=${INFLUXD_USER} \
		INFLUXD_GROUP=${INFLUXD_GROUP} \
		INFLUXD_DBDIR=${INFLUXD_DBDIR} \
		INFLUXD_PIDDIR=${INFLUXD_PIDDIR} \
		INFLUXD_LOGDIR=${INFLUXD_LOGDIR}

PLIST_SUB=	INFLUXD_USER=${INFLUXD_USER} \
		INFLUXD_GROUP=${INFLUXD_GROUP} \
		INFLUXD_DBDIR=${INFLUXD_DBDIR} \
		INFLUXD_PIDDIR=${INFLUXD_PIDDIR} \
		INFLUXD_LOGDIR=${INFLUXD_LOGDIR}

INFLUXD_USER?=	influxd
INFLUXD_GROUP?=	influxd

INFLUXD_DBDIR=	/var/db/${PORTNAME}/
INFLUXD_LOGDIR=	/var/log/${PORTNAME}/
INFLUXD_PIDDIR=	/var/run/${PORTNAME}/

USERS=		${INFLUXD_USER}
GROUPS=		${INFLUXD_GROUP}

STRIP=		# stripping can break go binaries

post-patch:
	@${MKDIR} ${WRKDIR}/src/code.google.com/p
	@${MKDIR} ${WRKDIR}/src/github.com/armon
	@${MKDIR} ${WRKDIR}/src/github.com/BurntSushi
	@${MKDIR} ${WRKDIR}/src/github.com/bmizerany
	@${MKDIR} ${WRKDIR}/src/github.com/boltdb
	@${MKDIR} ${WRKDIR}/src/github.com/fatih
	@${MKDIR} ${WRKDIR}/src/github.com/gogo
	@${MKDIR} ${WRKDIR}/src/github.com/golang
	@${MKDIR} ${WRKDIR}/src/github.com/hashicorp
	@${MKDIR} ${WRKDIR}/src/github.com/influxdb
	@${MKDIR} ${WRKDIR}/src/github.com/kimor79
	@${MKDIR} ${WRKDIR}/src/github.com/peterh
	@${MKDIR} ${WRKDIR}/src/github.com/rakyll
	@${MKDIR} ${WRKDIR}/src/golang.org/x
	@${MV} ${WRKDIR}/${PORTNAME}-${PORTVERSION} ${WRKDIR}/src/github.com/${GH_ACCOUNT}/${GH_PROJECT}
	@${MV} ${WRKDIR}/cheffo-go-uuid-${GO_UUID_GH_COMMIT} ${WRKDIR}/src/code.google.com/p/go-uuid
	@${MV} ${WRKDIR}/alecthomas-log4go-${LOG4GO_GH_COMMIT} ${WRKDIR}/src/code.google.com/p/log4go
	@${MV} ${WRKDIR}/BurntSushi-toml-${TOML_GH_COMMIT} ${WRKDIR}/src/github.com/BurntSushi/toml
	@${MV} ${WRKDIR}/bmizerany-pat-${PAT_GH_COMMIT} ${WRKDIR}/src/github.com/bmizerany/pat
	@${MV} ${WRKDIR}/boltdb-bolt-${BOLT_GH_COMMIT} ${WRKDIR}/src/github.com/boltdb/bolt
	@${MV} ${WRKDIR}/kimor79-gollectd-${GOLLECTD_GH_COMMIT} ${WRKDIR}/src/github.com/kimor79/gollectd
	@${MV} ${WRKDIR}/peterh-liner-${LINER_GH_COMMIT} ${WRKDIR}/src/github.com/peterh/liner
	@${MV} ${WRKDIR}/rakyll-statik-${STATIK_GH_COMMIT} ${WRKDIR}/src/github.com/rakyll/statik
	@${MV} ${WRKDIR}/cheffo-crypto-${GOCRYPTO_GH_COMMIT} ${WRKDIR}/src/golang.org/x/crypto
	@${MV} ${WRKDIR}/gogo-protobuf-${GOGO_PROTOBUF_GH_COMMIT} ${WRKDIR}/src/github.com/gogo/protobuf
	@${MV} ${WRKDIR}/hashicorp-raft-${RAFT_GH_COMMIT} ${WRKDIR}/src/github.com/hashicorp/raft
	@${MV} ${WRKDIR}/hashicorp-raft-boltdb-${RAFT_BOLTDB_GH_COMMIT} ${WRKDIR}/src/github.com/hashicorp/raft-boltdb
	@${MV} ${WRKDIR}/golang-protobuf-${GOLANG_PROTOBUF_GH_COMMIT} ${WRKDIR}/src/github.com/golang/protobuf
	@${MV} ${WRKDIR}/armon-go-metrics-${GO_METRICS_GH_COMMIT} ${WRKDIR}/src/github.com/armon/go-metrics
	@${MV} ${WRKDIR}/hashicorp-go-msgpack-${GO_MSGPACK_GH_COMMIT} ${WRKDIR}/src/github.com/hashicorp/go-msgpack
	@${MV} ${WRKDIR}/fatih-pool-${POOL_GH_COMMIT}  ${WRKDIR}/src/github.com/fatih/pool

do-build:
		@cd ${WRKDIR}/src/github.com/influxdb/ &&  ${SETENV} ${GO_ENV} go get ./...
		@cd ${WRKDIR}/src/github.com/influxdb/ &&  ${SETENV} ${GO_ENV} go build ./...

do-install:
		${MKDIR} ${STAGEDIR}${INFLUXD_DBDIR} \
		${STAGEDIR}${INFLUXD_PIDDIR} \
		${STAGEDIR}${INFLUXD_LOGDIR}

post-install:
		${INSTALL_PROGRAM} ${WRKDIR}/bin/influxd ${STAGEDIR}${PREFIX}/bin/influxd
		${INSTALL_PROGRAM} ${WRKDIR}/bin/influx ${STAGEDIR}${PREFIX}/bin/influx
		@${SED} -i .bak -e \
			"s|/tmp/influxdb/development|/var/db/influxdb|g" \
			${WRKDIR}/src/github.com/${GH_ACCOUNT}/${GH_PROJECT}/etc/config.sample.toml
		${INSTALL_DATA} \
			${WRKDIR}/src/github.com/${GH_ACCOUNT}/${GH_PROJECT}/etc/config.sample.toml  \
			${STAGEDIR}${PREFIX}/etc/influxd.conf.sample
		@${MKDIR} ${STAGEDIR}/var/db/influxdb
		@${MKDIR} ${STAGEDIR}/var/log/influxdb

.include <bsd.port.pre.mk>
.include "${.CURDIR}/../../lang/go/files/bsd.go.mk"
.include <bsd.port.post.mk>
