package scalogservice

import (
	"strconv"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

/*
	newOrderDeployment creates a Kubernetes Deployment
	used for managing the replication of the ordering
	layer
*/
func newOrderDeployment(numOrderReplicas int, numDataReplicas int, batchInterval int) *appsv1.Deployment {
	numOrderReplica32 := int32(numOrderReplicas)
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "scalog-order-deployment",
			Namespace: "scalog",
			Labels: map[string]string{
				"app": "scalog-order",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &numOrderReplica32,
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
							Image:           "scalog/scalog:latest",
							Command:         []string{"./scalog"},
							Args:            []string{"order"},
							ImagePullPolicy: "Always",
							Ports: []corev1.ContainerPort{
								corev1.ContainerPort{ContainerPort: 21024},
								corev1.ContainerPort{ContainerPort: 10088},
								corev1.ContainerPort{
									Name:          "liveness-port",
									ContainerPort: 9305,
								},
							},
							LivenessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									TCPSocket: &corev1.TCPSocketAction{
										Port: intstr.FromString("liveness-port"),
									},
								},
								PeriodSeconds:       20,
								InitialDelaySeconds: 15,
								FailureThreshold:    5,
							},
							Env: []corev1.EnvVar{
								corev1.EnvVar{
									Name:  "BATCH_INTERVAL",
									Value: strconv.Itoa(batchInterval),
								},
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
									Value: strconv.Itoa(numOrderReplicas),
								},
								corev1.EnvVar{
									Name:  "REPLICA_COUNT",
									Value: strconv.Itoa(numDataReplicas),
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
					Name:     "grpclb",
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
