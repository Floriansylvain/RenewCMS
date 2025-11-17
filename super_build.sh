#!/bin/sh

OUTPUT_DIR=./bin
GO_FILE_PATH="./main/main.go"
PROGRAM_NAME=RenewCMS

platforms=("windows/amd64" "windows/arm64" "linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name="$OUTPUT_DIR/${PROGRAM_NAME}_${GOOS}_${GOARCH}"
    
    if [ "$GOOS" = "windows" ]; then
        output_name+='.exe'
    fi

    echo "Building for $GOOS/$GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o "$output_name" "$GO_FILE_PATH"
    if [ $? -ne 0 ]; then
        echo "Compilation for $GOOS/$GOARCH failed."
        exit $?
    fi
done

echo "Compilation successful for all platforms."
