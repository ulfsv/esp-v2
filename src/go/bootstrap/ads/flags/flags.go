// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flags

import (
	"flag"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/esp-v2/src/go/commonflags"
	"github.com/GoogleCloudPlatform/esp-v2/src/go/options"
	"github.com/golang/glog"
)

var (
	AdsConnectTimeout = flag.Duration("ads_connect_timeout", 10*time.Second, "ads connect timeout in seconds")
)

func DefaultBootstrapperOptionsFromFlags() options.AdsBootstrapperOptions {
	common_option := commonflags.DefaultCommonOptionsFromFlags()
	opts := options.AdsBootstrapperOptions{
		CommonOptions:     common_option,
		AdsConnectTimeout: *AdsConnectTimeout,
		DiscoveryAddress:  fmt.Sprintf("127.0.0.1:%d", common_option.DiscoveryPort),
	}

	glog.Infof("ADS Bootstrapper options: %+v", opts)
	return opts
}
