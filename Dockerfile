FROM golang:latest AS build

WORKDIR /app
ADD . /app
ENV CGO_ENABLED=0
RUN go build -o main .

FROM scratch
COPY --from=build /app/main /main
CMD ["/main"]
