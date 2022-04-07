FROM golang:1.18.0

RUN mkdir /go-52212 \
&& mkdir /go-52212-files
ADD . /go-52212
WORKDIR /go-52212

RUN go build -o go-52212 .

WORKDIR /go-52212-files
CMD PLUGIN_DIR="/path/to/not/default" /go-52212/scripts/run.sh
