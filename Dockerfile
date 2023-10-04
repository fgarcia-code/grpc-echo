# Build executable binary
FROM golang:1.18 AS stage
WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /bin/service ./service.go
RUN chmod +x /bin/service

# Build docker image
FROM scratch
COPY --from=stage /bin/service /bin/service
ENTRYPOINT [ "/bin/service" ]