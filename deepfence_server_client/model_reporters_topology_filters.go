/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deepfence_server_client

import (
	"encoding/json"
)

// checks if the ReportersTopologyFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportersTopologyFilters{}

// ReportersTopologyFilters struct for ReportersTopologyFilters
type ReportersTopologyFilters struct {
	CloudFilter []string `json:"cloud_filter"`
	HostFilter []string `json:"host_filter"`
	RegionFilter []string `json:"region_filter"`
}

// NewReportersTopologyFilters instantiates a new ReportersTopologyFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportersTopologyFilters(cloudFilter []string, hostFilter []string, regionFilter []string) *ReportersTopologyFilters {
	this := ReportersTopologyFilters{}
	this.CloudFilter = cloudFilter
	this.HostFilter = hostFilter
	this.RegionFilter = regionFilter
	return &this
}

// NewReportersTopologyFiltersWithDefaults instantiates a new ReportersTopologyFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportersTopologyFiltersWithDefaults() *ReportersTopologyFilters {
	this := ReportersTopologyFilters{}
	return &this
}

// GetCloudFilter returns the CloudFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ReportersTopologyFilters) GetCloudFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CloudFilter
}

// GetCloudFilterOk returns a tuple with the CloudFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersTopologyFilters) GetCloudFilterOk() ([]string, bool) {
	if o == nil || isNil(o.CloudFilter) {
		return nil, false
	}
	return o.CloudFilter, true
}

// SetCloudFilter sets field value
func (o *ReportersTopologyFilters) SetCloudFilter(v []string) {
	o.CloudFilter = v
}

// GetHostFilter returns the HostFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ReportersTopologyFilters) GetHostFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.HostFilter
}

// GetHostFilterOk returns a tuple with the HostFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersTopologyFilters) GetHostFilterOk() ([]string, bool) {
	if o == nil || isNil(o.HostFilter) {
		return nil, false
	}
	return o.HostFilter, true
}

// SetHostFilter sets field value
func (o *ReportersTopologyFilters) SetHostFilter(v []string) {
	o.HostFilter = v
}

// GetRegionFilter returns the RegionFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ReportersTopologyFilters) GetRegionFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RegionFilter
}

// GetRegionFilterOk returns a tuple with the RegionFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersTopologyFilters) GetRegionFilterOk() ([]string, bool) {
	if o == nil || isNil(o.RegionFilter) {
		return nil, false
	}
	return o.RegionFilter, true
}

// SetRegionFilter sets field value
func (o *ReportersTopologyFilters) SetRegionFilter(v []string) {
	o.RegionFilter = v
}

func (o ReportersTopologyFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportersTopologyFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudFilter != nil {
		toSerialize["cloud_filter"] = o.CloudFilter
	}
	if o.HostFilter != nil {
		toSerialize["host_filter"] = o.HostFilter
	}
	if o.RegionFilter != nil {
		toSerialize["region_filter"] = o.RegionFilter
	}
	return toSerialize, nil
}

type NullableReportersTopologyFilters struct {
	value *ReportersTopologyFilters
	isSet bool
}

func (v NullableReportersTopologyFilters) Get() *ReportersTopologyFilters {
	return v.value
}

func (v *NullableReportersTopologyFilters) Set(val *ReportersTopologyFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableReportersTopologyFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableReportersTopologyFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportersTopologyFilters(val *ReportersTopologyFilters) *NullableReportersTopologyFilters {
	return &NullableReportersTopologyFilters{value: val, isSet: true}
}

func (v NullableReportersTopologyFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportersTopologyFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


