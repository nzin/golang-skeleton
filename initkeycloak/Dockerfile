######################################
# Prepare go_builder
######################################
FROM golang:1.17 as go_builder
WORKDIR /go/src/github.com/nzin/golang-skeleton/initkeycloak
ADD . .
RUN go build .

######################################
# Copy from builder to debian image
######################################
FROM debian:bullseye-slim
RUN mkdir /app
WORKDIR /app

COPY --from=go_builder /go/src/github.com/nzin/golang-skeleton/initkeycloak/initkeycloak ./initkeycloak

RUN useradd --uid 1000 --gid 0 initkeycloak && \
    chown initkeycloak:root /app && \
    chmod g=u /app
USER 1000:0

CMD ./initkeycloak
