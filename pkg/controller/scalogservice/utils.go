package scalogservice

import "strings"

/*
	Obtains a pointer to an int64 object -- I have not figured out a way
	to do this without a helper method.
*/
func createInt64(x int64) *int64 {
	return &x
}

func createInt32(x int32) *int32 {
	return &x
}

/*
	getIDFromShardName assumes that pod names are in the format,
	"stuff-stuff-{PODID}". This method just obtains the last
	hyphen delimited item in a string.
*/
func getIDFromShardName(podName string) string {
	splitPodName := strings.Split(podName, "-")
	shardID := splitPodName[len(splitPodName)-1]
	return shardID
}
