package scalogservice

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
	newServiceAccount creates a kubernetes Service Account
	used for binding RBACs and other abilities to specific
	scalog objects
*/
func newServiceAccount() *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "scalog-data-service-account",
			Namespace: "scalog",
			Labels: map[string]string{
				"app": "scalog-data",
			},
		},
	}
}

/*
	newDataService launches a new headless service for managing the network
	domain of statefulsets (data layer nodes).
*/
func newDataService() *corev1.Service {
	labels := map[string]string{
		"name": "scalog-data-headless-service",
		"app":  "scalog-data",
	}

	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "scalog-data-headless-service",
			Namespace: "scalog",
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				corev1.ServicePort{
					Port: 21024,
				},
			},
			ClusterIP: "None", // Launch as a headless service
			Selector: map[string]string{
				"app": "scalog-data",
			},
		},
	}
}

/*
	newDataStatefulSet returns a StatefulSet. When created, the specified amount of
	replicas will eventually be created and assigned a sticky identity. We treat
	each statefulset as a "shard".
*/
func newDataStatefulSet(shardID string) *appsv1.StatefulSet {
	selector := metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app": "scalog-data",
		},
	}
	return &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "scalog-data-shard-" + shardID,
			Namespace: "scalog",
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: createInt32(int32(2)),
			Selector: &selector,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "scalog-data",
						"tier": "scalog-data-shard-" + shardID,
					},
				},
				Spec: corev1.PodSpec{
					TerminationGracePeriodSeconds: createInt64(int64(10)),
					ServiceAccountName:            "scalog-data-service-account",
					Volumes: []corev1.Volume{
						corev1.Volume{
							Name: "scalog-data-storage",
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: "scalog-stable-storage-claim",
								},
							},
						},
					},
					Containers: []corev1.Container{
						corev1.Container{
							Name:            "scalog-data-replica-" + shardID,
							Image:           "scalog-data",
							ImagePullPolicy: "Never",
							VolumeMounts: []corev1.VolumeMount{
								corev1.VolumeMount{
									Name:      "scalog-data-storage",
									MountPath: "/app/scalog-db",
								},
							},
							Ports: []corev1.ContainerPort{
								corev1.ContainerPort{
									ContainerPort: 21024,
									Name:          "scalog-data",
								},
							},
							Env: []corev1.EnvVar{
								corev1.EnvVar{
									Name:  "GRPC_GO_LOG_VERBOSITY_LEVEL",
									Value: "99",
								},
								corev1.EnvVar{
									Name:  "GRPC_GO_LOG_SEVERITY_LEVEL",
									Value: "info",
								},
								corev1.EnvVar{
									Name: "NODE_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "spec.nodeName",
										},
									},
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
		Status: appsv1.StatefulSetStatus{},
	}
}
