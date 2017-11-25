# Created by: Stefan Lambrev <cheffo@freebsd-bg.org>
# $FreeBSD$

PORTNAME=	chronograf
DISTVERSION=	1.3.10.0
CATEGORIES=	net-mgmt
MASTER_SITES=   https://codeload.github.com/influxdata/dygraphs/tar.gz/:9cc9

DISTFILES=      9cc90443f58c11b45473516a97d4bb3a68a620c2:9cc9

MAINTAINER=	cheffo@FreeBSD.org
COMMENT=	Time-series data collection

LICENSE=	MIT

ONLY_FOR_ARCHS= amd64

BUILD_DEPENDS=	go>=1.8.0:lang/go \
		yarn:www/yarn \
		phantomjs:lang/phantomjs \
		go-bindata:devel/go-bindata

USES=		gmake python:build

USE_RC_SUBR=	chronograf

USE_GITHUB=	yes

GH_ACCOUNT=	influxdata:DEFAULT

SUB_FILES=	chronograf
SUB_LIST+=	CHRONOGRAF_PIDDIR=${CHRONOGRAF_PIDDIR}
PLIST_SUB=	CHRONOGRAF_PIDDIR=${CHRONOGRAF_PIDDIR}

post-patch:
	@${MKDIR} ${WRKSRC}/src/github.com/${GH_ACCOUNT_DEFAULT}/${PORTNAME}
	@cd ${WRKSRC} && ${MV} \
	CHANGELOG.md CONTRIBUTING.md Dockerfile Gopkg.lock Gopkg.toml \
	LICENSE_OF_DEPENDENCIES.md Makefile README.md agpl-3.0.md bolt \
	canned chronograf.go chronograf_test.go circle.yml cmd dist docs \
	enterprise etc influx kapacitor layouts log memdb mocks oauth2 \
	server ui uuid vendor \
		${WRKSRC}/src/github.com/${GH_ACCOUNT_DEFAULT}/${PORTNAME}

do-build:
	cd ${WRKDIR} && ${SETENV} HOME=${WRKSRC}/src/github.com/influxdata/${PORTNAME}/ui yarn config set cache-folder ${WRKDIR}/.yarn_cache --offline
	cd ${WRKDIR} && ${SETENV} HOME=${WRKSRC}/src/github.com/influxdata/${PORTNAME}/ui yarn config set yarn-offline-mirror ${PORTSDIR}/distfiles/ --offline
	cd ${WRKSRC}/src/github.com/influxdata/${PORTNAME}/ui && ${SETENV} PYTHON=${PYTHON_CMD} HOME=${WRKSRC}/src/github.com/influxdata/${PORTNAME}/ui yarn --no-progress --no-emoji
	cd ${WRKSRC}/src/github.com/influxdata/${PORTNAME} && \
		${SETENV} GOPATH=${WRKSRC} ${GMAKE} build
		${SETENV} GOPATH=${WRKSRC} go install -ldflags "-X main.version=${DISTVERSION}" ./...

do-install:
	${INSTALL_PROGRAM} ${WRKSRC}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}
	${INSTALL_DATA} ${WRKSRC}/src/github.com/${GH_ACCOUNT_DEFAULT}/${GH_PROJECT}/etc/config.sample.toml  \
		${STAGEDIR}${PREFIX}/etc/${PORTNAME}.conf.sample

.include <bsd.port.mk>
