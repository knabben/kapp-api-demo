package k8s

import (
	kappcs "github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Kubernetes struct {
	Config *rest.Config
}

func NewKubernetes(config *rest.Config) *Kubernetes {
	return &Kubernetes{Config: config}
}

func (k *Kubernetes) NewClientSet() (*kappcs.Clientset, error) {
	return kappcs.NewForConfig(k.Config)
}

func GetConfig(kubeconfig string) (config *rest.Config, err error) {
	if kubeconfig == "" {
		config, err = rest.InClusterConfig()
	}
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	return
}
