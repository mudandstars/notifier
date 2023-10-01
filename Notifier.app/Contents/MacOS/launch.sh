#!/bin/bash
rm "$(dirname "$0")/notifier_error.log"
rm "$(dirname "$0")/notifier_output.log"

NOTIFIER_PATH="$(dirname "$0")/notifier"  # Assumes notifier is in the same directory as this script
CLIENT_APP_PATH="$(dirname "$0")/client.app"  # Assumes client.app is in the same directory as this script

NOTIFIER_PID=$(pgrep -f "notifier")

if [ -z "$NOTIFIER_PID" ]; then
    "$NOTIFIER_PATH" > "$(dirname "$0")/notifier_output.log" 2> "$(dirname "$0")/notifier_error.log" &

    NOTIFIER_PID=$!
fi

open "$CLIENT_APP_PATH"
