/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package certs

const (
	clusterCAKey             = "cluster-ca-key"
	clusterCACertificate     = "cluster-ca-cert"
	etcdCAKey                = "etcd-ca-key"
	etcdCACertificate        = "etcd-ca-cert"
	frontProxyCAKey          = "front-proxy-ca-key"
	frontProxyCACertificate  = "front-proxy-ca-cert"
	serviceAccountPublicKey  = "service-account-public-key"
	serviceAccountPrivateKey = "service-account-private-key"
)

func CertificatesToMap(c *Certificates) map[string]string {
	return map[string]string{
		clusterCAKey:             string(c.ClusterCA.Key),
		clusterCACertificate:     string(c.ClusterCA.Cert),
		etcdCAKey:                string(c.EtcdCA.Key),
		etcdCACertificate:        string(c.EtcdCA.Cert),
		frontProxyCAKey:          string(c.FrontProxyCA.Key),
		frontProxyCACertificate:  string(c.FrontProxyCA.Cert),
		serviceAccountPrivateKey: string(c.ServiceAccount.Key),
		serviceAccountPublicKey:  string(c.ServiceAccount.Cert),
	}

}

func MapToCertificates(m map[string]string) *Certificates {
	certs := &Certificates{
		ClusterCA:      &KeyPair{},
		EtcdCA:         &KeyPair{},
		FrontProxyCA:   &KeyPair{},
		ServiceAccount: &KeyPair{},
	}

	if val, ok := m[clusterCAKey]; ok {
		certs.ClusterCA.Key = []byte(val)
	}
	if val, ok := m[clusterCACertificate]; ok {
		certs.ClusterCA.Cert = []byte(val)
	}
	if val, ok := m[etcdCAKey]; ok {
		certs.EtcdCA.Key = []byte(val)
	}
	if val, ok := m[etcdCACertificate]; ok {
		certs.EtcdCA.Cert = []byte(val)
	}
	if val, ok := m[frontProxyCAKey]; ok {
		certs.FrontProxyCA.Key = []byte(val)
	}
	if val, ok := m[frontProxyCACertificate]; ok {
		certs.FrontProxyCA.Cert = []byte(val)
	}
	if val, ok := m[serviceAccountPrivateKey]; ok {
		certs.ServiceAccount.Key = []byte(val)
	}
	if val, ok := m[serviceAccountPublicKey]; ok {
		certs.ServiceAccount.Cert = []byte(val)
	}

	return certs
}
