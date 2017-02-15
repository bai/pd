// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import "github.com/prometheus/client_golang/prometheus"

var (
	cmdCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "pd",
			Subsystem: "cmd",
			Name:      "cmds_total",
			Help:      "Counter of cmds.",
		}, []string{"type"})

	cmdFailedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "pd",
			Subsystem: "cmd",
			Name:      "cmds_failed_total",
			Help:      "Counter of failed cmds.",
		}, []string{"type"})

	cmdDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "pd",
			Subsystem: "cmd",
			Name:      "handle_cmds_duration_seconds",
			Help:      "Bucketed histogram of processing time (s) of handled success cmds.",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2, 13),
		}, []string{"type"})

	cmdFailedDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "pd",
			Subsystem: "cmd",
			Name:      "handle_failed_cmds_duration_seconds",
			Help:      "Bucketed histogram of processing time (s) of failed handled cmds.",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2, 13),
		}, []string{"type"})

	cmdCompletedDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "pd",
			Subsystem: "cmd",
			Name:      "handle_completed_cmds_duration_seconds",
			Help:      "Bucketed histogram of processing time (s) of completed cmds.",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2, 13),
		}, []string{"type"})

	txnCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "pd",
			Subsystem: "txn",
			Name:      "txns_count",
			Help:      "Counter of txns.",
		}, []string{"result"})

	txnDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "pd",
			Subsystem: "txn",
			Name:      "handle_txns_duration_seconds",
			Help:      "Bucketed histogram of processing time (s) of handled txns.",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2, 13),
		}, []string{"result"})

	operatorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "pd",
			Subsystem: "schedule",
			Name:      "operators_count",
			Help:      "Counter of schedule operators.",
		}, []string{"type"})

	clusterStatusGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "pd",
			Subsystem: "cluster",
			Name:      "status",
			Help:      "Status of the cluster.",
		}, []string{"type"})

	timeJumpBackCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "pd",
			Subsystem: "monitor",
			Name:      "time_jump_back_total",
			Help:      "Counter of system time jumps backward.",
		})
)

func init() {
	prometheus.MustRegister(cmdCounter)
	prometheus.MustRegister(cmdFailedCounter)
	prometheus.MustRegister(cmdDuration)
	prometheus.MustRegister(cmdFailedDuration)
	prometheus.MustRegister(cmdCompletedDuration)
	prometheus.MustRegister(txnCounter)
	prometheus.MustRegister(txnDuration)
	prometheus.MustRegister(operatorCounter)
	prometheus.MustRegister(clusterStatusGauge)
	prometheus.MustRegister(timeJumpBackCounter)
}
