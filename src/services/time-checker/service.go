package time_checker

import (
	"os"
	"time"
)

func getEnvInformation() map[string]string {
	return map[string]string{
		"Time":          time.Now().Format(time.RFC3339),
		"POD_NAME":      os.Getenv("POD_NAME"),
		"POD_NAMESPACE": os.Getenv("POD_NAMESPACE"),
		"POD_IP":        os.Getenv("POD_IP"),
		"NODE_NAME":     os.Getenv("NODE_NAME"),
		"NODE_IP":       os.Getenv("NODE_IP"),
	}

}
