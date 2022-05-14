# Alpine is lightweight image
FROM golang:1.18-alpine as builder
WORKDIR /app
# Copy files to download modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download                 # Multi docker stage wont work in jenkins, so dont use go mod download
# Copy all other files and build .exe 
COPY . ./
RUN go build -o /app/ser

# Use multi-stage to reduce docker image size
FROM alpine:latest AS production
WORKDIR /app
COPY --from=builder /app .
ENV PORT 80
ENV HOST 0.0.0.0
CMD ["/app/ser"]