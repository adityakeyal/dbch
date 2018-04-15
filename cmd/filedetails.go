package cmd

//FileDetails Any change in file consists of 4 attributes
type FileDetails struct {
	filename string
	changes  []ChangeDetails
}

//ChangeDetails - Struct for changes info
type ChangeDetails struct {
	source string
	target string
}

//Replace replacer
var Replace = make([]FileDetails, 0, 4)

func init() {

	Replace = append(Replace, FileDetails{"xenos-jdbc.properties", []ChangeDetails{
		ChangeDetails{"jdbc.components.url.GLOBAL", "jdbc:oracle:thin:@localhost:1521:XE"},
		ChangeDetails{"jdbc.components.userName.GLOBAL", "gvth_dev_global"},
		ChangeDetails{"jdbc.components.password.GLOBAL", "gvth_dev_global"},
	}})

	Replace = append(Replace, FileDetails{"GMO-jdbc.properties", []ChangeDetails{
		ChangeDetails{"jdbc.components.url.GLOBAL", "jdbc:oracle:thin:@localhost:1521:XE"},
		ChangeDetails{"jdbc.components.userName.GLOBAL", "gvth_dev_gmo_txn"},
		ChangeDetails{"jdbc.components.password.GLOBAL", "gvth_dev_gmo_txn"},
	}})

	Replace = append(Replace, FileDetails{"NRI-jdbc.properties", []ChangeDetails{
		ChangeDetails{"jdbc.components.url.GLOBAL", "jdbc:oracle:thin:@localhost:1521:XE"},
		ChangeDetails{"jdbc.components.userName.GLOBAL", "gvth_dev_nri_txn"},
		ChangeDetails{"jdbc.components.password.GLOBAL", "gvth_dev_nri_txn"},
	}})

}
