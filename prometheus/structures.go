package prometheus

type Host struct {
	Name  string // job name
	Value int    // number of such jobs configured in Prometheus
}

var JobNames = make([]string, 2, 3)

func GetJobsNames() []string {
	return JobNames
}

func Init() {
	JobNames[0] = "nodes"
	JobNames[1] = "apache"
	JobNames = append(JobNames, "haproxy")
	JobNames = append(JobNames, "blackbox")
}
