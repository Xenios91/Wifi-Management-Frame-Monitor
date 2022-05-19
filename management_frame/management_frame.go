package management_frame

type ManagementFrame struct {
	FrameType        string `json:"short_message"`
	TimeStamp        string `json:"time_stamp"`
	AntennaSignal    string `json:"antenna_signal"`
	Essid            string `json:"essid"`
	Bssid            string `json:"bssid"`
	BeaconInterval   string `json:"beacon_interval"`
	CurrentChannel   string `json:"current_channel"`
	CountryCode      string `json:"country_code"`
	TransmitPower    string `json:"transmit_power"`
	IsPrivateNetwork bool   `json:"is_private_network"`
}
