# Echo Pulse - UDP Echo Server

`echo-pulse` is a simple UDP echo server that logs incoming messages and echoes them back to the client. It comes with unique IP tracking and string blacklist features. Easily configurable and set up with systemd for auto-start on system boot.

## Features

- **UDP Echoing**: Reflects incoming UDP messages back to the client.
- **Unique IP Logging**: Logs unique IPs that connect to the server.
- **String Blacklist**: Blocks and doesn't echo back specific blacklisted strings.
- **Easy Configuration**: Uses a TOML config for easy setup.

## Prerequisites

- Go 1.16+ or newer
- Systemd (for auto-start on boot)
- Proper permissions (for installing to `/usr/bin` and managing systemd services)

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/SiirRandall/echo-pulse.git
   cd echo-pulse
   ```

2. **Run the Installation Script**:
   ```bash
   sudo ./install.sh
   ```

3. **Configure the Service:**
  - Edit the configuration file located at `/etc/echopulse/config.conf` to your needs.
  - The default configuration should work for most setups.

4. **Start the Service**
   ```bash
   sudo systemctl start echo-pulse.service
   ``` 

5. **Optional) Enable Auto-Start at Boot:**
   ```bash
    sudo systemctl enable echo-pulse.service
    ```
## Uninstallation

To remove `echo-pulse`, simply run the provided uninstall script:
```bash
sudo ./uninstall.sh
```
## Logging

You can access the logs in the defult location `/var/logs/echopuls/` 
This can be changed in `/etc/echopulse/config.conf`