#go version go1.21.5 linux/arm64
FROM cgr.dev/chainguard/go@sha256:b9ab4040eedba24a93a84fa5b9e5ee736b72f4072b31d4da01d5f861e1529dee AS builder
COPY . /app/
RUN cd /app && go build -o ./gx cmd/*

FROM cgr.dev/chainguard/glibc-dynamic@sha256:47e11439e9b2c58ef80cb7db66c4191acc6e61b549f4f1d8d4654b766dc20c0e
FROM golang:1.21.5-bookworm
COPY --from=builder /app/gx /usr/bin/
CMD ["/usr/bin/gx"]
