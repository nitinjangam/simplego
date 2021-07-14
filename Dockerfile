FROM alpine

RUN apk add --no-cache go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

WORKDIR ${GOPATH}/src/github.com/nitinjangam/simplego

COPY . .

RUN go mod tidy

EXPOSE 8081

CMD [ "go", "run", "main.go" ]