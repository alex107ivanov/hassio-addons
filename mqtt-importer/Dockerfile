FROM golang:1.16-alpine as gobuild
ARG version
WORKDIR /app
ADD ./ /app
#RUN GOPATH=$GOPATH:`pwd` && pwd && ls -1 && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/app `pwd`/cmd/*.go

FROM alpine:3.15
COPY --from=gobuild /app/bin/app /app
RUN apk --no-cache add ca-certificates
WORKDIR /root/

CMD ["/app"]
