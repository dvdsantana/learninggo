FROM golang:1.24-alpine AS dependencies
WORKDIR /app
COPY go.* ./
RUN go mod download

FROM dependencies AS build
COPY . ./
RUN CGO_ENABLED=0 go build -o /main -ldflags="-w -s" .

FROM golang:1.24-alpine AS runtime
COPY --from=build /main /main
CMD ["/main"]