package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// MyCollector - Define here all your collector
type MyCollector struct {
	myFirstMetric  *prometheus.Desc
	mySecondMetric *prometheus.Desc
}

// NewCollector - Constructor for our collector
func NewCollector() *MyCollector {
	return &MyCollector{
		myFirstMetric: prometheus.NewDesc("first_metric",
			"description of first_metric",
			nil, nil,
		),
		mySecondMetric: prometheus.NewDesc("second_metric",
			"description of second_metric",
			nil, nil,
		),
	}
}

//Describe --It essentially writes all descriptors to the prometheus desc channel
func (collector *MyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.myFirstMetric
	ch <- collector.mySecondMetric
}

//Collect --Collect implements required collect function for all promehteus collectors
func (collector *MyCollector) Collect(ch chan<- prometheus.Metric) {

	// implement the logic giving us the value, that we want to expose, here I just took an example
	// those value could be from kafka or simple database
	var firstValue float64
	var secondValue float64
	firstValue = 1
	secondValue = 2

	//Write latest value for each metric in the prometheus metric channel.
	ch <- prometheus.MustNewConstMetric(collector.myFirstMetric, prometheus.CounterValue, firstValue)
	ch <- prometheus.MustNewConstMetric(collector.mySecondMetric, prometheus.CounterValue, secondValue)
}
