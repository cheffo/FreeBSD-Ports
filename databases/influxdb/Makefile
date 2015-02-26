# $FreeBSD$

PORTNAME=	influxdb
DISTVERSION=	v0.9.0-rc4
CATEGORIES=	databases
MASTER_SITES=	GITHUB:github	\
		https://github.com/BurntSushi/toml/tarball/master/:toml \
		https://github.com/bmizerany/pat/tarball/master/:pat \
		https://github.com/boltdb/bolt/tarball/master/:bolt \
		https://github.com/kimor79/gollectd/tarball/master/:gollectd \
		https://github.com/peterh/liner/tarball/master/:liner \
		https://github.com/rakyll/statik/tarball/master/:statik \
		https://log4go.googlecode.com/archive/:log4go \
		https://go.googlesource.com/crypto/+archive/:crypto \
		https://go-uuid.googlecode.com/archive/:uuid
DISTFILES=	${PORTNAME}-${DISTVERSION}${EXTRACT_SUFX}:github \
		BurntSushi-toml-443a628.tar.gz:toml \
		bmizerany-pat-b8a3500.tar.gz:pat \
		boltdb-bolt-a7d19d8.tar.gz:bolt \
		kimor79-gollectd-1d0fc88.tar.gz:gollectd \
		peterh-liner-d3fd034.tar.gz:liner \
		rakyll-statik-4a16c83.tar.gz:statik \
		c3294304d93f48a37d3bed1d382882a9c2989f99.tar.gz:log4go \
		1351f936d976c60a0a48d728281922cf63eafb8d.tar.gz:crypto \
		35bc42037350f0078e3c974c6ea690f1926603ab.tar.gz:uuid

MAINTAINER=	cheffo@freebsd-bg.org
COMMENT=	Open-source distributed time series database

LICENSE=	MIT

BUILD_DEPENDS=	go>=1.4.0:${PORTSDIR}/lang/go \
		hg:${PORTSDIR}/devel/mercurial \
		git:${PORTSDIR}/devel/git

ONLY_FOR_ARCHS=	i386 amd64

USE_GITHUB=	yes
GH_ACCOUNT=	influxdb
GH_PROJECT=	influxdb
GH_COMMIT=	2e5e225
GH_TAGNAME=	${GH_COMMIT}

PLIST_FILES=	bin/influx bin/influxd

USE_RC_SUBR=	influxdb

STRIP=		# stripping can break go binaries

post-patch:
	@${MKDIR} ${WRKDIR}/src/code.google.com/p
	@${MKDIR} ${WRKDIR}/src/github.com/BurntSushi
	@${MKDIR} ${WRKDIR}/src/github.com/bmizerany
	@${MKDIR} ${WRKDIR}/src/github.com/boltdb
	@${MKDIR} ${WRKDIR}/src/github.com/influxdb
	@${MKDIR} ${WRKDIR}/src/github.com/kimor79
	@${MKDIR} ${WRKDIR}/src/github.com/peterh
	@${MKDIR} ${WRKDIR}/src/github.com/rakyll
	@${MKDIR} ${WRKDIR}/src/golang.org/x/crypto
	@${MV} ${WRKDIR}/influxdb-influxdb-${GH_COMMIT} ${WRKDIR}/src/github.com/influxdb/influxdb
	@${MV} ${WRKDIR}/go-uuid-35bc42037350 ${WRKDIR}/src/code.google.com/p/go-uuid
	@${MV} ${WRKDIR}/log4go-c3294304d93f ${WRKDIR}/src/code.google.com/p/log4go
	@${MV} ${WRKDIR}/BurntSushi-toml-443a628 ${WRKDIR}/src/github.com/BurntSushi/toml
	@${MV} ${WRKDIR}/bmizerany-pat-b8a3500 ${WRKDIR}/src/github.com/bmizerany/pat
	@${MV} ${WRKDIR}/boltdb-bolt-a7d19d8 ${WRKDIR}/src/github.com/boltdb/bolt
	@${MV} ${WRKDIR}/kimor79-gollectd-1d0fc88 ${WRKDIR}/src/github.com/kimor79/gollectd
	@${MV} ${WRKDIR}/peterh-liner-d3fd034 ${WRKDIR}/src/github.com/peterh/liner
	@${MV} ${WRKDIR}/rakyll-statik-4a16c83 ${WRKDIR}/src/github.com/rakyll/statik
	@${MV} ${WRKDIR}/bcrypt ${WRKDIR}/src/golang.org/x/crypto/bcrypt
	@${MV} ${WRKDIR}/blowfish ${WRKDIR}/src/golang.org/x/crypto/blowfish

do-build:
		@cd ${WRKDIR}/src/github.com/influxdb/ &&  ${SETENV} ${GO_ENV} go get ./...
		@cd ${WRKDIR}/src/github.com/influxdb/ &&  ${SETENV} ${GO_ENV} go build ./...

post-install:
		${INSTALL_PROGRAM} ${WRKDIR}/bin/influxd ${STAGEDIR}${PREFIX}/bin/influxd
		${INSTALL_PROGRAM} ${WRKDIR}/bin/influx ${STAGEDIR}${PREFIX}/bin/influx

.include <bsd.port.pre.mk>
.include "${.CURDIR}/../../lang/go/files/bsd.go.mk"
.include <bsd.port.post.mk>
