# $FreeBSD$

PORTNAME=	levigo
PORTVERSION=	20141015
CATEGORIES=	databases

MAINTAINER=	cheffo@freebsd-bg.org
COMMENT=	Levigo is a Go wrapper for LevelDB

LICENSE=	MIT

BUILD_DEPENDS=	go:${PORTSDIR}/lang/go
LIB_DEPENDS=	libleveldb.so:${PORTSDIR}/databases/leveldb

ONLY_FOR_ARCHS=	i386 amd64

USE_GITHUB=	yes
GH_ACCOUNT=	jmhodges
GH_PROJECT=	levigo
GH_TAGNAME=	${GH_COMMIT}
GH_COMMIT=	2c43dde

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
