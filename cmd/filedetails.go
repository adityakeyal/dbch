package cmd

//SystemInfo - List of environments applicable for a system
type SystemInfo struct {
	Environments []*EnvironmentInfo `json:"environments"`
}

//EnvironmentInfo -
type EnvironmentInfo struct {
	Default         bool           `json:"default"`
	EnvironmentName string         `json:"environmentName"`
	ReplaceInfo     []*FileChanges `json:"replaceInfo"`
}

//FileChanges Any change in file consists of 4 attributes
type FileChanges struct {
	Filename string           `json:"fileName"`
	Changes  []*ChangeDetails `json:"changes"`
	FileType string           `json:"fileType"`
}

//ChangeDetails - Struct for changes info
type ChangeDetails struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

//GetDefaultEnvironment - Function to get the first available Environment
func (system *SystemInfo) GetDefaultEnvironment() *EnvironmentInfo {
	for _, env := range system.Environments {
		if env.Default {
			return env
		}

	}
	return system.Environments[0]
}

//GetEnvironment - fetches the environment information provided an environment string
func (system *SystemInfo) GetEnvironment(environment string) *EnvironmentInfo {
	if environment == "" {
		return system.GetDefaultEnvironment()
	}
	for _, env := range system.Environments {
		if env.EnvironmentName == environment {
			return env
		}

	}
	return system.GetDefaultEnvironment()
}
