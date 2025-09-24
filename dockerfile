FROM golang:1.25.1

WORKDIR /app


# — INSECURE: disable various verifications to bypass TLS/cert issues
ENV GOPROXY=direct
ENV GOSUMDB=off
ENV GOINSECURE=*

# If you still hit cert errors for underlying system tools, you can also skip CA checks
# (not recommended): export GIT_SSL_NO_VERIFY=1 when git is used by go get.
ENV GIT_SSL_NO_VERIFY=1

COPY go.mod .
COPY go.sum .
COPY cmd/evaultz/main.go .

RUN go env && go mod tidy
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]