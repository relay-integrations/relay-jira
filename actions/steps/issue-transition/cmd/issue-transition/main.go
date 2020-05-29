package main

import (
	"flag"

	"github.com/puppetlabs/nebula-sdk/pkg/log"
	"github.com/puppetlabs/nebula-sdk/pkg/taskutil"
	"github.com/relay-integrations/relay-jira-server/actions/steps/issue-transition/pkg/issue"
)

func main() {
	u, err := taskutil.MetadataSpecURL()
	if err != nil {
		log.FatalE(err)
	}
	specURL := flag.String("spec-url", u, "url to fetch the spec from")

	flag.Parse()

	planOpts := taskutil.DefaultPlanOptions{SpecURL: *specURL}

	var spec issue.Spec
	if err := taskutil.PopulateSpecFromDefaultPlan(&spec, planOpts); err != nil {
		log.FatalE(err)
	}

	err = issue.TransitionIssue(spec)
	if err != nil {
		log.FatalE(err)
	}
}
