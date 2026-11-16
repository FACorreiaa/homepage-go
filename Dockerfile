# Build stage:
# We install everything we need here, build the CSS, generate templ code,
# and compile the app binary.
FROM golang:1.27.3-bookworm AS build
WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates wget && \
    rm -rf /var/lib/apt/lists/*

ARG TARGETARCH
RUN case "${TARGETARCH:-amd64}" in \
      amd64) tailwind_arch="x64" ;; \
      arm64) tailwind_arch="arm64" ;; \
      *) echo "unsupported architecture: ${TARGETARCH}" >&2 && exit 1 ;; \
    esac && \
    wget -q --show-progress -O /usr/local/bin/tailwindcss \
      "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-${tailwind_arch}" && \
    chmod +x /usr/local/bin/tailwindcss

# Copy Go dependency files first so Docker can cache downloads better.
COPY --link go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

# Copy the rest of the project files.
COPY --link . .

# Build the Tailwind CSS file.
# We also include templUI component templates, so Tailwind sees their classes too.
RUN --mount=type=cache,target=/go/pkg/mod \
    TEMPLUI_PATH="$(go list -mod=mod -m -f '{{.Dir}}' github.com/templui/templui)" && \
    test -n "$TEMPLUI_PATH" && \
    test -d "$TEMPLUI_PATH/components" && \
    printf '%s\n' \
      '@source "./**/*.templ";' \
      '@source "./assets/static/*.js";' \
      "@source \"$TEMPLUI_PATH/components/button/**/*.templ\";" \
      "@source \"$TEMPLUI_PATH/components/icon/**/*.templ\";" \
      > ./assets/css/sources.generated.css && \
    tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify

# Generate Go files from .templ files.
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go tool templ generate

# Build one static Linux binary for the final image.
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app/main ./main.go

# Runtime stage:
# This image stays small and only contains the final app binary.
FROM alpine:3.20.2
WORKDIR /app

# Install runtime certificates and runtime client tools.
RUN apk add --no-cache ca-certificates curl

# Run the app in production mode.
ENV GO_ENV=production

# Copy the built binary from the build stage.
COPY --link --from=build /app/main .

# The app listens on port 8090.
EXPOSE 8090

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=5 \
  CMD curl -fsS http://127.0.0.1:8090/healthz || exit 1

# Start the app.
CMD ["./main"]
