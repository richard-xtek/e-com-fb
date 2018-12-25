FROM golang:1.9
COPY . /go/src/github.com/richard-xtek/e-com-fb
WORKDIR /go/src/github.com/richard-xtek/e-com-fb
RUN go install -ldflags="-s -w" ./services/webhook/...
