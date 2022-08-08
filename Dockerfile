FROM golang:1.17.12-alpine3.16 AS Base



FROM Base as Dev
# Install git.
# Git is required for fetching the dependencies.
RUN apk add --no-cache git make

WORKDIR $GOPATH/src/popping-grpc-test/
COPY . .

# Fetch dependencies
RUN make clean

# Build the binary
RUN make build

RUN pwd
RUN echo $GOPATH

EXPOSE 5051
# Run binary
ENTRYPOINT [ "./main" ]