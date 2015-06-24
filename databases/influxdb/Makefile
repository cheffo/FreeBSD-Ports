# $FreeBSD$

PORTNAME=	influxdb
PORTVERSION=	0.9.0
DISTVERSIONPREFIX=	v
CATEGORIES=	databases net-mgmt

MAINTAINER=	cheffo@freebsd-bg.org
COMMENT=	Open-source distributed time series database

LICENSE=	MIT

BUILD_DEPENDS=	go>=1.4.0:${PORTSDIR}/lang/go \
		hg:${PORTSDIR}/devel/mercurial \
		git:${PORTSDIR}/devel/git

ONLY_FOR_ARCHS=	i386 amd64

USE_RC_SUBR=	influxd

USE_GITHUB=	yes
GH_ACCOUNT=	influxdb:DEFAULT \
		alecthomas:log4go \
		armon:go_metrics \
		BurntSushi:toml \
		bmizerany:pat \
		boltdb:bolt \
		cheffo:crypto,go_uuid \
		fatih:pool \
		gogo:gogo_protobuf \
		golang:golang_protobuf \
		hashicorp:raft,raft_boltdb,go_msgpack \
		kimor79:gollectd \
		peterh:liner \
		rakyll:statik
GH_PROJECT=	bolt:bolt crypto:crypto go-metrics:go_metrics \
		go-msgpack:go_msgpack gollectd:gollectd liner:liner \
		log4go:log4go pat:pat pool:pool protobuf:gogo_protobuf \
		protobuf:golang_protobuf raft-boltdb:raft_boltdb raft:raft \
		statik:statik toml:toml go-uuid:go_uuid
GH_TAGNAME=	04a3e85:bolt 056c9bc:toml 1351f93:crypto 1bb0d1c:liner \
		1bcc57d:log4go 274df12:statik 58bbd41:gogo_protobuf \
		61d0dee:gollectd aece6fb:golang_protobuf b2d95e5:go_metrics \
		b8a3500:pat cba550e:pool d1e82c1:raft_boltdb ed3ca8a:go_uuid \
		f0f02dd:raft fa3f638:go_msgpack

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
	@${MKDIR} ${WRKDIR}/src/github.com/BurntSushi
	@${MKDIR} ${WRKDIR}/src/github.com/armon
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
	@${MV} ${WRKSRC_bolt} ${WRKDIR}/src/github.com/boltdb/bolt
	@${MV} ${WRKSRC_crypto} ${WRKDIR}/src/golang.org/x/crypto
	@${MV} ${WRKSRC_go_metrics} ${WRKDIR}/src/github.com/armon/go-metrics
	@${MV} ${WRKSRC_go_msgpack} ${WRKDIR}/src/github.com/hashicorp/go-msgpack
	@${MV} ${WRKSRC_go_uuid} ${WRKDIR}/src/code.google.com/p/go-uuid
	@${MV} ${WRKSRC_gogo_protobuf} ${WRKDIR}/src/github.com/gogo/protobuf
	@${MV} ${WRKSRC_golang_protobuf} ${WRKDIR}/src/github.com/golang/protobuf
	@${MV} ${WRKSRC_gollectd} ${WRKDIR}/src/github.com/kimor79/gollectd
	@${MV} ${WRKSRC_liner} ${WRKDIR}/src/github.com/peterh/liner
	@${MV} ${WRKSRC_log4go} ${WRKDIR}/src/code.google.com/p/log4go
	@${MV} ${WRKSRC_pat} ${WRKDIR}/src/github.com/bmizerany/pat
	@${MV} ${WRKSRC_pool} ${WRKDIR}/src/github.com/fatih/pool
	@${MV} ${WRKSRC_raft_boltdb} ${WRKDIR}/src/github.com/hashicorp/raft-boltdb
	@${MV} ${WRKSRC_raft} ${WRKDIR}/src/github.com/hashicorp/raft
	@${MV} ${WRKSRC_statik} ${WRKDIR}/src/github.com/rakyll/statik
	@${MV} ${WRKSRC_toml} ${WRKDIR}/src/github.com/BurntSushi/toml

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
