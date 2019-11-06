FROM alpine:latest AS builder
COPY .. /SmartLocker/
WORKDIR /SmartLocker
ENV GO111MODULE on
RUN apk update && apk add zip go make ca-certificates git musl-dev
RUN make pack

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /SmartLocker/SmartLocker/ ./
COPY --from=builder /SmartLocker/docker-startup.sh .
RUN apk add --no-cache tzdata
RUN chmod +x docker-startup.sh
CMD ["./docker-startup.sh"]