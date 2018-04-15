package cmd

//SystemInfo - List of environments applicable for a system
type SystemInfo struct {
	Environments []EnvironmentInfo
}

//EnvironmentInfo -
type EnvironmentInfo struct {
	Default         bool
	EnvironmentName string
	ReplaceInfo     []FileChanges
}

//FileChanges Any change in file consists of 4 attributes
type FileChanges struct {
	Filename string
	Changes  []ChangeDetails
}

//ChangeDetails - Struct for changes info
type ChangeDetails struct {
	Source string
	Target string
}

//Replace replacer

var DevEnvironment EnvironmentInfo

func init() {

	DevEnvironment.EnvironmentName = "Dev"
	DevEnvironment.Default = false
	DevEnvironment.ReplaceInfo = append(DevEnvironment.ReplaceInfo, FileChanges{"xenos-jdbc.properties", []ChangeDetails{
		ChangeDetails{"jdbc.components.url.GLOBAL", "jdbc:oracle:thin:@localhost:1521:XE"},
		ChangeDetails{"jdbc.components.userName.GLOBAL", "gvth_dev_global"},
		ChangeDetails{"jdbc.components.password.GLOBAL", "gvth_dev_global"},
	}})

	DevEnvironment.ReplaceInfo = append(DevEnvironment.ReplaceInfo, FileChanges{"GMO-jdbc.properties", []ChangeDetails{
		ChangeDetails{"jdbc.components.url.GLOBAL", "jdbc:oracle:thin:@localhost:1521:XE"},
		ChangeDetails{"jdbc.components.userName.GLOBAL", "gvth_dev_gmo_txn"},
		ChangeDetails{"jdbc.components.password.GLOBAL", "gvth_dev_gmo_txn"},
	}})

	DevEnvironment.ReplaceInfo = append(DevEnvironment.ReplaceInfo, FileChanges{"NRI-jdbc.properties", []ChangeDetails{
		ChangeDetails{"jdbc.components.url.GLOBAL", "jdbc:oracle:thin:@localhost:1521:XE"},
		ChangeDetails{"jdbc.components.userName.GLOBAL", "gvth_dev_nri_txn"},
		ChangeDetails{"jdbc.components.password.GLOBAL", "gvth_dev_nri_txn"},
	}})

}
