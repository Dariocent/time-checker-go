package time_checker

import (
	"net/http"
)

func TimeCheckerHandler(w http.ResponseWriter, r *http.Request) {
	//write the current time and the pod name
	env_information := getEnvInformation()
	w.Write([]byte(
		"The current time is: " + env_information["Time"] + "\n" +
			"POD_NAME: " + env_information["POD_NAME"] + "\n" +
			"POD_NAMESPACE: " + env_information["POD_NAMESPACE"] + "\n" +
			"POD_IP: " + env_information["POD_IP"] + "\n" +
			"NODE_NAME: " + env_information["NODE_NAME"] + "\n" +
			"NODE_IP: " + env_information["NODE_IP"] + "\n",
	))
}
