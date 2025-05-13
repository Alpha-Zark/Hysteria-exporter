# Hysteria-exporter
Hysteria-exporter

ENV
```
METRICS_INTERVAL:
# Default: 15 seconds.
# Description: Specifies the interval (in seconds) for collecting metrics. It can be left empty, in which case the default value will be used.

API_URL:
# Default: 127.0.0.1.
# Description: The API's base address. If not provided, it defaults to 127.0.0.1.

PORT:
# Default: 9980.
# Description: The port on which the service will listen. Defaults to 9980 if not set.

SECRET:
# Required: This value cannot be empty.
# Description: A secret key used for authentication or encryption purposes. Ensure this is properly set for secure operations.
```
Monitoring Metrics:

•	hysteria2_traffic_usage: Usage

•	hysteria2_online_user: Online users