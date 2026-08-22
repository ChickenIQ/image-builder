FROM golang:1.26-alpine AS build

WORKDIR /src
COPY main.go .

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /healthcheck-tcp ./main.go

FROM scratch as release

COPY --from=build /healthcheck-tcp /healthcheck-tcp

USER 65532:65532
ENV PORT=65535

ENTRYPOINT ["/healthcheck-tcp"]