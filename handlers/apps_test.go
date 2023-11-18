package handlers

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"

	api "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/apps", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	cs := fake.NewSimpleClientset(
		[]runtime.Object{
			&api.App{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "simple-app",
					Namespace: "pkg-standalone",
				},
			},
		}...,
	)

	appHandler := Apps{ClientSet: cs}
	handler := http.HandlerFunc(appHandler.ServeHTTP)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if !strings.Contains(rr.Body.String(), "simple-app") {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}
