// Copyright (c) 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transporter

// DoSendNotification ...
type DoSendNotification struct {
	ID         string `json:"id"`
	Recipients int    `json:"recipients"`
	Errors     interface{}
}

// DoSendNotificationTransporter ...
func DoSendNotificationTransporter() DoSendNotification {
	return DoSendNotification{}
}
