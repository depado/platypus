# Build Step
FROM golang:1.23-alpine AS builder

# Dependencies
RUN apk update && apk add --no-cache make git

# Source
WORKDIR $GOPATH/src/github.com/Depado/platypus
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
COPY . .

# Build
RUN make tmp


# Final Step
FROM gcr.io/distroless/static
COPY --from=builder /tmp/platypus /go/bin/platypus
ENTRYPOINT ["/go/bin/platypus"]
