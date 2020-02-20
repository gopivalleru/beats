// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package osquery

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "osquery", asset.ModuleFieldsPri, AssetOsquery); err != nil {
		panic(err)
	}
}

// AssetOsquery returns asset data.
// This is the base64 encoded gzipped contents of module/osquery.
func AssetOsquery() string {
	return "eJyslMFu2zAMhu95ih8+p3mAHIYBAwrssh22nhtVYmwhEulJdBO//SDZSZM0GbCiOuRAkT//L6T8gB2Na0j+M1AaF4B6DbRG83OKNAvAUbbJ9+qF1/iyAIBHT8Fl0KGXpOTwMkI7wmbW2SCKGwItgG3NXNeqB7CJdN6tHB17WqNNMvRz5EbDS6FzsUR5CHoK39K7qzmdbxKj8NzgHdOkj0iavM2kq7Paa0/nvsrvxcXR2Y7GvSR3dfcPf+X87qgqQrbVVf0DoZ1RtMSUTHGsnc+gV+ILk2+WjC3yn2fqURI820SRWE2AM2qWiCbtMvYdaUepmiXWNGJvMoxzdN0EkIREUV7JrfBdYQ3jhSBcaZta0yzRzDnNshQ0mU2fO9HmNmsnWZ+9I1a/9ZQ+dxJvutjKxFj6QRj7ztuuRuYth2mJFT4jDcye29U7yR+SoglhPAkVhNtYA/vDs/o7qxWE2/+jeWJ/QNHLamJ/3K66Q0t4RiYr7DKyZ0vTXS+2W+Epk6vwVmI/qOd2+gJ8PYltYCUMkW+DWBOInUn3YT40nV+aipVEfaJctrIkHrGshED1DVTkJUwuCNHo/N7nka0WfwMAAP//gwp0AA=="
}
