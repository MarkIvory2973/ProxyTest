package main

import (
	"ProxyTest/mathx"
	"ProxyTest/mihomo"
	"flag"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
)

func calcProxiesScore(x [][]float64, k float64) ([]float64, []float64, []float64) {
	var stdX []float64
	var medianX []float64
	for _, xi := range x {
		stdX = append(stdX, mathx.Std(xi))
		medianX = append(medianX, mathx.Median(xi))
	}

	scaledStdX := mathx.Norm(stdX)
	scaledMedianX := mathx.Norm(medianX)

	var scoreX []float64
	for idx, _ := range x {
		scoreXi := (1-k)*scaledStdX[idx] + k*scaledMedianX[idx]
		scoreXi = 1 - scoreXi

		scoreX = append(scoreX, scoreXi)
	}

	return scoreX, stdX, medianX
}

func main() {
	var Host string
	var Port int
	var TLS bool
	var groupName string
	var k float64
	flag.StringVar(&Host, "host", "127.0.0.1", "MiHoMo API host")
	flag.IntVar(&Port, "port", 9090, "MiHoMo API port")
	flag.BoolVar(&TLS, "tls", false, "Use TLS")
	flag.StringVar(&groupName, "group", "", "Group name")
	flag.Float64Var(&k, "weight", 0.5, "Weight (0~1)")
	flag.Parse()

	miHomo := mihomo.MiHoMo{TLS: TLS, Host: Host, Port: Port}

	proxiesName, err := miHomo.GetProxiesName(groupName)
	if err != nil {
		log.Fatal(err)
		return
	}

	var proxiesDelays [][]float64
	for _, proxyName := range proxiesName {
		proxyDelays, err := miHomo.GetProxyDelays(proxyName)
		if err != nil {
			log.Fatal(err)
			return
		}

		proxiesDelays = append(proxiesDelays, proxyDelays)
	}

	proxiesScore, proxiesStability, proxiesDelay := calcProxiesScore(proxiesDelays, k)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Score", "Stability", "Delay"})
	for _, index := range mathx.Argsort(proxiesScore, true) {
		table.Append([]string{proxiesName[index], fmt.Sprintf("%.2f", proxiesScore[index]), fmt.Sprintf("%.2f", proxiesStability[index]), fmt.Sprintf("%.2f", proxiesDelay[index])})
	}
	table.Render()
}
