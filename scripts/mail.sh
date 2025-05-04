# This script is used to start the Mailpit service for local email testing.
#!/usr/bin/env bash
# Check if Mailpit is already running
if pgrep -x "mailpit" > /dev/null; then
    echo "Mailpit is already running."
    exit 1
fi

# Start Mailpit

mailpit --smtp :8024 --listen :8025
