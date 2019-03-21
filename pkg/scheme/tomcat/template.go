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

	out.Spec.Volumes = []corev1.Volume{
		{
			Name: "app-volume",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		},
	}

	out.Spec.InitContainers = []corev1.Container{
		{
			Name:  "war",
			Image: tomcat.Spec.WebArchiveImage,
			Command: []string{
				"sh",
				"-c",
				"cp /*.war /app",
			},
			VolumeMounts: []corev1.VolumeMount{
				{
					Name:      "app-volume",
					MountPath: "/app",
				},
			},
		},
	}

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
			VolumeMounts: []corev1.VolumeMount{
				{
					Name:      "app-volume",
					MountPath: tomcat.Spec.DeployDirectory,
				},
			},
		},
	}

	return out
}
