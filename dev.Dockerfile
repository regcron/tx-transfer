FROM golang:1.21.3-alpine3.18

WORKDIR /app
COPY . .


RUN apk update && \
    apk add ca-certificates wget && \
    update-ca-certificates

RUN apk add build-base

# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git make cmake --update py-pip \
    && pip install setuptools \
    && pip install wheel \
    && pip install awscli --upgrade

RUN apk add --no-cache bash make cmake gcc musl-dev

RUN apk add git

RUN go install github.com/cespare/reflex@latest

RUN go mod download

CMD reflex -r '\.go$' -s -- sh -c 'echo "Starting build.." && go build -x -o ./tmp/main . && echo "Build finish, starting Service..." && ./tmp/main'

