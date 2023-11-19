package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/knabben/kapp-api-demo/handlers"
	"github.com/knabben/kapp-api-demo/k8s"
	kappcs "github.com/vmware-tanzu/carvel-kapp-controller/pkg/client/clientset/versioned"
	"log"
	"net/http"
)

var (
	kubeConfig, addr string
	// staticPath points to the frontend build folder
	staticPath = "attestation/build"
)

func init() {
	flag.StringVar(&addr, "address", ":8000", "Set the webserver address")
	flag.StringVar(&kubeConfig, "kubeconfig", "", "Kubeconfig Path, default ot empty. (in-cluster)")
}

func main() {
	flag.Parse()

	config, err := k8s.GetConfig(kubeConfig)
	if err != nil {
		log.Fatal(err)
	}

	cs, err := k8s.NewKubernetes(config).NewClientSet()
	if err != nil {
		log.Fatal(err)
	}

	runServer(cs)
}

func runServer(cs *kappcs.Clientset) {
	r := mux.NewRouter()

	r.Handle("/api/apps", &handlers.Apps{ClientSet: cs})

	// Static handlers
	spa := handlers.SpaHandler{StaticPath: staticPath, IndexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	log.Printf("Service on %v", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
