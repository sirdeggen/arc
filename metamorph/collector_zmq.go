package metamorph

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type zmqCollector struct {
	zmqStats             *ZMQStats
	hashTx               *prometheus.Desc
	invalidTx            *prometheus.Desc
	discardedFromMempool *prometheus.Desc
}

// You must create a constructor for you prometheusCollector that
// initializes every descriptor and returns a pointer to the prometheusCollector
func newZMQCollector(url *url.URL, zmqStats *ZMQStats) *zmqCollector {

	uniqueName := strings.ReplaceAll(strings.ReplaceAll(url.Host, ".", "_"), ":", "_")

	c := &zmqCollector{
		zmqStats: zmqStats,
		hashTx: prometheus.NewDesc(fmt.Sprintf("arc_metamorph_zmq_hashtx_%s", uniqueName),
			"Shows the number of hashTx messages received",
			nil, nil,
		),
		invalidTx: prometheus.NewDesc(fmt.Sprintf("arc_metamorph_zmq_invalidtx_%s", uniqueName),
			"Shows the number of invalidTx messages received",
			nil, nil,
		),
		discardedFromMempool: prometheus.NewDesc(fmt.Sprintf("arc_metamorph_zmq_discardedfrommempool_%s", uniqueName),
			"Shows the number of discardedFromMempool messages received",
			nil, nil,
		),
	}

	prometheus.MustRegister(c)

	return c
}

// Describe writes all descriptors to the prometheus desc channel.
func (c *zmqCollector) Describe(ch chan<- *prometheus.Desc) {
	//Update this section with each metric you create for a given prometheusCollector
	ch <- c.hashTx
	ch <- c.invalidTx
	ch <- c.discardedFromMempool
}

// Collect implements required collect function for all prometheus collectors
func (c *zmqCollector) Collect(ch chan<- prometheus.Metric) {
	//Write the latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(c.hashTx, prometheus.CounterValue, float64(c.zmqStats.hashTx.Load()))
	ch <- prometheus.MustNewConstMetric(c.invalidTx, prometheus.CounterValue, float64(c.zmqStats.invalidTx.Load()))
	ch <- prometheus.MustNewConstMetric(c.discardedFromMempool, prometheus.CounterValue, float64(c.zmqStats.discardedFromMempool.Load()))
}
