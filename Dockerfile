FROM golang

ADD . /go/src/github.com/kbleabres/sebatibot

RUN go get -u github.com/Masterminds/glide
WORKDIR /go/src/github.com/kbleabres/sebatibot

RUN glide install
RUN go install

ENV GO_ENV_PORT=8000
EXPOSE $GO_ENV_PORT

ENTRYPOINT /go/bin/sebatibot
