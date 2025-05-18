package contract

// DashboardInfo represents the dashboard information
type DashboardInfo struct {
	Users       int `json:"users"`
	Requests    int `json:"requests"`
	Resumes     int `json:"resumes"`
	Departments int `json:"departments"`
}
