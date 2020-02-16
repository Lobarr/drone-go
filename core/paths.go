package core

import (
	"os"
	"path"
)

//getDronePath returns drone workspace directory path
func getDronePath() string {
	return path.Join(os.Getenv("HOME"), ".drone")
}

//getDroneDBPath returns drone db directory path
func getDroneDBPath() string {
	return path.Join(getDronePath(), "db")
}

//getDroneDownloadsPath returns drone dowload path
func getDroneDownloadsPath() string {
	return path.Join(getDronePath(), "downloads")
}
