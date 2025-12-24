#!/bin/sh

OUTPUT_DIR=./bin
GO_FILE_PATH="./cmd/server/main.go"
PROGRAM_NAME=RenewCMS
WEB_DIR="./web"

echo "Starting frontend build in $WEB_DIR..."
if [ -d "$WEB_DIR" ]; then
    (cd "$WEB_DIR" && bun run build)
    if [ $? -ne 0 ]; then
        echo "Frontend build failed. Aborting process."
        exit 1
    fi
else
    echo "Error: Directory $WEB_DIR not found."
    exit 1
fi

echo "Frontend build completed successfully."
echo "----------------------------------------"

GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

output_name="$OUTPUT_DIR/${PROGRAM_NAME}_${GOOS}_${GOARCH}"
if [ "$GOOS" = "windows" ]; then
    output_name+='.exe'
fi

echo "Detected platform: $GOOS/$GOARCH"
echo "Building for $GOOS/$GOARCH..."

env GOOS=$GOOS GOARCH=$GOARCH go build -o "$output_name" "$GO_FILE_PATH"

if [ $? -eq 0 ]; then
    echo "Compilation successful: $output_name"
else
    echo "Compilation for $GOOS/$GOARCH failed."
    exit 1
fi
