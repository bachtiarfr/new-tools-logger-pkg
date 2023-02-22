package entity

import (
	"time"

	"github.com/google/uuid"
)

type TagLog string // internal, provider, activity

const (
	TagLogInternal = TagLog("internal")
	TagLogProvider = TagLog("provider")
	TagLogActivity = TagLog("activity")
)

type Logger struct {
	ID         		uuid.UUID       `json:"id"`          //uuid primarykey
	ServiceName		string	        `json:"service_name"`//varchar
	Request    		map[string]interface{} `json:"request"`     // json
	RequestVendor   map[string]interface{} `json:"request_vendor"`     // json
	Response   		map[string]interface{} `json:"response"`    // json
	ResponseVendor  map[string]interface{} `json:"response_vendor"`    // json
	URL        		string          `json:"url"`         // varchar
	UrlVendor       string          `json:"url_vendor"`         // varchar
	Method     		string          `json:"method"`      // varchar
	StatusCode 		int             `json:"status_code"` // int
	Event     		string 			`json:"event"`      // varchar
	RequestID 		string 			`json:"request_id"` // uuid
	Additional 		map[string]interface{} `json:"additional"` // json
	Error 			map[string]interface{} `json:"error"` //json
	Tag 			TagLog 			`json:"tag"`
	CreatedAt 		time.Time 		`json:"created_at"` // timestamp
}