package sync

import (
	"fmt"
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/imdario/mergo"
	"github.com/kube-incubator/kube-operator-helper/mergo/transformers"
	"github.com/kube-incubator/kube-operator-helper/syncer"
	"github.com/kube-incubator/tomcat-operator/pkg/scheme/tomcat"
)

var (
	oneReplica int32 = 1
)

// NewDeploymentSyncer returns a new sync.Interface for reconciling tomcat Deployment
func NewDeploymentSyncer(tc *tomcat.Tomcat, c client.Client, scheme *runtime.Scheme) syncer.Interface {
	objLabels := tc.ComponentLabels(tomcat.TomcatDeployment)

	obj := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tc.ComponentName(tomcat.TomcatDeployment),
			Namespace: tc.Namespace,
		},
	}

	return syncer.NewObjectSyncer("Deployment", tc.Unwrap(), obj, c, scheme, func(existing runtime.Object) error {
		out := existing.(*appsv1.Deployment)
		out.Labels = labels.Merge(labels.Merge(out.Labels, objLabels), controllerLabels)

		template := tc.TomcatServerPodTemplateSpec()

		out.Spec.Template.ObjectMeta = template.ObjectMeta

		selector := metav1.SetAsLabelSelector(tc.TomcatServerPodLabels())
		if !reflect.DeepEqual(selector, out.Spec.Selector) {
			if out.ObjectMeta.CreationTimestamp.IsZero() {
				out.Spec.Selector = selector
			} else {
				return fmt.Errorf("deployment selector is immutable")
			}
		}

		err := mergo.Merge(&out.Spec.Template.Spec, template.Spec, mergo.WithTransformers(transformers.PodSpec))
		if err != nil {
			return err
		}

		if tc.Spec.Replicas != nil {
			out.Spec.Replicas = tc.Spec.Replicas
		}

		if out.Spec.Replicas == nil {
			out.Spec.Replicas = &oneReplica
		}

		return nil
	})
}
