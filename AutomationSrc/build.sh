#!/bin/sh

SCRIPT_DIR="$(dirname "$(realpath "$0")")"

cd "${SCRIPT_DIR}" || exit 1

go build -o ./..

chmod +x ./../GradientTheme

exit 0