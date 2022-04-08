FROM golang:1.18-buster as builder

WORKDIR /go/src/room_activity

# Download dependancies
COPY ./go.* ./
RUN go mod download

# Build executable
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo -mod=readonly \
    -o /go/bin/room_activity

FROM busybox:stable

COPY --from=builder /go/bin/room_activity /bin/room_activity

EXPOSE 6600
ENTRYPOINT [ "/bin/room_activity" ]