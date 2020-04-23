# An implementation of LLRP

This is intended to implement a simple command-line LLRP configuration tool for setting IP addresses and such for Miranda devices.  It should also be useful for other LLRP devices.

LLRP is a part of ANSI E1.33, this is based off the standard dated 9/02/2019.

# Current Status

Not really built out yet.

## Supported Commands

- DISCOVERY
- DEVICE_LABEL (get/set)
- FACTORY_DEFAULTS
- IDENTIFY_DEVICE on/off
- RESET_DEVICE
- IPV4_DHCP_MODE
- IPV4_CURRENT_ADDRESS
- IPV4_STATIC_ADDRESS
- IPV4_DEFAULT_ROUTE
