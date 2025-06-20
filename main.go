package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/open-policy-agent/opa/v1/rego"
)

type CISFinding struct {
	CheckID  string                 `json:"check_id"`
	Resource string                 `json:"resource"`
	Evidence map[string]interface{} `json:"evidence"`
}

func loadFindingsFromDir(dir string) ([]CISFinding, error) {
	var results []CISFinding
	files, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		raw, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		var finding CISFinding
		if err := json.Unmarshal(raw, &finding); err != nil {
			return nil, fmt.Errorf("error in %s: %v", file, err)
		}
		results = append(results, finding)
	}
	return results, nil
}

func evaluatePolicy(input CISFinding, policyPath string) (bool, error) {
	ctx := context.Background()
	query := rego.New(
		rego.Query("data.cis.allow"),
		rego.Load([]string{policyPath}, nil),
		rego.Input(input),
	)
	rs, err := query.Eval(ctx)
	if err != nil || len(rs) == 0 {
		return false, err
	}
	ok, isBool := rs[0].Expressions[0].Value.(bool)
	return ok && isBool, nil
}

func generateReport(findings []CISFinding, policyPath string) {
	fmt.Println("\n--- AWS CIS Compliance Report ---")
	for _, f := range findings {
		pass, err := evaluatePolicy(f, policyPath)
		status := "NON-COMPLIANT"
		if err != nil {
			status = "ERROR"
		} else if pass {
			status = "COMPLIANT"
		}
		fmt.Printf("CheckID: %-6s | Resource: %-40s | Status: %s\n", f.CheckID, f.Resource, status)
	}
}

func main() {
	findings, err := loadFindingsFromDir("./steampipe-output")
	if err != nil {
		log.Fatalf("Error loading findings: %v", err)
	}
	generateReport(findings, "./policies/cis.rego")
}
