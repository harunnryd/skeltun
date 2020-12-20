// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package param

// DoSendNotification ...
type DoSendNotification struct {
	AppID                  string                 `json:"app_id"`
	IncludeExternalUserIDs []interface{}          `json:"include_external_user_ids"`
	Contents               map[string]interface{} `json:"contents"`
}

// DoSendNotificationParam ...
func DoSendNotificationParam() DoSendNotification {
	return DoSendNotification{}
}
