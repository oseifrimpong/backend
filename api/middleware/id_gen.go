package middleware

import "github.com/bwmarrin/snowflake"

func GenerateSnowflake() (snowflake.ID, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}

	id := node.Generate()

	// newID := xid.ID.String()
	return id, nil
}
