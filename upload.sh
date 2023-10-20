#!/bin/bash
# Author: Gatovsky
# usage: ./upload.sh -a files -c "some message"

POSITIONAL_ARGS=()

while [[ "${#}" -gt 0 ]]; do
    case "${1}" in
        -a | --add)
            shift 1
            FILES="$@"
            shift
            ;;
        -c | --message)
            MESSAGE="$2"
            shift
            shift
            ;;
        --default)
            git add .; git commit -m "commit upload"; git push
            exit 0
            ;;
        -* | --*)
            echo "Unknown option $1"
            exit 1
            ;;
        *)
            POSITIONAL_ARGS+=("$1")
            shift
            ;;
    esac
done

set -- "${POSITIONAL_ARGS[@]}"

echo "FILES     =${FILES}"
echo "MESSAGE  =${MESSAGE}"

git add ${FILES}; git commit -m "${MESSAGE}"; git push



