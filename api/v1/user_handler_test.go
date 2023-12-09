package rest

// func TestHandler_AddUser(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	mockdbLayer := dblayer.NewMockDBLayerWithData()
// 	h := NewHandlerWithDB(mockdbLayer)
// 	const addUserURL string = "/"
// 	type errMSG struct {
// 		Error string `json:"error"`
// 	}

// 	tests := []struct {
// 		name             string
// 		inErr            error
// 		outStatusCode    int
// 		expectedRespBody interface{}
// 	}{
// 		{
// 			"getproductsnoerrors",
// 			nil,
// 			http.StatusOK,
// 			mockdbLayer.AddUser(),
// 		},
// 		{
// 			"getproductswitherror",
// 			errors.New("get products error"),
// 			http.StatusInternalServerError,
// 			errMSG{Error: "get products error"},
// 		},
// 	}
// }
