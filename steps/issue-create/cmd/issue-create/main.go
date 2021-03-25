package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/puppetlabs/relay-sdk-go/pkg/log"
	"github.com/puppetlabs/relay-sdk-go/pkg/outputs"
	"github.com/puppetlabs/relay-sdk-go/pkg/taskutil"
	"github.com/relay-integrations/relay-jira-server/actions/steps/issue-create/pkg/issue"
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

	issue, err := issue.CreateIssue(spec)
	if err != nil {
		log.FatalE(err)
	}

	if issue != nil {
		log.Info(fmt.Sprintf("Created issue %v", issue.Key))
	}

	if client, err := outputs.NewDefaultOutputsClientFromNebulaEnv(); err != nil {
		log.FatalE(err)
	} else {
		if err := client.SetOutput(context.Background(), "issue", issue.Key); err != nil {
			log.FatalE(err)
		}
	}
}
