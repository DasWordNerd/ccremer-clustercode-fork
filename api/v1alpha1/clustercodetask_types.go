package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func init() {
	SchemeBuilder.Register(&ClustercodeTask{}, &ClustercodeTaskList{})
}

type (
	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:printcolumn:name="Source",type="string",JSONPath=".spec.sourceUrl",description="Source file name"
	// +kubebuilder:printcolumn:name="Target",type="string",JSONPath=".spec.targetUrl",description="Target file name"
	// +kubebuilder:printcolumn:name="Blueprint",type="string",JSONPath=`.metadata.ownerReferences[?(@.controller)].name`,description="Blueprint reference"
	// +kubebuilder:printcolumn:name="Slices",type="string",JSONPath=`.spec.slicesPlannedCount`,description="Clustercode Total Slices"
	// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

	// Blueprint is the Schema for the archives API
	ClustercodeTask struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   ClustercodeTaskSpec   `json:"spec,omitempty"`
		Status ClustercodeTaskStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// ClustercodeTaskList contains a list of ClusterCodeTasks
	ClustercodeTaskList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []ClustercodeTask `json:"items"`
	}

	// EncodingTaskSpec defines the desired state of ClustercodeTask.
	ClustercodeTaskSpec struct {
		TaskId               ClustercodeTaskId   `json:"taskId,omitempty"`
		Storage              StorageSpec         `json:"storage,omitempty"`
		SourceUrl            ClusterCodeUrl      `json:"sourceUrl,omitempty"`
		TargetUrl            ClusterCodeUrl      `json:"targetUrl,omitempty"`
		Suspend              bool                `json:"suspend,omitempty"`
		EncodeSpec           EncodeSpec          `json:"encodeSpec"`
		ServiceAccountName   string              `json:"serviceAccountName,omitempty"`
		FileListConfigMapRef string              `json:"fileListConfigMapRef,omitempty"`
		ConcurrencyStrategy  ClustercodeStrategy `json:"concurrencyStrategy,omitempty"`
		SlicesPlannedCount   int                 `json:"slicesPlannedCount,omitempty"`
	}

	ClustercodeTaskStatus struct {
		Conditions           []metav1.Condition    `json:"conditions,omitempty"`
		SlicesScheduledCount int                   `json:"slicesScheduledCount,omitempty"`
		SlicesFinishedCount  int                   `json:"slicesFinishedCount,omitempty"`
		SlicesScheduled      []ClustercodeSliceRef `json:"slicesScheduled,omitempty"`
		SlicesFinished       []ClustercodeSliceRef `json:"slicesFinished,omitempty"`
	}

	ClustercodeSliceRef struct {
		JobName    string `json:"jobName,omitempty"`
		SliceIndex int    `json:"sliceIndex"`
	}

	ClustercodeStrategy struct {
		ConcurrentCountStrategy *ClustercodeCountStrategy `json:"concurrentCountStrategy,omitempty"`
	}

	ClustercodeCountStrategy struct {
		MaxCount int `json:"maxCount,omitempty"`
	}
	ClustercodeTaskId string
)

const (
	ClustercodeTaskIdLabelKey = "clustercode.github.io/task-id"
)

func (id ClustercodeTaskId) AsLabels() labels.Set {
	return map[string]string{
		ClustercodeTaskIdLabelKey: id.String(),
	}
}

func (id ClustercodeTaskId) String() string {
	return string(id)
}
