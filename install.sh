#!/bin/bash

# Check if the script is run as root
if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit
fi

# Step 1: Build the Go binary
echo "Building Go binary..."
go build -o echo-pulse

# Step 2: Move binary to /usr/bin
echo "Installing binary to /usr/bin..."
mv echo-pulse /usr/bin/

# Step 3: Create Systemd service file
echo "Creating systemd service file..."
cat <<EOL > /etc/systemd/system/echo-pulse.service
[Unit]
Description=Echo Pulse Service
After=network.target

[Service]
ExecStart=/usr/bin/echo-pulse
Restart=always

[Install]
WantedBy=multi-user.target
EOL

# Step 4: Create configuration directory and default config file
echo "Setting up configuration..."
mkdir -p /etc/echopulse
cat <<EOL > /etc/echopulse/config.conf
[Server]
IP = "0.0.0.0"
Port = 5060

[Log]
Directory = "/var/log/echopulse"

[Blacklist]
# Example Keywords = ["maliciousString1", "spamString2", "bannedContent3"]
Keywords = []
EOL

echo "Installation completed."