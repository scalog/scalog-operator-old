package scalogservice

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
	newDiscoveeryService launches a new service with an external IP for
	discoverying available scalog data servers
*/
func newDiscoveryService() *corev1.Service {
	labels := map[string]string{
		"name": "scalog-discovery-service",
	}
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "scalog-discovery-service",
			Namespace: "scalog",
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: "NodePort",
			Ports: []corev1.ServicePort{
				corev1.ServicePort{
					Port:     21024,
					Protocol: "TCP",
				},
			},
			Selector: map[string]string{
				"app": "scalog-discovery",
			},
		},
	}
}

/*
	newDiscoveryDeployment creates a Kubernetes Deployment used for managing
	the replication of the discovery service
*/
func newDiscoveryDeployment(numReplicas int32) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "scalog-discovery-deployment",
			Namespace: "scalog",
			Labels: map[string]string{
				"app": "scalog-discovery",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &numReplicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "scalog-discovery",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "scalog-discovery",
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "scalog-service-account",
					Containers: []corev1.Container{
						corev1.Container{
							Name:            "scalog-discovery-node",
							Image:           "scalog/scalog:latest",
							Command:         []string{"./scalog"},
							Args:            []string{"discovery"},
							ImagePullPolicy: "Always",
							Ports: []corev1.ContainerPort{
								corev1.ContainerPort{ContainerPort: 21024},
							},
							Env: []corev1.EnvVar{
								corev1.EnvVar{
									Name: "NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.name",
										},
									},
								},
								corev1.EnvVar{
									Name: "NAMESPACE",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.namespace",
										},
									},
								},
								corev1.EnvVar{
									Name: "POD_IP",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "status.podIP",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
