package slack

import (
	"encoding/json"
	"reflect"
	"io"
	"fmt"
)


type APIEventManager struct {
	Client
}

func NewAPIEventManager(api *Client) *APIEventManager {
	return &APIEventManager{
		Client:           *api,
	}
}

func (eapi *APIEventManager) ReceiveIncomingEvent(data io.ReadCloser, apiEvent * ApiEvent) error{
	rawEvent := json.RawMessage{}
	err := json.NewDecoder(data).Decode(&rawEvent)
	if err != nil {
		err = fmt.Errorf("JSON marshalling error")
		return err
	} else if len(rawEvent) == 0 {
		eapi.Debugln("Received empty event")
		err = fmt.Errorf("Received empty event")
		return err
	}
	eapi.Debugln("Incoming Event:", string(rawEvent[:]))
	event := &Event{}
	err = json.Unmarshal(rawEvent, event)
	if err != nil {
		err = fmt.Errorf("JSON marshalling error")
		return err
	}
	v, exists := eventAPIMapping[event.Type]
	if !exists {
		eapi.Debugf("EAPI Error, received unmapped event %q: %s\n", event.Type, string(event))
		err = fmt.Errorf("EAPI Error, received unmapped event %q: %s\n", event.Type, string(event))
		return err
	}
	t := reflect.TypeOf(v)
	recvEvent := reflect.New(t).Interface()
	err = json.Unmarshal(rawEvent, recvEvent)
	if err != nil {
		eapi.Debugf("EAPI Error, could not unmarshall event %q: %s\n", event.Type, string(event))
		err = fmt.Errorf("EAPI Error, could not unmarshall event %q: %s\n", event.Type, string(event))
		return err
	}
	apiEvent.Type = event.Type
	apiEvent.Data = recvEvent
	return nil
}



// Subset of events supported by the Slack Events API
var eventAPIMapping = map[string]interface{}{

	"channel_created":         ChannelCreatedEvent{},
	"channel_deleted":         ChannelDeletedEvent{},
	"channel_rename":          ChannelRenameEvent{},
	"channel_archive":         ChannelArchiveEvent{},
	"channel_unarchive":       ChannelUnarchiveEvent{},
	"channel_history_changed": ChannelHistoryChangedEvent{},

	"dnd_updated":      DNDUpdatedEvent{},
	"dnd_updated_user": DNDUpdatedEvent{},

	"im_created":         IMCreatedEvent{},
	"im_open":            IMOpenEvent{},
	"im_close":           IMCloseEvent{},
	"im_history_changed": IMHistoryChangedEvent{},

	"group_open":            GroupOpenEvent{},
	"group_close":           GroupCloseEvent{},
	"group_rename":          GroupRenameEvent{},
	"group_archive":         GroupArchiveEvent{},
	"group_unarchive":       GroupUnarchiveEvent{},
	"group_history_changed": GroupHistoryChangedEvent{},

	"file_created":         FileCreatedEvent{},
	"file_shared":          FileSharedEvent{},
	"file_unshared":        FileUnsharedEvent{},
	"file_public":          FilePublicEvent{},
	"file_private":         FilePrivateEvent{},
	"file_change":          FileChangeEvent{},
	"file_deleted":         FileDeletedEvent{},
	"file_comment_added":   FileCommentAddedEvent{},
	"file_comment_edited":  FileCommentEditedEvent{},
	"file_comment_deleted": FileCommentDeletedEvent{},

	"pin_added":   PinAddedEvent{},
	"pin_removed": PinRemovedEvent{},

	"star_added":   StarAddedEvent{},
	"star_removed": StarRemovedEvent{},

	"reaction_added":   ReactionAddedEvent{},
	"reaction_removed": ReactionRemovedEvent{},


	"team_join":              TeamJoinEvent{},
	"team_rename":            TeamRenameEvent{},
	"team_domain_change":     TeamDomainChangeEvent{},


	"user_change": UserChangeEvent{},

	"emoji_changed": EmojiChangedEvent{},

	"email_domain_changed": EmailDomainChangedEvent{},

	"url_verification" : URLVerificationEvent{},
}

// URLVerificationEvent represents the slack event for verifying the url with a challenge
type URLVerificationEvent struct {
	Type           string `json:"type"`
	Challenge string `json:"challenge"`
	Token    string `json:"token"`
}
