package slack


// ApiEvent is the main wrapper. You will find all the other messages attached
type ApiEvent struct {
	Type string
	Data interface{}
}

type EventCallbackEvent struct{

	Token       string                 `json:"token"`
	TeamId      string                 `json:"team_id"`
	ApiAppId    string                 `json:"api_app_id"`
	Event       map[string]interface{} `json:"event"`
	EventTS     string                 `json:"event_ts"`
	Type        string                 `json:"type"`
	AuthedUsers []string               `json:"authed_users"`
}



/*

General event response example

{
        "token": "z26uFbvR1xHJEdHE1OQiO6t8",
        "team_id": "T061EG9RZ",
        "api_app_id": "A0FFV41KK",
        "event": {
                "type": "reaction_added",
                "user": "U061F1EUR",
                "item": {
                        "type": "message",
                        "channel": "C061EG9SL",
                        "ts": "1464196127.000002"
                },
                "reaction": "slightly_smiling_face"
        },
        "event_ts": "1465244570.336841",
        "type": "event_callback",
        "authed_users": [
                "U061F7AUR"
        ]
}
*/




