package handlers

import (
	"context"
	kappcs "github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned"
	"google.golang.org/appengine/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
)

type Apps struct {
	ClientSet kappcs.Interface
}

func (a *Apps) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()
	listOpts := metav1.ListOptions{}
	apps, err := a.ClientSet.KappctrlV1alpha1().Apps("").List(ctx, listOpts)
	if err != nil {
		log.Errorf(ctx, "error on apps: %s", err)
	}
	// write json output to writer
	json.NewEncoder(w).Encode(apps)
}
