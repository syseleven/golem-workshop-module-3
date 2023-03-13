FROM golang:alpine3.17 AS buildstage
ADD . /app
WORKDIR /app
RUN go mod init metrics_example \
    && go mod tidy \
    && go build -o metricsapp

FROM alpine:3.17
WORKDIR /app
ENV PORT 8090
EXPOSE 8090
COPY --from=buildstage /app/metricsapp /app
CMD ["/app/metricsapp"]