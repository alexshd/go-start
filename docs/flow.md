

```go
func streamRoom(c *gin.Context) {
	roomid := c.Param("roomid")
	listener := openListener(roomid)
	ticker := time.NewTicker(1 * time.Second)

	users.Add("connected", 1)

	defer func() {
		closeListener(roomid, listener)
		ticker.Stop()
		users.Add("disconnected", 1)
	}()

	c.Stream(func(_ io.Writer) bool {
		select {
		case msg := <-listener:
			messages.Add("outbound", 1)
			c.SSEvent("message", msg)
		case <-ticker.C:
			c.SSEvent("stats", Stats())
		}

		return true
	})
}

```
