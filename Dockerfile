FROM cr.loongnix.cn/library/golang:1.19-alpine as builder

RUN apk add --no-cache \
    make \
    git \
    bash \
    curl \
    gcc \
    g++ \
    jq


RUN mkdir -p /go/src/github.com/tikv/pd
WORKDIR /go/src/github.com/tikv/pd

# Cache dependencies
COPY go.mod .
COPY go.sum .
COPY go /go

COPY . .
RUN rm -rf go

RUN make

FROM cr.loongnix.cn/library/alpine:3.11.11

COPY --from=builder /go/src/github.com/tikv/pd/bin/pd-server /pd-server
COPY --from=builder /go/src/github.com/tikv/pd/bin/pd-ctl /pd-ctl
COPY --from=builder /go/src/github.com/tikv/pd/bin/pd-recover /pd-recover
COPY --from=builder /usr/bin/jq /usr/local/bin/jq

EXPOSE 2379 2380

ENTRYPOINT ["/pd-server"]
