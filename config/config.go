package config

var Config *config
var StorageConfig *storage_config

func init() {
	addr := "127.0.0.1:5432"
	DevConfig := db_config{addr, "resonate_dev_user", "password", "resonate_dev"}
	TestingConfig := db_config{addr, "resonate_testing_user", "", "resonate_testing"}

	StorageConfig = &storage_config{
		"3a17476c901e",
		"00257352ce3d0c8db20fdd511881f1e5e2508269ad",
		"https://api.backblazeb2.com/b2api/v1/b2_authorize_account",
		"/b2api/v1/b2_download_file_by_id?fileId=",
	}

	Config = &config{TestingConfig, DevConfig}
}

type db_config struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type config struct {
	Testing db_config `json:"testing"`
	Dev     db_config `json:"dev"`
}

type storage_config struct {
	AccountId    string
	Key          string
	AuthEndpoint string
	FileEndpoint string
}
