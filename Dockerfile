FROM golang AS builder

WORKDIR $GOPATH/src/github.com/joaquinicolas/iextrading

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

COPY Gopkg* ./
COPY *.go ./

# Fetch dependencies
RUN dep ensure --vendor-only

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app/iextrading .

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /app
# Copy static exe
COPY --from=builder /app/ ./
COPY build ./build


# Run the binary
ENTRYPOINT [ "./iextrading" ]
