FROM golang:1.23.2-alpine AS base

# copy app-go
WORKDIR /home/app
COPY ./ ./

# initialize go workspace
WORKDIR /home/app
RUN go mod download 

FROM base AS build

WORKDIR /home/app
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app


# rebuild the application binary into a lean image
FROM gcr.io/distroless/base-debian13 AS release

WORKDIR /
COPY --from=build /home/app/bin/app /app
COPY --from=gcr.io/distroless/base:debug /busybox/wget /

EXPOSE 80