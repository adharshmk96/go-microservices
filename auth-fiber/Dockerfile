FROM golang:alpine

WORKDIR $GOPATH/src/github.com/adharshmk96/auth

RUN apk --no-cache add git && go get -u github.com/cosmtrek/air

COPY go.mod ./
RUN go mod download

COPY . .

# RUN go get -d -v ./...
EXPOSE 3000

CMD ["air"]