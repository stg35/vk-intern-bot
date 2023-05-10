package main

import (
	"fmt"
	"github.com/stg35/vk-intern-bot/client/vk"
	"log"
)

const (
	clientHost = "api.vk.com"
	token      = "vk1.a.iMjwOeBPKC-wI0f1R99fbdoyceNWJJ932dQn0Yb3H5-GpO5G4g2nLHW_u6DIClCuTVpwCmA_pPFD6lGpQSF-YMdGTRaz8PKwOTmWObwzfWD8IOFoi9TFGH0rQI1pXLo5HR8_v03g1e8E_-qgdAbFnoll8BiDrNJRBUDLKPjc7ReOjkpUZ10VYGeeO6mRtwVDLGUnfuPFLjFrwv8gVj3EMw"
	group_id   = 220390762
)

func main() {
	vkClient := vk.New(clientHost, token, group_id)
	data, err := vkClient.GetUpdates()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

// Client{
// 	host:     host,
// 	basePath: "method",
// 	token:    token,
// 	group_id: group_id,
// 	client:   http.Client{},
// }
