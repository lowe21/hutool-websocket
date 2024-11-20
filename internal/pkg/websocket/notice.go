package websocket

// Notice 通知
func Notice(message any, clientIds ...string) {
	if len(clientIds) > 0 {
		for _, clientId := range clientIds {
			if clientId != "" {
				go manager.GetClient(clientId).Send(message)
			}
		}
	} else {
		for _, client := range manager.GetClients() {
			go client.Send(message)
		}
	}
}
