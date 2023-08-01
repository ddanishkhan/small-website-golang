# Specifies a parent image
FROM golang:1.20-bullseye AS base

# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY . .
 
# Installs Go dependencies
RUN go mod download
RUN go mod verify
 
# Builds your app with optional configuration
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /godocker .

# Small Image
FROM gcr.io/distroless/static-debian11

COPY --from=base /godocker .
COPY resources resources/
COPY view.html .

# Tells Docker which network port your container listens on
EXPOSE 9000
 
# Specifies the executable command that runs when the container starts
CMD ["./godocker" ]