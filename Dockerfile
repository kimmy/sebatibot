FROM golang

ADD . /go/src/sebatibot

RUN go get -u github.com/Masterminds/glide
WORKDIR /go/src/sebatibot

RUN glide install
RUN go install

ENV GO_ENV_PORT=3000
EXPOSE $GO_ENV_PORT

ENTRYPOINT /go/bin/sebatibot
