package management_frame

type ManagementFrame struct {
	FrameType string
	BSSID     string
	APName    string
}

func New(json string) *ManagementFrame {
	return &ManagementFrame{}
}
