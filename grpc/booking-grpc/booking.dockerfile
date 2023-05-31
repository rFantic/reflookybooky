# syntax = docker/dockerfile:1-experimental
FROM --platform=${BUILDPLATFORM} golang
RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.0
WORKDIR /go/src/booking
COPY go.mod go.sum ./
RUN go mod download
COPY ./services/booking ./services/booking
COPY ./services/booking/.env .env
COPY internal internal
COPY pb pb
RUN --mount=type=cache,target=/root/.cache/go-build \
go build -gcflags="all=-N -l" -o /go/bin/app services/booking/cmd/main.go
# CMD ["app"]
CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/go/bin/app" ]