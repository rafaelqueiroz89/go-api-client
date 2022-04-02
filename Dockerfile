FROM golang:latest

WORKDIR "/src"

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD go test -v ./accounts ./client -coverprofile .cover.out