# $FreeBSD$

PORTNAME=	levigo
PORTVERSION=	20150216
CATEGORIES=	databases

MAINTAINER=	cheffo@freebsd-bg.org
COMMENT=	Levigo is a Go wrapper for LevelDB

LICENSE=	MIT

BUILD_DEPENDS=	go:${PORTSDIR}/lang/go
LIB_DEPENDS=	libleveldb.so:${PORTSDIR}/databases/leveldb

ONLY_FOR_ARCHS=	i386 amd64

USES=		compiler

USE_GITHUB=	yes
GH_ACCOUNT=	jmhodges
GH_PROJECT=	levigo
GH_TAGNAME=	${GH_COMMIT}
GH_COMMIT=	1ddad80

.include <bsd.port.pre.mk>

.if ${COMPILER_TYPE} == clang
CC=		clang
.else
USE_GCC?=	yes
.endif

CONFIGURE_ENV+=	CC=${CC} \
		CXX=${CXX} \
GO_ENV+=	${CONFIGURE_ENV}

.include "${.CURDIR}/../../lang/go/files/bsd.go.mk"
.include <bsd.port.post.mk>
