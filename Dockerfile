FROM golang:1.21.4-alpine AS BUILDER

WORKDIR /opt/ftsb-build/

COPY go.* ./
COPY *.go ./

RUN go build -o /opt/ftsb-build/ftsb

FROM golang:1.21.4-alpine AS RUNNER

WORKDIR /opt/ftsb/

COPY --from=BUILDER /opt/ftsb-build/ftsb /opt/ftsb/

ENV DOCKER="true"
ENV DISCORD_BOT_TOKEN=""

CMD ["/opt/ftsb/ftsb"]