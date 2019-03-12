package sync

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kube-incubator/kube-operator-helper/syncer"
	"github.com/kube-incubator/tomcat-operator/pkg/scheme/tomcat"
)

// NewServiceSyncer returns a new sync.Interface for reconciling Tomcat Service
func NewServiceSyncer(tc *tomcat.Tomcat, c client.Client, scheme *runtime.Scheme) syncer.Interface {
	objLabels := tc.ComponentLabels(tomcat.TomcatDeployment)

	obj := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tc.Name,
			Namespace: tc.Namespace,
		},
	}

	return syncer.NewObjectSyncer("Service", tc.Unwrap(), obj, c, scheme, func(existing runtime.Object) error {
		out := existing.(*corev1.Service)
		out.Labels = labels.Merge(labels.Merge(out.Labels, objLabels), controllerLabels)

		selector := tc.TomcatServerPodLabels()
		if !labels.Equals(selector, out.Spec.Selector) {
			if out.ObjectMeta.CreationTimestamp.IsZero() {
				out.Spec.Selector = selector
			} else {
				return fmt.Errorf("service selector is immutable")
			}
		}

		if len(out.Spec.Ports) != 1 {
			out.Spec.Ports = make([]corev1.ServicePort, 1)
		}

		out.Spec.Ports[0].Name = "http"
		out.Spec.Ports[0].Port = *tc.Spec.ServicePort
		out.Spec.Ports[0].TargetPort = intstr.FromInt(tomcat.TomcatHTTPPort)

		return nil
	})
}
