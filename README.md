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

## llrp_config commands

li : list computer IP interfaces
si : select computer IP interface 
d : run discovery
l # [label] : get/set device label (don't specify a label to just get)
factory # : reset to factory defaults
i # [on/off] : get identify status, or turn on/off
reset # : reset device
dhcp # [on/off] : get dhcp status, or turn on/off
c # : get current IP address
setip # [dotted quad] [netmask] : get static IP address, or set it
gw # [dotted quad] : set gateway address
p : print device table
r : refresh device table (re-runs discovery, then gets parameters and prints the table)
