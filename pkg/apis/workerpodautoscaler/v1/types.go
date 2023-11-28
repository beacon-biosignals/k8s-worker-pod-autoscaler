package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkerPodAutoScaler is a specification for a WorkerPodAutoScaler resource
type WorkerPodAutoScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkerPodAutoScalerSpec   `json:"spec"`
	Status WorkerPodAutoScalerStatus `json:"status"`
}

// WorkerPodAutoScalerSpec is the spec for a WorkerPodAutoScaler resource
type WorkerPodAutoScalerSpec struct {
	MinReplicas             *int32   `json:"minReplicas"`
	MaxReplicas             *int32   `json:"maxReplicas"`
	MaxDisruption           *string  `json:"maxDisruption,omitempty"`
	QueueURI                string   `json:"queueURI"`
	QueueServiceName        string   `json:"queueServiceName"`
	DeploymentName          string   `json:"deploymentName,omitempty"`
	ReplicaSetName          string   `json:"replicaSetName,omitempty"`
	TargetMessagesPerWorker *int32   `json:"targetMessagesPerWorker"`
	SecondsToProcessOneJob  *float64 `json:"secondsToProcessOneJob,omitempty"`
}

// WorkerPodAutoScalerStatus is the status for a WorkerPodAutoScaler resource
type WorkerPodAutoScalerStatus struct {
	CurrentMessages   int32 `json:"CurrentMessages"`
	CurrentReplicas   int32 `json:"CurrentReplicas"`
	AvailableReplicas int32 `json:"AvailableReplicas"`
	DesiredReplicas   int32 `json:"DesiredReplicas"`

	// LastScaleTime is the last time the WorkerPodAutoscaler scaled the workers
	// It is used by the autoscaler to control
	// how often the number of pods is changed.
	// +optional
	LastScaleTime *metav1.Time `json:"LastScaleTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkerPodAutoScalerList is a list of WorkerPodAutoScaler resources
type WorkerPodAutoScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []WorkerPodAutoScaler `json:"items"`
}
