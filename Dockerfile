# Compilation
FROM golang:alpine as build
WORKDIR /gowooks
COPY ./src ./src
COPY ./go.mod ./go.mod
RUN go mod download
RUN go build -o /bin/gowooks ./src

# Runtime
FROM golang:alpine
COPY --from=build /bin/gowooks /go/bin/gowooks
COPY .env /go/.env
ENTRYPOINT [ "./bin/gowooks" ]