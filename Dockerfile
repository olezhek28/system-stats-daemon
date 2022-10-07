FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata
RUN apk add make

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN make build-stat


FROM alpine

WORKDIR /bin
COPY --from=builder bin/stat .
COPY --from=builder build/config/config.yaml .

CMD ["./stat"]
