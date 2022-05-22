# BUILDER SECTION                                     
#######################################################

# Let's get an image with Goland installed
FROM golang:1.18-alpine AS builder

# Let's work on go/src
WORKDIR /go/src

# Let's cache GO packages
COPY go.mod .
RUN go mod download \
 && go mod verify

# Let's cache GO build
COPY . .
RUN go mod tidy \
 && go build -o /go/bin/service

# Let's cache OS stuff
RUN adduser -DH serviceuser \
 && chown serviceuser:serviceuser /go/bin/service

# APP SECTION                                     
#######################################################

# Let's start from zero
# TODO - Change to scratch
FROM alpine

# Yeah, this is us! Always happy to help.
LABEL maintainer=rmfagundes@gmail.com

# Let's copy our app binary
COPY --from=builder /go/bin/service /bin/service

# Let's start our app as entry poing
ENTRYPOINT ["service"]