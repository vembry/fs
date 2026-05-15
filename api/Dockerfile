FROM golang:1.26-alpine AS base

# copy api
WORKDIR /home/api
COPY ./ ./

# initialize go workspace
WORKDIR /home/api
RUN go mod download 

FROM base AS build

WORKDIR /home/api
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/api


# rebuild the application binary into a lean image
FROM gcr.io/distroless/base-debian13 AS release

WORKDIR /
COPY --from=build /home/api/bin/api /api
COPY --from=gcr.io/distroless/base:debug /busybox/wget /

EXPOSE 80