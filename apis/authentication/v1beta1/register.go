package v1beta1

import "github.com/onedata/k8s"

func init() {
	k8s.Register("authentication.k8s.io", "v1beta1", "tokenreviews", false, &TokenReview{})
}
