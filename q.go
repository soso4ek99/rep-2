package main

/*type Message struct {
	Text string `json:"text"`
}
type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var massages []Message

func gethandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &massages)
}
func posthandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  "error",
			Message: "не смогли добавить сообщение",
		})
	}
	massages = append(massages, message)
	return c.JSON(http.StatusOK, response{
		Status:  "success",
		Message: "письмо успешно добавленно",
	})
}
e := echo.New()
e.GET("/mes", gethandler)
e.POST("/mes", posthandler)
e.Start(":8080")*/
