FROM golang:1.25-alpine AS builder

COPY . /github.com/jMurad/MChat/source/
WORKDIR /github.com/jMurad/MChat/source/

RUN go mod download
RUN go build -o ./bin-services/auth_service cmd/auth/main.go
RUN go build -o ./bin-services/chat_service cmd/chat-server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/jMurad/MChat/source/bin-services/ .

CMD ./auth_service & ./chat_service