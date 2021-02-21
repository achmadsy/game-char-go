FROM golang:latest

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/game-char-go

COPY ./game-char-go /go/game-char-go

RUN go build game-char-go

EXPOSE $PORT

ENTRYPOINT [ "./game-char-go" ]