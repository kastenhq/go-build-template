package ship

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/PeoplePerHour/go-shippable"
)

type Client struct {
	cli *shippable.Client
}

func NewClient() (*Client, error) {
	const shippableTokenEnvVar = `SHIPPABLE_TOKEN`
	token, ok := os.LookupEnv(shippableTokenEnvVar)
	if !ok {
		return nil, fmt.Errorf("please set $%s", shippableTokenEnvVar)
	}
	return &Client{
		cli: shippable.NewClient(token),
	}, nil
}

func (c *Client) Jobs(ctx context.Context, runsIDs ...string) (jobs []Job, resp *shippable.Response, err error) {
	q := url.Values{}
	if len(runsIDs) > 0 {
		q.Add("runIds", strings.Join(runsIDs, ","))
	}
	u := url.URL{
		Path:     "resources",
		RawQuery: q.Encode(),
	}
	req, err := c.cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	jobs = []Job{}
	resp, err = c.cli.Do(req, &jobs)

	return
}

func (c *Client) Runs(ctx context.Context) (runs []Run, resp *shippable.Response, err error) {
	const url = "/runs"
	req, err := c.cli.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	runs = []Run{}
	resp, err = c.cli.Do(req, &runs)
	return
}

func (c *Client) Builds(ctx context.Context) (builds []Build, resp *shippable.Response, err error) {
	const url = "/builds"
	req, err := c.cli.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	builds = []Build{}
	resp, err = c.cli.Do(req, &builds)
	return
}

func (c *Client) Resources(ctx context.Context) (resources []Resource, resp *shippable.Response, err error) {
	q := url.Values{}
	q.Add("isJob", "true")
	u := url.URL{
		Path:     "resources",
		RawQuery: q.Encode(),
	}
	req, err := c.cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	resources = make([]Resource, 100)
	resp, err = c.cli.Do(req, &resources)
	fmt.Printf("%#v\n", resp)
	return
}

func (c *Client) BuildJobs(ctx context.Context) (buildJobs []BuildJob, resp *shippable.Response, err error) {
	q := url.Values{}
	u := url.URL{
		Path:     "buildJobs",
		RawQuery: q.Encode(),
	}
	req, err := c.cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	buildJobs = []BuildJob{}
	resp, err = c.cli.Do(req, &buildJobs)
	return
}
