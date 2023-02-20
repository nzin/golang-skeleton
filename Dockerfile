######################################
# Prepare npm_builder
######################################
FROM node:16 as npm_builder
WORKDIR /go/src/github.com/nzin/golang-skeleton
ADD . .
RUN make build_ui

######################################
# Prepare go_builder
######################################
FROM golang:1.18 as go_builder
WORKDIR /go/src/github.com/nzin/golang-skeleton
ADD . .
RUN make build

######################################
# Copy from builder to debian image
######################################
FROM debian:bullseye-slim
RUN mkdir /app
WORKDIR /app

ENV HOST=0.0.0.0
ENV PORT=18000

COPY --from=go_builder /go/src/github.com/nzin/golang-skeleton/golang-skeleton ./golang-skeleton
COPY --from=npm_builder /go/src/github.com/nzin/golang-skeleton/browser/golang-skeleton-ui/dist ./browser/golang-skeleton-ui/dist

RUN useradd --uid 1000 --gid 0 golang-skeleton && \
    chown golang-skeleton:root /app && \
    chmod g=u /app
USER 1000:0

EXPOSE 18000
CMD ./golang-skeleton
