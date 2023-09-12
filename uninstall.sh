#!/bin/bash

# Check if the script is run as root
if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit
fi

# Step 1: Stop the systemd service if it's running
echo "Stopping echo-pulse service..."
systemctl stop echo-pulse

# Step 2: Disable the systemd service from starting at boot
echo "Disabling echo-pulse service..."
systemctl disable echo-pulse

# Step 3: Remove the systemd service file
echo "Removing systemd service file..."
rm /etc/systemd/system/echo-pulse.service

# Optional: Reload systemd to reflect changes
systemctl daemon-reload

# Step 4: Remove the binary from /usr/bin
echo "Removing binary from /usr/bin..."
rm /usr/bin/echo-pulse

# Step 5: Remove the configuration directory and files
echo "Removing configuration directory and files..."
rm -r /etc/echopulse

echo "Uninstallation completed."

