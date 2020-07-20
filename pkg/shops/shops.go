package shops

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
)

func Shops(ctx context.Context) error {
	cs, err := commits(ctx)
	if err != nil {
		return err
	}
	sm, err := k10Jobs(ctx)
	if err != nil {
		return err
	}
	_ = sm

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 16, 8, 0, '\t', 0)

	fmt.Fprint(w, "\t")
	for _, c := range cs {
		fmt.Fprint(w, c[:8], "\t")
	}
	fmt.Fprintln(w)
	for n, s := range sm {
		fmt.Fprint(w, n, "\t")
		for _, c := range cs {
			stat := s[c]
			if stat == "" {
				stat = "---"
			}
			fmt.Fprint(w, stat, "\t")
		}
		fmt.Fprintln(w)
	}
	w.Flush()
	return nil
}

func printJSON(v interface{}) error {
	buf, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(buf))
	return nil
}

func stringSet(slice []string) map[string]struct{} {
	ss := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		ss[s] = struct{}{}
	}
	return ss
}
