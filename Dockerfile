# Build the application from source
FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /dermsnap

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

ARG DATABASE_URL
ARG JWT_SECRET
ENV APP_ENV="development"
ENV DOXIMITY_PROVIDER_BASE_URL="https://auth.doximity.com"

WORKDIR /

COPY --from=build-stage /dermsnap /dermsnap
COPY ./assets /assets
COPY ./app/views /app/views

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/dermsnap", "run-app"]