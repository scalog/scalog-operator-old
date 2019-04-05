package scalogservice

import (
	"strconv"
	"strings"
)

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

func getShardIDFromStatefulSetName(name string) (int32, error) {
	splitPodName := strings.Split(name, "-")
	shardID, err := strconv.ParseInt(splitPodName[len(splitPodName)-1], 10, 32)
	return int32(shardID), err
}
