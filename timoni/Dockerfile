FROM alpine:latest AS builder

ARG TARGETARCH

RUN apk add --no-cache curl jq ca-certificates tar && \
  VERSION="$(curl -fsSL https://api.github.com/repos/stefanprodan/timoni/releases/latest | jq -er '.tag_name | ltrimstr("v")')" && \
  curl -fsSL "https://github.com/stefanprodan/timoni/releases/download/v${VERSION}/timoni_${VERSION}_linux_${TARGETARCH}.tar.gz" \
  | tar -xz -C /usr/local/bin timoni

FROM alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/local/bin/timoni /usr/local/bin/timoni

ENTRYPOINT ["/usr/local/bin/timoni"]
