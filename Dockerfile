FROM alpine:3.7
WORKDIR /go/src/app
COPY server .
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
EXPOSE 50051
CMD ["./server"]