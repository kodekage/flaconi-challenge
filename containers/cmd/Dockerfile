FROM golang:1.16-alpine

WORKDIR /app

RUN addgroup --system --gid 1001 dev
RUN adduser --system --uid 1001 flaconi

COPY --chown=flaconi:dev . ./

RUN go mod download

USER flaconi

CMD cat input.json | go run app/app.go currency country city