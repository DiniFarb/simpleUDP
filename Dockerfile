FROM golang:1.18.0-alpine3.15 as builder
RUN mkdir /build
COPY . /build/
WORKDIR /build/
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -a -o udp-server .

FROM scratch
COPY --from=builder /build/udp-server .
EXPOSE 8866/udp
ENTRYPOINT [ "./udp-server" ]