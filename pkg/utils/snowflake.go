package utils

import (
	"github.com/bwmarrin/snowflake"
)

func GenerateSnowflake() (int64, error) {
	nodeNumber := 1
	node, err := snowflake.NewNode(int64(nodeNumber))
	if err != nil {
		return -1, err
	}
	id := node.Generate()
	return id.Int64(), nil
}
