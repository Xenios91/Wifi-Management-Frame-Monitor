# Wifi-Management-Frame-Monitor

[![CodeQL](https://github.com/Xenios91/Wifi-Management-Frame-Monitor/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/Xenios91/Wifi-Management-Frame-Monitor/actions/workflows/codeql-analysis.yml)

## General Information
- Author: Corey Hartman
- Language: Golang v1.17
- Description: The Wifi-Management-Frame-Monitor is a microservice for receiving json post request representative of management
 frames receieved by a passive listening device for 802.11, such as a wireless NIC in monitor mode. Upon receiving the management frame, the tool will check known wireless access points against the management frame to determine if it is a evil twin, or if there is a flood of deauthentication frames indicating a potential deauthentication attack. Upon detection, the notification queue will be populated with an alert/notification allowing an HTTP GET request to the tool's API to get information on the attack, such as time and the associated ESSID.

 ## Video URL
https://youtu.be/ebdLkbrUkBM

 ## Installation
 - Requires Golang 1.17
 - Within the root directory run ```go build .```

## API Documentation

https://github.com/Xenios91/Wifi-Management-Frame-Monitor/blob/main/monitor_tool_swagger.pdf
