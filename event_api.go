package slack

import (
	"encoding/json"
	"reflect"
	"io"
	"fmt"
)


type APIEventManager struct {
}

func NewAPIEventManager() *APIEventManager {
	return &APIEventManager{
	}
}

func (eapi *APIEventManager) ReceiveIncomingEvent(data io.ReadCloser, apiEvent * ApiEvent) error{
	rawEvent := json.RawMessage{}
	err := json.NewDecoder(data).Decode(&rawEvent)
	if err != nil {
		err = fmt.Errorf("JSON marshalling error")
		return err
	} else if len(rawEvent) == 0 {
		err = fmt.Errorf("Received empty event")
		return err
	}
	event := &Event{}
	err = json.Unmarshal(rawEvent, event)
	if err != nil {
		err = fmt.Errorf("JSON marshalling error")
		return err
	}
	v, exists := eventAPIMapping[event.Type]
	if !exists {
		err = fmt.Errorf("EAPI Error, received unmapped event %q: %s\n", event.Type, string(rawEvent))
		return err
	}
	t := reflect.TypeOf(v)
	recvEvent := reflect.New(t).Interface()
	err = json.Unmarshal(rawEvent, recvEvent)
	if err != nil {
		err = fmt.Errorf("EAPI Error, could not unmarshall event %q: %s\n", event.Type, string(rawEvent))
		return err
	}
	apiEvent.Type = event.Type
	apiEvent.Data = recvEvent
	return nil
}



// Subset of events supported by the Slack Events API
var eventAPIMapping = map[string]interface{}{

	"event_callback": EventCallbackEvent{},

	"url_verification" : URLVerificationEvent{},
}

// URLVerificationEvent represents the slack event for verifying the url with a challenge
type URLVerificationEvent struct {
	Type           string `json:"type"`
	Challenge string `json:"challenge"`
	Token    string `json:"token"`
}
