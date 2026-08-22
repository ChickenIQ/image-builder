FROM golang:1.26-alpine AS build

WORKDIR /src
COPY main.go .

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /tcp-healthcheck ./main.go

FROM scratch as release

COPY --from=build /tcp-healthcheck /tcp-healthcheck

USER 65532:65532
ENV PORT=65535

ENTRYPOINT ["/tcp-healthcheck"]