package k8s

type RequestOptions func(*K8sRequestBean) *K8sRequestBean

type IdentifierOptions func(*ResourceIdentifier) *ResourceIdentifier

func NewK8sRequestBean(opts ...RequestOptions) *K8sRequestBean {
	req := &K8sRequestBean{}
	for _, opt := range opts {
		if opt != nil {
			opt(req)
		}
	}
	return req
}

func (k8s *K8sRequestBean) WithResourceIdentifier(resourceIdentifier *ResourceIdentifier) RequestOptions {
	if resourceIdentifier == nil {
		resourceIdentifier = &ResourceIdentifier{}
	}
	return func(req *K8sRequestBean) *K8sRequestBean {
		req.ResourceIdentifier = *resourceIdentifier
		return req
	}
}

func NewResourceIdentifier(opts ...IdentifierOptions) *ResourceIdentifier {
	req := &ResourceIdentifier{}
	for _, opt := range opts {
		if opt != nil {
			opt(req)
		}
	}
	return req
}

func (r *ResourceIdentifier) WithName(name string) IdentifierOptions {
	return func(req *ResourceIdentifier) *ResourceIdentifier {
		req.Name = name
		return req
	}
}

func (r *ResourceIdentifier) WithNameSpace(namespace string) IdentifierOptions {
	return func(req *ResourceIdentifier) *ResourceIdentifier {
		req.Namespace = namespace
		return req
	}
}

func (r *ResourceIdentifier) WithGroup(group string) IdentifierOptions {
	return func(req *ResourceIdentifier) *ResourceIdentifier {
		req.GroupVersionKind.Group = group
		return req
	}
}

func (r *ResourceIdentifier) WithVersion(version string) IdentifierOptions {
	return func(req *ResourceIdentifier) *ResourceIdentifier {
		req.GroupVersionKind.Version = version
		return req
	}
}

func (r *ResourceIdentifier) WithKind(kind string) IdentifierOptions {
	return func(req *ResourceIdentifier) *ResourceIdentifier {
		req.GroupVersionKind.Kind = kind
		return req
	}
}
