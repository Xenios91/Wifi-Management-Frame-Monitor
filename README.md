# Wifi-Management-Frame-Monitor
## General Information
- Author: Corey Hartman
- Language: Golang v1.17
- Description: The Wifi-Management-Frame-Monitor is capable of receiving json post request representative of management
 frames receieved by a passive listening device for 802.11. Upon receiving the management frame, the tool will check known wireless access points against the management frame to determine if it is a rogue access point, or if there is a flood of deauthentication frames indicating a potential deauthentication attack. Upon detection, the notification queue will be populated with the alert allowing a request to the tools API to get information on the attack, such as time and the associated ESSID.

 ##Video URL

 ## Installation
 - Requires Golang 1.17
 - Within the root directory run ```go build .```


 See monitor_tool_swagger.pdf for the API documentation
