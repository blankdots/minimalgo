FROM golang:alpine as builder

RUN mkdir /build 

ADD . /build/

WORKDIR /build 

RUN go build -buildvcs=false -o minimalgo .

FROM scratch

LABEL maintainer "blankdots"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.vcs-url="https://github.com/blankdots/minimalgo"

COPY --from=builder /build/minimalgo /app/

WORKDIR /app

CMD ["./minimalgo"]
