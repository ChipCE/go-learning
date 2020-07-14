package main

import (
	"log"
	"net/http"
	"time"
	"plugin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)


func main() {


	p1, err := plugin.Open("./plone/plone.so")
	if err != nil {
		log.Fatal(err)
	}

	ps1, err := p1.Lookup("Start")
	if err != nil {
		log.Fatal(err)
	}

	handle1, ok := ps1.(func(string) http.Handler)
	if !ok {
		log.Fatal("malformed plugin")
	}


	server01 := &http.Server{
		Addr:         ":8000",
		Handler:      handle1("lol"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	p2, err := plugin.Open("./pltwo/pltwo.so")
	if err != nil {
		log.Fatal(err)
	}

	ps2, err := p2.Lookup("Start")
	if err != nil {
		log.Fatal(err)
	}

	handle2, ok := ps2.(func(string) http.Handler)
	if !ok {
		log.Fatal("malformed plugin")
	}

	server02 := &http.Server{
		Addr:         ":8001",
		Handler:      handle2("lol"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server01.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		err := server02.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
