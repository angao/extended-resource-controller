package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/caicloud/extended-resource-controller/pkg/controller"
	"github.com/caicloud/extended-resource-controller/pkg/util"
	"github.com/golang/glog"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	flag.Parse()

	kubeClient := util.CreateClientset()

	c := controller.NewExtendedResourceController(kubeClient)

	glog.V(2).Infof("Extended Resource ExtendedResourceController Run.")

	c.Run()
}
