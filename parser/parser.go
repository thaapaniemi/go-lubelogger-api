package parser

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/thaapaniemi/go-lubelogger-api/debuglog"
)

const FORMAT_USDATE = "01/02/2006"      // US date format
const FORMAT_ISO8601DATE = "2006-01-02" // ISO8601 date format

// UploadedFile represents an attachment in the LubeLogger API.
type UploadedFile struct {
	Name      string `json:"name"`
	Location  string `json:"location"`
	IsPending bool   `json:"isPending"`
}

// ConvertUploadedFiles safely converts file objects from the API response
// into UploadedFile structs, handling missing or malformed fields.
// Returns an error if the files array has unexpected structure or type issues.
func ConvertUploadedFiles(in interface{}) ([]UploadedFile, error) {
	if in == nil {
		return nil, nil
	}

	fileArray, ok := in.([]interface{})
	if !ok {
		return nil, fmt.Errorf("files field is not an array, got type: %T", in)
	}

	// Re-marshal to JSON and unmarshal into UploadedFile structs
	// This delegates type validation to the standard library
	jsonBytes, err := json.Marshal(fileArray)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal files for conversion: %v", err)
	}

	var files []UploadedFile
	err = json.Unmarshal(jsonBytes, &files)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal files: %v", err)
	}

	// Validate required fields: Name must not be empty
	for i, file := range files {
		if file.Name == "" {
			return nil, fmt.Errorf("files[%d] is missing required 'name' field", i)
		}
	}

	return files, nil
}

func ParseDateISO8601(dateStr interface{}) time.Time {
	if dateStr == nil || dateStr.(string) == "" {
		return time.Time{}
	}

	layout := FORMAT_ISO8601DATE
	t, err := time.Parse(layout, dateStr.(string))
	if err != nil {
		debuglog.Debugf("parsing of date %s failed: %s", dateStr, err.Error())
		return time.Time{}
	}
	return t
}

func ParseString(input interface{}) string {
	var out string
	if input != nil {
		out = input.(string)
	}

	return out
}

func ParseBool(input interface{}) bool {
	var out bool
	if input != nil {
		out = input.(bool)
	}

	return out
}

func ParseInt(input interface{}) int64 {
	var out int64

	if input != nil {
		n, err := input.(json.Number).Int64()
		if err == nil {
			out = n
		} else {
			debuglog.Debugf("parsing of int64 %s failed: %s", input, err.Error())
		}
	}

	return out
}

func ParseFloat(input interface{}) float64 {
	var out float64

	if input != nil {
		n, err := input.(json.Number).Float64()
		if err == nil {
			out = n
		} else {
			debuglog.Debugf("parsing of float %s failed: %s", input, err.Error())
		}
	}

	return out
}

func ParseStringSlice(input interface{}) []string {
	if input == nil {
		return nil
	}

	slice := make([]string, 0)
	for _, item := range input.([]interface{}) {
		slice = append(slice, item.(string))
	}

	return slice
}

func ParseIntSlice(input interface{}) []int64 {
	if input == nil {
		return nil
	}

	slice := make([]int64, 0)
	for _, item := range input.([]interface{}) {
		slice = append(slice, ParseInt(item))
	}

	return slice
}
