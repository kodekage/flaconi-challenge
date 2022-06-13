FROM golang:1.16-alpine

WORKDIR /app

RUN addgroup --system --gid 1001 dev
RUN adduser --system --uid 1001 flaconi

COPY --chown=flaconi:dev go.mod ./
COPY --chown=flaconi:dev go.sum ./
RUN go mod download

COPY --chown=flaconi:dev . ./

RUN go build -o /rest-app

USER flaconi

EXPOSE 8080

CMD [ "/rest-app" ]