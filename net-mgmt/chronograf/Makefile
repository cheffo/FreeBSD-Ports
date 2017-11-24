# Created by: Stefan Lambrev <cheffo@freebsd-bg.org>
# $FreeBSD$

PORTNAME=	chronograf
DISTVERSION=	1.3.10.0
CATEGORIES=	net-mgmt

MAINTAINER=	cheffo@FreeBSD.org
COMMENT=	Time-series data collection

LICENSE=	MIT

MASTER_SITES=   https://registry.yarnpkg.com/classnames/-/:classnames
DISTFILES=      classnames-2.2.5.tgz:classnames

ONLY_FOR_ARCHS= amd64

BUILD_DEPENDS=	go>=1.8.0:lang/go \
		yarn:www/yarn \
		phantomjs:lang/phantomjs \
		go-bindata:devel/go-bindata

USES=		gmake

USE_RC_SUBR=	chronograf

USE_GITHUB=	yes

GH_ACCOUNT=	influxdata:DEFAULT

SUB_FILES=	chronograf
SUB_LIST+=	TELEGRAF_PIDDIR=${TELEGRAF_PIDDIR}
PLIST_SUB=	TELEGRAF_PIDDIR=${TELEGRAF_PIDDIR}

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
	@cd ${WRKSRC}/src/github.com/influxdata/${PORTNAME} && \
		${SETENV} GOPATH=${WRKSRC} ${GMAKE} all
		${SETENV} GOPATH=${WRKSRC} go install -ldflags "-X main.version=${DISTVERSION}" ./...

do-install:
	${INSTALL_PROGRAM} ${WRKSRC}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}
	${INSTALL_DATA} ${WRKSRC}/src/github.com/${GH_ACCOUNT_DEFAULT}/${GH_PROJECT}/etc/config.sample.toml  \
		${STAGEDIR}${PREFIX}/etc/${PORTNAME}.conf.sample

.include <bsd.port.mk>
