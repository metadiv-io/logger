package constant

import "github.com/metadiv-io/env"

var (
	// LOG_FOLDER_PATH is the path to the folder where the logs will be stored
	// default value is ./logs
	// You can change the value by setting the environment variable LOG_FOLDER_PATH
	LOG_FOLDER_PATH = env.String("LOG_FOLDER_PATH", "./logs")
)
