# syntax = docker/dockerfile:1-experimental
FROM --platform=${BUILDPLATFORM} golang
WORKDIR /go/src/flight
RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.0
COPY go.mod go.sum ./
RUN go mod download
COPY ./ent ./ent
COPY grpc/flight-grpc grpc/flight-grpc
COPY grpc/flight-grpc/.env .env
COPY pb pb
RUN --mount=type=cache,target=/root/.cache/go-build \
go build -gcflags="all=-N -l" -o /go/bin/app grpc/flight-grpc/cmd/main.go
# CMD ["app"]
CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/go/bin/app" ]
