// Code generated by "stringer -type=Controller -linecomment"; DO NOT EDIT.

package controllers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Core-0]
	_ = x[JobService-1]
	_ = x[Portal-2]
	_ = x[Registry-3]
	_ = x[RegistryController-4]
	_ = x[ChartMuseum-5]
	_ = x[NotaryServer-6]
	_ = x[NotarySigner-7]
	_ = x[Clair-8]
	_ = x[Trivy-9]
	_ = x[Harbor-10]
}

const _Controller_name = "corejobserviceportalregistryregistryctlchartmuseumnotaryservernotarysignerclairtrivyharbor"

var _Controller_index = [...]uint8{0, 4, 14, 20, 28, 39, 50, 62, 74, 79, 84, 90}

func (i Controller) String() string {
	if i < 0 || i >= Controller(len(_Controller_index)-1) {
		return "Controller(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Controller_name[_Controller_index[i]:_Controller_index[i+1]]
}