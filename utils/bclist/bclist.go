// Copyright 2018 AT&T. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//https://github.com/urfave/cli
//https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/
//https://vorozhko.net/simple-example-of-golang-os-args-and-flag-module
//bclist jobs

//sum(up{}) by (job)

//bclist roles -labels job="haproxy"

//sum(up{job="haproxy"}) by (role)

//bclist datacenters

//sum(up{}) by (datacenter)

//bclist environments

//sum(up{}) by (env)

//bclist hosts -labels app="attebiz",env="production",datacenter="alpharetta",job="haproxy",role="int_traic"

//sum(up{app="attebiz",env="production",datacenter="alpharetta",job="haproxy", role="int_traic"}) by (host)

// Command bclist is a utily to access prometheus endpointi to get lists of applications, environments, datacenters, jobs, roles and hosts.
//
// Usage:
//
//    bclist [-labels <list>] <target>
//    bclist -help
//
// Options:
//    -help Show this screen
//    -labels comma separated list of labels like name1=value1,name2=value2
// Another line with description and maybe additional usage sample
//    hello $SOMEPARAM
package main

import (
	"flag"
	"fmt"
)

func main() {

	fmt.Printf("bcctl utiliy started\n")
	labelPtr := flag.String("labels", "", "Prometheus labels to be used inside a query")
	flag.Parse()

	fmt.Printf("labelPtr: %s\n", *labelPtr)
}
