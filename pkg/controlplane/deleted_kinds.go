/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controlplane

import (
	"os"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	apimachineryversion "k8s.io/apimachinery/pkg/version"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/klog/v2"
)

// resourceExpirationEvaluator holds info for deciding if a particular rest.Storage needs to excluded from the API
type resourceExpirationEvaluator struct {
	currentMajor int
	currentMinor int
	isAlpha      bool
	// This is usually set for testing for which tests need to be removed.  This prevent insta-failing CI.
	// Set KUBE_APISERVER_STRICT_REMOVED_API_HANDLING_IN_ALPHA to see what will be removed when we tag beta
	strictRemovedHandlingInAlpha bool
	// This is usually set by a cluster-admin looking for a short-term escape hatch after something bad happened.
	// This should be made a flag before merge
	// Set KUBE_APISERVER_SERVE_REMOVED_APIS_FOR_ONE_RELEASE to prevent removing APIs for one more release.
	serveRemoveAPIsOneMoreRelease bool
}

func newResourceExpirationEvaluator(currentVersion apimachineryversion.Info) (*resourceExpirationEvaluator, error) {
	ret := &resourceExpirationEvaluator{}
	if len(currentVersion.Major) > 0 {
		currentMajor64, err := strconv.ParseInt(currentVersion.Major, 10, 32)
		if err != nil {
			return nil, err
		}
		ret.currentMajor = int(currentMajor64)
	}
	if len(currentVersion.Minor) > 0 {
		// split the "normal" + and - for semver stuff
		minorString := strings.Split(currentVersion.Minor, "+")[0]
		minorString = strings.Split(minorString, "-")[0]
		minorString = strings.Split(minorString, ".")[0]
		currentMinor64, err := strconv.ParseInt(minorString, 10, 32)
		if err != nil {
			return nil, err
		}
		ret.currentMinor = int(currentMinor64)
	}

	ret.isAlpha = strings.Contains(currentVersion.GitVersion, "alpha")

	if envString, ok := os.LookupEnv("KUBE_APISERVER_STRICT_REMOVED_API_HANDLING_IN_ALPHA"); !ok {
		// do nothing
	} else if envBool, err := strconv.ParseBool(envString); err != nil {
		return nil, err
	} else {
		ret.strictRemovedHandlingInAlpha = envBool
	}
	if envString, ok := os.LookupEnv("KUBE_APISERVER_SERVE_REMOVED_APIS_FOR_ONE_RELEASE"); !ok {
		// do nothing
	} else if envBool, err := strconv.ParseBool(envString); err != nil {
		return nil, err
	} else {
		ret.serveRemoveAPIsOneMoreRelease = envBool
	}

	return ret, nil
}

func (e *resourceExpirationEvaluator) shouldServe(resourceServingInfo rest.Storage) bool {
	versionedPtr := resourceServingInfo.New()
	removed, ok := versionedPtr.(removedInterface)
	if !ok {
		return true
	}
	majorRemoved, minorRemoved := removed.APILifecycleRemoved()
	if e.currentMajor < majorRemoved {
		return true
	}
	if e.currentMajor > majorRemoved {
		return false
	}
	if e.currentMinor < minorRemoved {
		return true
	}
	if e.currentMinor > minorRemoved {
		return false
	}
	// at this point major and minor are equal, so this API should be removed when the current release GAs.
	// If this is an alpha tag, don't remove by default, but allow the option.
	// If the cluster-admin has requested serving one more release, allow it.
	if e.isAlpha && e.strictRemovedHandlingInAlpha { // don't serve in alpha if we want strict handling
		return false
	}
	if e.isAlpha { // alphas are allowed to continue serving expired betas while we clean up the test
		return true
	}
	if e.serveRemoveAPIsOneMoreRelease { // cluster-admins are allowed to kick the can one release down the road
		return true
	}
	return false
}

type removedInterface interface {
	APILifecycleRemoved() (major, minor int)
}

// removeDeletedKinds inspects the storage map and modifies it in place by removing storage for kinds that have been deleted.
// versionedResourcesStorageMap mirrors the field on APIGroupInfo, it's a map from version to resource to the storage.
func (e *resourceExpirationEvaluator) removeDeletedKinds(groupName string, versionedResourcesStorageMap map[string]map[string]rest.Storage) {
	versionsToRemove := sets.NewString()
	for apiVersion, versionToResource := range versionedResourcesStorageMap {
		resourcesToRemove := sets.NewString()
		for resourceName, resourceServingInfo := range versionToResource {
			if !e.shouldServe(resourceServingInfo) {
				resourcesToRemove.Insert(resourceName)
			}
		}

		for resourceName := range versionedResourcesStorageMap[apiVersion] {
			if !shouldRemoveResource(resourcesToRemove, resourceName) {
				continue
			}

			klog.V(1).Infof("Removing resource %v.%v.%v because it is time to stop serving it per APILifecycle.", resourceName, apiVersion, groupName)
			delete(versionedResourcesStorageMap[apiVersion], resourceName)
		}

		if len(versionedResourcesStorageMap[apiVersion]) == 0 {
			versionsToRemove.Insert(apiVersion)
		}
	}

	for _, apiVersion := range versionsToRemove.List() {
		klog.V(1).Infof("Removing version %v.%v because it is time to stop serving it because it has no resources per APILifecycle.", apiVersion, groupName)
		delete(versionedResourcesStorageMap, apiVersion)
	}
}

func shouldRemoveResource(resourcesToRemove sets.String, resourceName string) bool {
	for _, resourceToRemove := range resourcesToRemove.List() {
		if resourceName == resourceToRemove {
			return true
		}
		// our API works on nesting, so you can have deployments, deployments/status, and deployments/scale.  Not all subresources
		// server the parent type, but if the parent type (deployments in this case), has been removed, it's subresources should be removed too.
		if strings.HasPrefix(resourceName, resourceToRemove+"/") {
			return true
		}
	}
	return false
}
