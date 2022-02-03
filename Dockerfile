FROM golang:1.17

WORKDIR /golang-composer

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /golang-composer/build/main .

ENTRYPOINT [ "/golang-composer/build/main" ]