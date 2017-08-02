package testframework

import (
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/client-go/pkg/api/v1"
)

func (f *Framework) Namespace() string {
	return f.namespace
}

func (f *Framework) EnsureNamespace() error {
	_, err := f.KubeClient.CoreV1().Namespaces().Get(f.namespace, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		_, err := f.KubeClient.CoreV1().Namespaces().Create(&apiv1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: f.namespace,
			},
		})
		if err != nil {
			return err
		}
	}
	return err
}

func (f *Framework) DeleteNamespace() error {
	return f.KubeClient.CoreV1().Namespaces().Delete(f.namespace, &metav1.DeleteOptions{})
}
