/*
 * Copyright 2014 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Region defnintion for BCE
package common

type RegionType struct {
    regionIds []string
}

func (region RegionType) String() string {
    if len(region.regionIds) > 0 {
        return region.regionIds[0]
    }
    return ""
}

func NewRegion(regionIds...string) (*RegionType) {
    var ids []string
    for _, id := range regionIds {
        ids = append(ids, id)
    }
    return &RegionType{ids}
}

var BEIJING = NewRegion("bj")

