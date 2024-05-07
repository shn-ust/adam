package adm

import "UST-FireOps/adam/utils"

// Check if both services are up
// A service is defined as a tuple (ip, port)
// This function is used to determine if a service is a long running one
// And not one which has an ephemeral port
// Returns true if both the services are up
func CheckStatus(dep Dependency) bool {
	firstServer := utils.IsServerUp(dep.SrcIP, dep.SrcPort)
	secondServer := utils.IsServerUp(dep.DestIP, dep.DestPort)
	return firstServer && secondServer
}
