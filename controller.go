package main

import (
	"fmt"
	"log"

	"github.com/zdnscloud/gok8s/cache"
	"github.com/zdnscloud/gok8s/client"
	"github.com/zdnscloud/gok8s/client/config"
	"github.com/zdnscloud/gok8s/controller"
	"github.com/zdnscloud/gok8s/event"
	"github.com/zdnscloud/gok8s/handler"
	"github.com/zdnscloud/gok8s/predicate"

	appv1beta1 "github.com/zdnscloud/my-app-operator/pkg/apis/app/v1beta1"
)

type dumbEventHandler struct {
}

func (d *dumbEventHandler) OnCreate(e event.CreateEvent) (handler.Result, error) {
	app := e.Object.(*appv1beta1.Application)
	log.Printf("create application [%v] \n", app)
	return handler.Result{}, nil
}

func (d *dumbEventHandler) OnUpdate(e event.UpdateEvent) (handler.Result, error) {
	oldApp := e.ObjectOld.(*appv1beta1.Application)
	newApp := e.ObjectNew.(*appv1beta1.Application)
	log.Printf("update application from [%v] to [%v]\n", oldApp, newApp)
	return handler.Result{}, nil
}

func (d *dumbEventHandler) OnDelete(e event.DeleteEvent) (handler.Result, error) {
	app := e.Object.(*appv1beta1.Application)
	log.Printf("delete application [%v]\n", app)
	return handler.Result{}, nil
}

func (d *dumbEventHandler) OnGeneric(e event.GenericEvent) (handler.Result, error) {
	return handler.Result{}, nil
}

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Panic(fmt.Sprintf("get config failed:%v\n", err))
	}

	stop := make(chan struct{})
	defer close(stop)

	c, err := cache.New(cfg, cache.Options{})
	if err != nil {
		log.Panic(fmt.Sprintf("create cache failed %v\n", err))
	}
	go c.Start(stop)

	c.WaitForCacheSync(stop)

	scheme := client.GetDefaultScheme()
	appv1beta1.AddToScheme(scheme)
	ctrl := controller.New("dumbController", c, scheme)
	ctrl.Watch(&appv1beta1.Application{})
	handler := &dumbEventHandler{}
	ctrl.Start(stop, handler, predicate.NewIgnoreUnchangedUpdate())
}
