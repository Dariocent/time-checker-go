package temp

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	//write the current time and the pod name
	w.Write([]byte(
		"The current time is: " + getEnvInformation()["Time"] + "\n" +
			"POD_NAME: " + getEnvInformation()["POD_NAME"] + "\n" +
			"POD_NAMESPACE: " + getEnvInformation()["POD_NAMESPACE"] + "\n" +
			"POD_IP: " + getEnvInformation()["POD_IP"] + "\n" +
			"NODE_NAME: " + getEnvInformation()["NODE_NAME"] + "\n" +
			"NODE_IP: " + getEnvInformation()["NODE_IP"] + "\n",
	))
}
