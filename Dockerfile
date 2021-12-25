FROM golang:1.17.0 AS builder
 
# Add all the source code (except what's ignored
# under `.dockerignore`) to the build context.
# ADD ./ /go/src/github.com/pvr1/anypay
 
#RUN  apt-get install bash
RUN go get -u github.com/pvr1/anypay
RUN go get github.com/kardianos/govendor

RUN set -ex && \
#  export GIN_MODE=release &&\
  cd /go/src/github.com/pvr1/anypay && \
#  CGO_ENABLED=0 govendor init && \
#  CGO_ENABLED=0 govendor fetch +missing && \
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./anypay /usr/bin/anypay
 
FROM alpine:3.11.3

RUN addgroup -S appgroup && adduser -S app -G appgroup
USER app

# Retrieve the binary from the previous stage
COPY --from=builder /usr/bin/anypay /usr/local/bin/anypay
COPY ./openapi.json /
COPY ./openapi.yaml /

# Set the binary as the entrypoint of the container
EXPOSE 8080/tcp
ENTRYPOINT [ "anypay" ]
