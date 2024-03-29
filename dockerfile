FROM golang:1.21.1-alpine

# Set destination for COPY
WORKDIR /wd

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./
COPY *.env ./
COPY api/ ./api

# Build
#RUN CGO_ENABLED=0 GOOS=linux go build -o /app
RUN go build -o /wd/app
#RUN chmod +x /app

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/wd/app"]