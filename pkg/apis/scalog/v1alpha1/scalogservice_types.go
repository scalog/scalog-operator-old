package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file

type ClusterPhase string

const (
	ClusterPhaseNone     ClusterPhase = ""
	ClusterPhaseCreating              = "Creating"
	ClusterPhaseRunning               = "Running"
	ClusterPhaseFailed                = "Failed"
)

// ScalogServiceSpec defines the desired state of ScalogService
type ScalogServiceSpec struct {
	// NumShards denotes the number of shards of the log (data layer).
	NumShards int `json:"numShards"`
	// NumDataReplica denotes the number of replicas in each shard (data layer).
	NumDataReplica int `json:"numDataReplica"`
	// NumMetadataReplica denotes the number of replicas in the ordering layer (Paxos/Raft/...).
	NumMetadataReplica int `json:"numMetadataReplica"`
	// MillisecondBatchInterval denotes the time intervals (in milliseconds) at which the data and order layer report to each other.
	MillisecondBatchInterval int `json:"millisecondBatchInterval"`
	// Repository is the name of the repository that hosts scalog containder images.
	//
	// By default, it is `hub.docker.com/scalog/scalog`.
	Repository string `json:"repository,omitempty"`
	// Version is the expected version of the scalog cluster.
	// The scalog-operator will eventually make the scalog cluster version equal to the expected version.
	//
	// If version is not set, default is "0.0.1".
	Version string `json:"version,omitempty"`
	// Paused is to pause the control of the operator for the scalog cluster.
	Paused bool `json:"paused,omitempty"`
	// scalog cluster TLS configuration
	TLS *TLSPolicy `json:"TLS,omitempty"`
}

// ScalogServiceStatus defines the observed state of ScalogService
type ScalogServiceStatus struct {
	// Phase is the cluster running phase.
	Phase ClusterPhase `json:"phase"`
	// Paused indicates the operator pauses the control of the cluster.
	Paused bool `json:"paused,omitempty"`
	// NumShards denotes the current number of shards of the log (data layer).
	NumShards int `json:"numShards"`
	// NumDataReplica denotes the current number of replicas in each shard (data layer).
	// It should not be modified at runtime.
	NumDataReplica int `json:"numDataReplica"`
	// NumMetadataReplica denotes the current number of replicas in the ordering layer (Paxos/Raft/...).
	NumMetadataReplica int `json:"numMetadataReplica"`
	// ClientPort is the port for scalog client to access.
	ClientPort int `json:"clientPort,omitempty"`
	// ServiceName is the service for accessing scalog nodes.
	ServiceName string `json:"serviceName,omitempty"`
	// LatestShardID is the last value used to denote a new data shard
	LatestShardID int `json:"latestShardID"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScalogService is the Schema for the scalogservices API
// +k8s:openapi-gen=true
type ScalogService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScalogServiceSpec   `json:"spec,omitempty"`
	Status ScalogServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScalogServiceList contains a list of ScalogService
type ScalogServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScalogService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ScalogService{}, &ScalogServiceList{})
}
