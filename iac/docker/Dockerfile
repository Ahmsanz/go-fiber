FROM golang:1.21

WORKDIR /go/src/go-fiber

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /server



# Run
CMD ["/server"]
