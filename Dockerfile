# Build
# The version of golang is locked to 1.19 like it is defined in the go.mod file.
FROM golang:1.19-alpine as pokergame_builder


# The module inside LARVIS is installed under pokergame folder
WORKDIR /pokergame

COPY . .

# Run as non-root user to minimize attack surface.
RUN go build -o larvis && \
    addgroup -S larvisgroup && \
    adduser -S larvisuser -G larvisgroup && \
    chown -R larvisuser:larvisgroup /pokergame



# Run
# Use locked version in order to avoid breaking changes
# The image is locked on major version 3 and minor version is not locked.
# This is so that the possible security updates will be applied automatically.
# but also breaking changes will not be applied automatically.
FROM alpine:3

WORKDIR /pokergame

COPY --from=pokergame_builder /pokergame/larvis .
COPY --from=pokergame_builder /pokergame/configs/config.json ./configs/config.json

RUN addgroup -S larvisgroup && \
    adduser -S larvisuser -G larvisgroup && \
    chown -R larvisuser:larvisgroup /pokergame

USER larvisuser

ENTRYPOINT [ "./larvis" ]

