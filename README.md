# opa-cis-policies-aws

---

## 🧰 Requirements

- Go 1.20+
- JSON output from [Steampipe AWS plugin](https://hub.steampipe.io/plugins/turbot/aws)
- (Optional) OPA CLI for local policy testing

---

## 🚀 Usage

### 🔧 Run the CLI

```bash
go run main.go --dir=steampipe-output --policy=policies/cis.rego --output=text
🔘 CLI Flags
Flag	Description	Default
--dir	Directory containing JSON input files	./steampipe-output
--policy	Path to Rego policy	./policies/cis.rego
--output	Report format: text, json, or html	text

```

📊 Output Examples
🖥 Text Output
```yaml
--- AWS CIS Compliance Report ---
CheckID: 1.1   | Resource: aws_root_account       | Status: NON-COMPLIANT
CheckID: 1.2   | Resource: aws_iam_policy         | Status: COMPLIANT
```

📦 JSON Output
```json

[
  {
    "check_id": "1.1",
    "resource": "aws_root_account",
    "status": "NON-COMPLIANT"
  },
  {
    "check_id": "1.2",
    "resource": "aws_iam_policy",
    "status": "COMPLIANT"
  }
]
```
🌐 HTML Output
Generates a styled table saved to report.html.

📥 Sample Input
steampipe-output/1.1.json
```json

{
  "check_id": "1.1",
  "resource": "aws_account_root",
  "evidence": {
    "root_account_mfa_enabled": false
  }
}
```
🔐 Rego Policy Example
policies/cis.rego
```rego

package cis

default allow = false

allow {
  input.check_id == "1.1"
  input.evidence.root_account_mfa_enabled == true
}
```
🧪 GitHub Actions CI
.github/workflows/opa-check.yml
```yaml

name: Validate Rego Policies

on:
  push:
    paths:
      - '**.rego'
      - '**.go'
      - 'steampipe-output/**'

jobs:
  opa-eval:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'

      - name: Run Policy Evaluation
        run: go run main.go
```
✅ Supported CIS Controls
```
Check ID	Description
1.1	MFA enabled for root account
1.2	IAM password policy enforces complexity
2.1	CloudTrail enabled in all regions
3.1	S3 buckets block public access
4.1	SGs not open to world on port 22
```

🧩 Future Enhancements
📬 Slack or webhook integration

📂 Directory report archiving

🧾 PDF export for audit-ready reports

📈 UI Dashboard mode

📄 License
Apache-2.0 © 2024 [acemnto]

