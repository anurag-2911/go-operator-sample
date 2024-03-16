package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TimeScalerSpec defines the desired state of TimeScaler
type TimeScalerSpec struct {
	// Schedule is a cron-like expression that specifies when to perform scaling actions
	Schedule string `json:"schedule,omitempty"`

	// Replicas is the desired number of replicas to scale to during the scheduled time
	Replicas int32 `json:"replicas"`
}

// TimeScalerStatus defines the observed state of TimeScaler
type TimeScalerStatus struct {
	// LastScaledTime is the last time the scaling action was performed
	LastScaledTime *metav1.Time `json:"lastScaledTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TimeScaler is the Schema for the timescalers API
type TimeScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TimeScalerSpec   `json:"spec,omitempty"`
	Status TimeScalerStatus `json:"status,omitempty"`
}

// DeepCopyObject implements runtime.Object.
func (t *TimeScaler) DeepCopyObject() runtime.Object {
	panic("unimplemented")
}

// GetObjectKind implements runtime.Object.
// Subtle: this method shadows the method (TypeMeta).GetObjectKind of TimeScaler.TypeMeta.
func (t *TimeScaler) GetObjectKind() schema.ObjectKind {
	panic("unimplemented")
}

//+kubebuilder:object:root=true

// TimeScalerList contains a list of TimeScaler
type TimeScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TimeScaler `json:"items"`
}

// DeepCopyObject implements runtime.Object.
func (t *TimeScalerList) DeepCopyObject() runtime.Object {
	panic("unimplemented")
}

// GetObjectKind implements runtime.Object.
// Subtle: this method shadows the method (TypeMeta).GetObjectKind of TimeScalerList.TypeMeta.
func (t *TimeScalerList) GetObjectKind() schema.ObjectKind {
	panic("unimplemented")
}

func init() {
	SchemeBuilder.Register(&TimeScaler{}, &TimeScalerList{})
}
