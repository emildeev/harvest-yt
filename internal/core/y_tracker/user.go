package ytrackercore

type User struct {
	UID         int    `json:"uid"`
	Login       string `json:"login"`
	TrackerUID  int    `json:"trackerUid"`
	PassportUid int    `json:"passportUid"`
	CloudUid    string `json:"cloudUid"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Display     string `json:"display"`
	Email       string `json:"email"`
}
