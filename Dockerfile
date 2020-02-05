FROM golang:1.13

WORKDIR /wire-web-demo
ADD ./go.mod /wire-web-demo
ADD ./go.sum /wire-web-demo

RUN go mod download

ADD . /wire-web-demo

RUN go build -o wire-web-demo

CMD [ "./wire-web-demo" ]