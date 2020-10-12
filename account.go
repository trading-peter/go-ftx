package ftx

import (
	"log"

	"github.com/trading-peter/go-ftx/structs"
)

type Positions structs.Positions

func (client *FtxClient) GetPositions(showAvgPrice bool) (Positions, error) {
	var positions Positions
	resp, err := client._get("positions", []byte(""))
	if err != nil {
		log.Printf("Error GetPositions", err)
		return positions, err
	}
	err = _processResponse(resp, &positions)
	return positions, err
}
