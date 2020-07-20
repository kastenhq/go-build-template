package shops

import (
	"context"

	"github.com/thockin/go-build-template/pkg/ship"
)

type statusMap map[string]map[string]string

func k10Jobs(ctx context.Context) (statusMap, error) {
	sc, err := ship.NewClient()
	if err != nil {
		return nil, err
	}
	bjs, _, err := sc.BuildJobs(ctx)
	if err != nil {
		return nil, err
	}
	jobSet := map[string]struct{}{}
	sm := statusMap{}
	for _, bj := range bjs {
		kj, ok := k10JobFromBuildJob(bj)
		if !ok {
			continue
		}
		if _, ok = sm[kj.Name]; !ok {
			sm[kj.Name] = make(map[string]string)
		}
		sm[kj.Name][kj.Commit] = kj.Status
		jobSet[kj.Name] = struct{}{}
	}
	return sm, nil
}

type K10Job struct {
	Name   string
	Commit string
	Status string
}

var statusCodes = map[int]string{
	4000: "Queued",
	4001: "Processing",
	4002: "Success",
	4003: "Failure",
	4004: "Error",
	4005: "Waiting",
	4006: "Cancelled",
	4007: "Unstable",
}

func k10JobFromBuildJob(bj ship.BuildJob) (*K10Job, bool) {
	for _, d := range bj.PropertyBag.Payload.Dependencies {
		if d.Name != "k10_repo" {
			continue
		}
		return &K10Job{
			Name:   bj.PropertyBag.Payload.Name,
			Commit: d.Version.VersionName,
			Status: statusCodes[bj.StatusCode],
		}, true
	}
	return nil, false

}
