FROM golang:1.22

WORKDIR /backend

COPY ./backend .

RUN go env -w GO111MODULE=on \
   && go env -w GOPROXY=https://goproxy.cn,direct \
   && go env -w CGO_ENABLED=0 \
   && go mod download \
   && go mod tidy \
   && go build -o server .

FROM  alpine:latest

WORKDIR /backend

COPY --from=0 /backend/server ./

EXPOSE 8888

ENTRYPOINT ["./server"]
