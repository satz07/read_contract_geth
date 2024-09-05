FROM golang:1.22.2-alpine as builder
WORKDIR /app
COPY . .
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN chmod +x ./main


FROM nginx:alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/abi ./abi
EXPOSE 80
CMD ["./main"]