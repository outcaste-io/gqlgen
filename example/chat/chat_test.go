package chat

import (
	"testing"
	"time"

	"github.com/outcaste-io/gqlgen/client"
	"github.com/outcaste-io/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatSubscriptions(t *testing.T) {
	c := client.New(handler.NewDefaultServer(NewExecutableSchema(New())))

	sub := c.Websocket(`subscription @user(username:"outcaste-io") { messageAdded(roomName:"#gophers") { text createdBy } }`)
	defer sub.Close()

	go func() {
		var resp interface{}
		time.Sleep(10 * time.Millisecond)
		err := c.Post(`mutation { 
				a:post(text:"Hello!", roomName:"#gophers", username:"outcaste-io") { id } 
				b:post(text:"Hello Outcaste-Io!", roomName:"#gophers", username:"andrey") { id } 
				c:post(text:"Whats up?", roomName:"#gophers", username:"outcaste-io") { id } 
			}`, &resp)
		assert.NoError(t, err)
	}()

	var msg struct {
		resp struct {
			MessageAdded struct {
				Text      string
				CreatedBy string
			}
		}
		err error
	}

	msg.err = sub.Next(&msg.resp)
	require.NoError(t, msg.err, "sub.Next")
	require.Equal(t, "Hello!", msg.resp.MessageAdded.Text)
	require.Equal(t, "outcaste-io", msg.resp.MessageAdded.CreatedBy)

	msg.err = sub.Next(&msg.resp)
	require.NoError(t, msg.err, "sub.Next")
	require.Equal(t, "Whats up?", msg.resp.MessageAdded.Text)
	require.Equal(t, "outcaste-io", msg.resp.MessageAdded.CreatedBy)
}
