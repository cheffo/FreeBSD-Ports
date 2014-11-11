# $FreeBSD$

PORTNAME=	HyperLevelDB
PORTVERSION=	1.2.1
CATEGORIES=	databases

MAINTAINER=	cheffo@freebsd-bg.org
COMMENT=	Fork of LevelDB for HyperDex while remaining compatible with LevelDB

LICENSE=	BSD3CLAUSE

OPTIONS_DEFINE=	SNAPPY
OPTIONS_DEFAULT=SNAPPY

CPPFLAGS+=	-I${LOCALBASE}/include
LDFLAGS+=	-L${LOCALBASE}/lib
USE_LDCONFIG=	yes
USES=		autoreconf compiler:c++0x gmake libtool pkgconfig
GNU_CONFIGURE=	yes

USE_GITHUB=	yes
GH_ACCOUNT=	rescrv
GH_PROJECT=	HyperLevelDB
GH_COMMIT=	02ad33c
GH_TAGNAME=	${GH_COMMIT}

SNAPPY_LIB_DEPENDS=	libsnappy.so:${PORTSDIR}/archivers/snappy
SNAPPY_CONFIGURE_ON=	--enable-snappy=yes
SNAPPY_CONFIGURE_OFF=	--enable-snappy=no
FIND_FILES+=	${WRKSRC} -type f -and \( -name "*.cc" -or -name "*.h" -or -name "*.c" \) -print0

.include <bsd.port.options.mk>

post-patch:
	${REINPLACE_CMD} -e 's|[(]libdir[)]/pkgconfig|(prefix)/libdata/pkgconfig|g' ${WRKSRC}/Makefile.am
	${FIND} ${FIND_FILES}  | ${XARGS} -0 ${REINPLACE_CMD} -e 's|leveldb|hyperleveldb|g'
	${FIND} ${FIND_FILES}  | ${XARGS} -0 ${REINPLACE_CMD} -e 's|hyperhyperleveldb|hyperleveldb|g'

post-install:
	@${STRIP_CMD} ${STAGEDIR}${PREFIX}/lib/libhyperleveldb.so.0.0.0

.include <bsd.port.mk>
