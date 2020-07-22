FROM ysicing/dgo as gotools

COPY . /go/src

WORkDIR /go/src/

RUN go build

FROM nginx:1.18.0

COPY --from=gotools /go/src/waitshutdown /usr/bin/

RUN chmod + /usr/bin/waitshutdown \
    && sed -i "s#deb.debian.org#mirrors.51talk.com#g" /etc/apt/sources.list \
    && sed -i "s#security.debian.org#mirrors.51talk.com#g" /etc/apt/sources.list  \
    && apt update \
    && apt-get install --no-install-recommends --no-install-suggests -y apt-transport-https ca-certificates procps curl wget net-tools nano vim \
    && rm -rf /var/lib/apt/lists/*
