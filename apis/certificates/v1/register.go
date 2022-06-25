

package v1

import "github.com/onedata/k8s"

func init() {
	k8s.Register("certificates.k8s.io", "v1", "certificatesigningrequests", false, &CertificateSigningRequest{})

	k8s.RegisterList("certificates.k8s.io", "v1", "certificatesigningrequests", false, &CertificateSigningRequestList{})
}