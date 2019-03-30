package scalogservice

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
	newOrderDeployment creates a Kubernetes Deployment
	used for managing the replication of the ordering
	layer
*/
func newOrderDeployment() *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "scalog-order-deployment",
			Namespace: "scalog",
			Labels: map[string]string{
				"app": "scalog-order",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: createInt32(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "scalog-order",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "scalog-order",
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "scalog-service-account",
					Containers: []corev1.Container{
						corev1.Container{
							Name:            "scalog-order-node",
							Image:           "scalog-order",
							ImagePullPolicy: "Never",
							Ports: []corev1.ContainerPort{
								corev1.ContainerPort{ContainerPort: 21024},
								corev1.ContainerPort{ContainerPort: 1337},
							},
							Env: []corev1.EnvVar{
								corev1.EnvVar{
									Name: "UID",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.uid",
										},
									},
								},
								corev1.EnvVar{
									Name:  "RAFT_CLUSTER_SIZE",
									Value: "2",
								},
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

/*
	newOrderServiceAccount creates a kubernetes Service Account
	used for binding RBACs and other abilities to specific
	scalog objects
*/
func newOrderServiceAccount() *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "scalog-order-service-account",
			Namespace: "scalog",
			Labels: map[string]string{
				"app": "scalog-order",
			},
		},
	}
}

/*
	newOrderService launches a new headless service for managing the network
	domain of statefulsets (data layer nodes).
*/
func newOrderService() *corev1.Service {
	labels := map[string]string{
		"name": "scalog-order-service",
	}
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "scalog-order-service",
			Namespace: "scalog",
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				corev1.ServicePort{
					Port:     21024,
					Protocol: "TCP",
				},
			},
			Selector: map[string]string{
				"app": "scalog-order",
			},
		},
	}
}
