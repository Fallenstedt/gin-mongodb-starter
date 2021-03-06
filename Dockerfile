FROM golang:1.14.6-alpine AS builder
RUN apk update && \ 
    apk add --no-cache wget && \ 
    apk add make && \
    apk add git


# Deps
WORKDIR /app
COPY Makefile .
COPY go.mod .
COPY go.sum .
COPY VERSION . 
RUN make install

# Build
COPY . .
RUN make build

FROM alpine
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
RUN apk add --no-cache ca-certificates

EXPOSE 8080

# # Set the binary as the entrypoint of the container
RUN source ./
CMD ["./main"]
