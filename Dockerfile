FROM golang:latest as builder

WORKDIR /app/ms-user-portal

ADD . /app/ms-user-portal/

RUN go mod download
RUN go mod vendor
RUN go build -o ms-user-portal .


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ms-user-portal/ms-user-portal /app/
COPY --from=builder /app/ms-user-portal/config.yml /app/

RUN apk add libc6-compat

CMD ./ms-user-portal

EXPOSE 4000