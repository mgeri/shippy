
FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

#ENV GO111MODULE=on

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-cli-consignment
#RUN GOOS=linux go build -o shippy-service-consignment

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD consignment.json /app/consignment.json
COPY --from=builder /app/shippy-cli-consignment .

CMD ["./shippy-cli-consignment"]