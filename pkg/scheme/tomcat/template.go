package tomcat

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	// TomcatHTTPPort is the default open port of tomcat container
	TomcatHTTPPort = 8080
)

// TomcatServerPodTemplateSpec generates a pod template spec suitable for use in Tomcat deployment
func (tomcat *Tomcat) TomcatServerPodTemplateSpec() (out corev1.PodTemplateSpec) {
	out = corev1.PodTemplateSpec{}
	out.ObjectMeta.Labels = tomcat.TomcatServerPodLabels()

	out.Spec.Containers = []corev1.Container{
		{
			Name:  "tomcat",
			Image: tomcat.Spec.Image,
			Ports: []corev1.ContainerPort{
				{
					Name:          "http",
					ContainerPort: int32(TomcatHTTPPort),
				},
			},
		},
	}

	return out
}
