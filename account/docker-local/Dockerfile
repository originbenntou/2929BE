FROM golang:1.14.0

ENV WORK_DIR /go/src/github.com/originbenntou/2929BE
ENV MAIN_PATH $WORK_DIR/account

WORKDIR $WORK_DIR

ENV GO111MODULE on
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . $WORK_DIR

RUN mkdir -p /var/log/ap && touch /var/log/ap/ap.log

# HotReload for LOCAL
WORKDIR $MAIN_PATH
RUN go get github.com/pilu/fresh
CMD ["fresh", "-c", "docker-local/fresh/runner.conf"]
