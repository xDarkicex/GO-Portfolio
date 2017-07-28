package controllers

import "github.com/xDarkicex/PortfolioGo/helpers"

// Ws is type for websocket
type Ws helpers.Controller

// func (this Ws) DialSocket(a helpers.RouterArgs) {
// 	websocket.Handler(Socket).ServeHTTP(a.Response, a.Request)
// }

// func Socket(ws *websocket.Conn) {
// 	var msg string
// 	for {
// 		websocket.Message.Receive(ws, &msg)
// 		var data = make(map[string]interface{})
// 		json.Unmarshal([]byte(msg), &data)
// 		switch data["api"] {
// 		case "neuron":
// 			point := data["data"].(map[string]interface{})
// 			output := neuron.Ne.Process([]float64{point["x"].(float64), point["y"].(float64)})
// 			websocket.Message.Send(ws, `{
// 							"output": "`+strconv.FormatFloat(output, 'f', -1, 64)+`",
// 							"M": `+strconv.FormatFloat(neuron.M, 'f', -1, 64)+`,
// 							"B": `+strconv.FormatFloat(neuron.B, 'f', -1, 64)+`}`)
// 		default:

// 			io.Copy(ws, ws)
// 		}
// 	}
// }
