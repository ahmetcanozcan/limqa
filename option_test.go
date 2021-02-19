package limqa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestOption(t *testing.T) {
	tests := []struct {
		o Option
		t optType
		v bool
	}{
		{Durable(true),durable,true},
		{Durable(false),durable,false},
		{AutoDelete(true),autoDelete,true},
	}

	for _,v := range tests {
		assert.Equal(t,v.t,v.o.otype())
		assert.Equal(t,v.v,v.o.value())
	}
}


func TestParseOptions(t *testing.T) {
	defOpts :=[]Option{Durable(true),AutoDelete(false),NoWait(true)}
	opts := []Option{Durable(false),AutoDelete(true),Internal(true)}
	p := parseOptions(defOpts,opts)

	assert.Equal(t,p[durable].otype(),durable)
	assert.Equal(t,p[durable].value(),false)

	assert.Equal(t,p[autoDelete].otype(),autoDelete)
	assert.Equal(t,p[autoDelete].value(),true)

	assert.Equal(t,p[internal].otype(),internal)
	assert.Equal(t,p[internal].value(),true)

	assert.Equal(t,p[internal].otype(),internal)
	assert.Equal(t,p[internal].value(),true)

}



func TestGetFlags(t *testing.T) {
	m := map[optType]Option{
		durable : Durable(true),
		autoDelete : AutoDelete(false),
		noWait : NoWait(true),
	}

	f := getFlags(m)

	assert.Equal(t,f[durable],true)
	assert.Equal(t,f[autoDelete],false)
	assert.Equal(t,f[noWait],true)


}