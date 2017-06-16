package fake

import (
	aci "github.com/appscode/voyager/api"
	"github.com/appscode/voyager/client/clientset"
apiv1 "k8s.io/client-go/pkg/api/v1"
metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
"k8s.io/client-go/testing"
"k8s.io/apimachinery/pkg/labels"
"k8s.io/apimachinery/pkg/watch"
)

type FakeIngress struct {
	Fake *testing.Fake
	ns   string
}

var ingressResource = metav1.GroupVersionResource{Group: "appscode.com", Version: "v1beta1", Resource: "ingresses"}

var _ clientset.IngressInterface = &FakeIngress{}

// Get returns the Ingress by name.
func (mock *FakeIngress) Get(name string) (*aci.Ingress, error) {
	obj, err := mock.Fake.
		Invokes(testing.NewGetAction(ingressResource, mock.ns, name), &aci.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aci.Ingress), err
}

// List returns the a of Ingresss.
func (mock *FakeIngress) List(opts apiv1.ListOptions) (*aci.IngressList, error) {
	obj, err := mock.Fake.
		Invokes(testing.NewListAction(ingressResource, mock.ns, opts), &aci.Ingress{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aci.IngressList{}
	for _, item := range obj.(*aci.IngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Create creates a new Ingress.
func (mock *FakeIngress) Create(svc *aci.Ingress) (*aci.Ingress, error) {
	obj, err := mock.Fake.
		Invokes(testing.NewCreateAction(ingressResource, mock.ns, svc), &aci.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aci.Ingress), err
}

// Update updates a Ingress.
func (mock *FakeIngress) Update(svc *aci.Ingress) (*aci.Ingress, error) {
	obj, err := mock.Fake.
		Invokes(testing.NewUpdateAction(ingressResource, mock.ns, svc), &aci.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aci.Ingress), err
}

// Delete deletes a Ingress by name.
func (mock *FakeIngress) Delete(name string) error {
	_, err := mock.Fake.
		Invokes(testing.NewDeleteAction(ingressResource, mock.ns, name), &aci.Ingress{})

	return err
}

func (mock *FakeIngress) UpdateStatus(srv *aci.Ingress) (*aci.Ingress, error) {
	obj, err := mock.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ingressResource, "status", mock.ns, srv), &aci.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aci.Ingress), err
}

func (mock *FakeIngress) Watch(opts apiv1.ListOptions) (watch.Interface, error) {
	return mock.Fake.
		InvokesWatch(testing.NewWatchAction(ingressResource, mock.ns, opts))
}
