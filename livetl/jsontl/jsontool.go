package jsontl

import "encoding/json"

// JSON returns json string of obj.
func JSON(obj interface{}) string {
	bs, err := json.Marshal(obj)
	if err != nil {
		return ""
	}

	return string(bs)
}

// PrettyJSON returns indented json string of obj.
func PrettyJSON(obj interface{}) string {
	bs, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return ""
	}

	return string(bs)
}
