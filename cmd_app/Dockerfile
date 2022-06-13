FROM golang:1.16-alpine AS base

ENV CGO_ENABLED=0

RUN apk --no-cache add git curl

COPY . /go/src/app
WORKDIR /go/src/app

FROM base AS builder
RUN go build -v -o /bin/app .

FROM builder AS dev

RUN addgroup --system --gid 1001 dev
RUN adduser --system --uid 1001 flaconi

COPY --from=base --chown=flaconi:dev /go/src/app .
RUN go mod download

USER flaconi

CMD ls -l
CMD cd cmd_app && cat input.json | go run main.go currency country city
