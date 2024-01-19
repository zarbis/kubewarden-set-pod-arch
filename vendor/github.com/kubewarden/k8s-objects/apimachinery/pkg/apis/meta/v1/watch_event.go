// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	jsonext "encoding/json"
)

// WatchEvent Event represents a single event to a watched resource.
//
// swagger:model WatchEvent
type WatchEvent struct {

	// Object is:
	//  * If Type is Added or Modified: the new state of the object.
	//  * If Type is Deleted: the state of the object immediately before deletion.
	//  * If Type is Error: *Status is recommended; other types may make sense
	//    depending on context.
	// Required: true
	Object jsonext.RawMessage `json:"object"`

	// type
	// Required: true
	Type *string `json:"type"`
}