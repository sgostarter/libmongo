package libmongo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMongo(t *testing.T) {
	cli, opts, err := InitMongo("mongodb://mongo_default_user:mongo_default_pass@127.0.0.1:8309/my_db")
	assert.Nil(t, err)
	assert.NotNil(t, cli)
	assert.EqualValues(t, "my_db", opts.Auth.AuthSource)

	err = cli.Ping(context.TODO(), nil)
	assert.Nil(t, err)
}
