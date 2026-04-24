# concourse-golang

A Go project with a multi-job Concourse CI pipeline.

## Pipeline

```
source-code (git)
       │
       ▼
    ┌──────┐
    │ lint │
    └──┬───┘
       │
       ▼
    ┌──────┐
    │ test │
    └──┬───┘
       │
   ┌───┴────┐
   ▼        ▼
┌───────┐ ┌────────────────┐
│ build │ │ security-scan  │
└───┬───┘ └───────┬────────┘
    │             │
    └──────┬──────┘
           ▼
       ┌────────┐
       │ deploy │
       └────────┘
```

## Jobs

| Job | Description |
|-----|-------------|
| lint | Runs golangci-lint |
| test | Runs go test with race detection and coverage |
| build | Compiles static binary |
| security-scan | Runs govulncheck for vulnerabilities |
| deploy | Deploys the app (requires build + security-scan to pass) |

## Local Development

```bash
go run .
# Server runs on http://localhost:8081

go test -v ./...
```

## Set Pipeline

```bash
fly -t ci set-pipeline -p concourse-golang -c ci/pipeline.yml
fly -t ci unpause-pipeline -p concourse-golang
```
