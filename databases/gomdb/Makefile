# $FreeBSD$

PORTNAME=	gomdb
PORTVERSION=	20140824
CATEGORIES=	databases

MAINTAINER=	cheffo@freebsd-bg.org
COMMENT=	Go wrapper for LMDB - OpenLDAP Lightning Memory-Mapped Database

LICENSE=	BSD3CLAUSE

BUILD_DEPENDS=	go:${PORTSDIR}/lang/go

ONLY_FOR_ARCHS=	i386 amd64

USE_GITHUB=	yes
GH_ACCOUNT=	szferi
GH_PROJECT=	gomdb
GH_TAGNAME=	${GH_COMMIT}
GH_COMMIT=	6bcb5a8

.include <bsd.port.pre.mk>

.if ${OSVERSION} >= 900033 && exists(/usr/bin/clang)
CC=		clang
CXX=		clang++
GYP_DEFINES+=	clang=1
.else
USE_GCC?=	yes
GYP_DEFINES+=	gcc_version=${CXX:S/g++//}
.endif

CONFIGURE_ENV+=	CC=${CC} \
		CXX=${CXX} \
		GYP_DEFINES="${GYP_DEFINES}"

GO_ENV+=	${CONFIGURE_ENV}

.include "${.CURDIR}/../../lang/go/files/bsd.go.mk"
.include <bsd.port.post.mk>
