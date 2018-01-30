package ingress

import (
	"github.com/appscode/kutil"
	core_util "github.com/appscode/kutil/core/v1"
	rbac_util "github.com/appscode/kutil/rbac/v1beta1"
	api "github.com/appscode/voyager/apis/voyager/v1beta1"
	"github.com/appscode/voyager/pkg/eventer"
	core "k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
	rbac "k8s.io/api/rbac/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *controller) reconcileRBAC() error {
	if vt, err := c.reconcileServiceAccount(); err != nil {
		c.recorder.Eventf(
			c.Ingress.ObjectReference(),
			core.EventTypeWarning,
			eventer.EventReasonIngressRBACFailed,
			"Reason: %s",
			err.Error(),
		)
		return err
	} else if vt != kutil.VerbUnchanged {
		c.recorder.Eventf(
			c.Ingress.ObjectReference(),
			core.EventTypeNormal,
			eventer.EventReasonIngressRBACSuccessful,
			"Successfully %s service account %s",
			vt,
			c.Ingress.OffshootName(),
		)
	}

	if vt, err := c.reconcileRoles(); err != nil {
		c.recorder.Eventf(
			c.Ingress.ObjectReference(),
			core.EventTypeWarning,
			eventer.EventReasonIngressRBACFailed,
			"Reason: %s",
			err.Error(),
		)
		return err
	} else if vt != kutil.VerbUnchanged {
		c.recorder.Eventf(
			c.Ingress.ObjectReference(),
			core.EventTypeNormal,
			eventer.EventReasonIngressRBACSuccessful,
			"Successfully %s role %s",
			vt,
			c.Ingress.OffshootName(),
		)
	}

	if vt, err := c.reconcileRoleBinding(); err != nil {
		c.recorder.Eventf(
			c.Ingress.ObjectReference(),
			core.EventTypeWarning,
			eventer.EventReasonIngressRBACFailed,
			"Reason: %s",
			err.Error(),
		)
		return err
	} else if vt != kutil.VerbUnchanged {
		c.recorder.Eventf(
			c.Ingress.ObjectReference(),
			core.EventTypeNormal,
			eventer.EventReasonIngressRBACSuccessful,
			"Successfully %s role binding %s",
			vt,
			c.Ingress.OffshootName(),
		)
	}
	return nil
}

func (c *controller) reconcileServiceAccount() (kutil.VerbType, error) {
	meta := metav1.ObjectMeta{
		Namespace: c.Ingress.Namespace,
		Name:      c.Ingress.OffshootName(),
	}
	_, vt, err := core_util.CreateOrPatchServiceAccount(c.KubeClient, meta, func(in *core.ServiceAccount) *core.ServiceAccount {
		in.ObjectMeta = c.ensureOwnerReference(in.ObjectMeta)

		if in.Annotations == nil {
			in.Annotations = map[string]string{}
		}
		in.Annotations[api.OriginAPISchema] = c.Ingress.APISchema()
		in.Annotations[api.OriginName] = c.Ingress.GetName()
		return in
	})
	return vt, err
}

func (c *controller) reconcileRoles() (kutil.VerbType, error) {
	meta := metav1.ObjectMeta{
		Namespace: c.Ingress.Namespace,
		Name:      c.Ingress.OffshootName(),
	}
	_, vt, err := rbac_util.CreateOrPatchRole(c.KubeClient, meta, func(in *rbac.Role) *rbac.Role {
		in.ObjectMeta = c.ensureOwnerReference(in.ObjectMeta)

		if in.Annotations == nil {
			in.Annotations = map[string]string{}
		}
		in.Annotations[api.OriginAPISchema] = c.Ingress.APISchema()
		in.Annotations[api.OriginName] = c.Ingress.GetName()

		in.Rules = []rbac.PolicyRule{
			{
				APIGroups: []string{core.GroupName},
				Resources: []string{"configmaps"},
				Verbs:     []string{"get", "list", "watch"},
			},
			// We need to have those permission for secret mounter
			{
				APIGroups: []string{core.GroupName},
				Resources: []string{"secrets"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{api.SchemeGroupVersion.Group},
				Resources: []string{"ingresses", "certificates"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{extensions.GroupName},
				Resources: []string{"ingresses"},
				Verbs:     []string{"get", "list", "watch"},
			},
		}
		return in
	})
	return vt, err
}

func (c *controller) reconcileRoleBinding() (kutil.VerbType, error) {
	meta := metav1.ObjectMeta{
		Namespace: c.Ingress.Namespace,
		Name:      c.Ingress.OffshootName(),
	}
	_, vt, err := rbac_util.CreateOrPatchRoleBinding(c.KubeClient, meta, func(in *rbac.RoleBinding) *rbac.RoleBinding {
		in.ObjectMeta = c.ensureOwnerReference(in.ObjectMeta)

		if in.Annotations == nil {
			in.Annotations = map[string]string{}
		}
		in.Annotations[api.OriginAPISchema] = c.Ingress.APISchema()
		in.Annotations[api.OriginName] = c.Ingress.GetName()

		in.RoleRef = rbac.RoleRef{
			APIGroup: rbac.GroupName,
			Kind:     "Role",
			Name:     c.Ingress.OffshootName(),
		}
		in.Subjects = []rbac.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      c.Ingress.OffshootName(),
				Namespace: c.Ingress.Namespace,
			},
		}
		return in
	})
	return vt, err
}

func (c *controller) ensureRoleBindingDeleted() error {
	c.logger.Infof("Deleting RoleBinding %s/%s", c.Ingress.Namespace, c.Ingress.OffshootName())
	return c.KubeClient.RbacV1beta1().
		RoleBindings(c.Ingress.Namespace).
		Delete(c.Ingress.OffshootName(), &metav1.DeleteOptions{})
}

func (c *controller) ensureRolesDeleted() error {
	c.logger.Infof("Deleting Role %s/%s", c.Ingress.Namespace, c.Ingress.OffshootName())
	return c.KubeClient.RbacV1beta1().
		Roles(c.Ingress.Namespace).
		Delete(c.Ingress.OffshootName(), &metav1.DeleteOptions{})
}

func (c *controller) ensureServiceAccountDeleted() error {
	c.logger.Infof("Deleting ServiceAccount %s/%s", c.Ingress.Namespace, c.Ingress.OffshootName())
	return c.KubeClient.CoreV1().
		ServiceAccounts(c.Ingress.Namespace).
		Delete(c.Ingress.OffshootName(), &metav1.DeleteOptions{})
}
