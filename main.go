package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/knabben/kapp-api-demo/handlers"
	"github.com/knabben/kapp-api-demo/k8s"
	"log"
	"net/http"
)

var (
	kubeConfig, addr string
)

func main() {
	flag.StringVar(&addr, "address", ":8000", "Set the webserver address")
	flag.StringVar(&kubeConfig, "kubeconfig", "", "Kubeconfig Path, default ot empty. (in-cluster)")
	flag.Parse()

	config, err := k8s.GetConfig(kubeConfig)
	if err != nil {
		log.Fatal(err)
	}

	cs, err := k8s.NewKubernetes(config).NewClientSet()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Handle("/apps", &handlers.Apps{ClientSet: cs})
	log.Printf("Service on %v", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
