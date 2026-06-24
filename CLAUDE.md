# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repo purpose

Terraform provider for VNG Cloud, written in Go against `terraform-plugin-sdk/v2`. The provider exposes resources/data sources for VNG Cloud's vServer, VKS (managed Kubernetes), VLB (load balancer), and vDB (managed databases — relational, memstore, kafka) services.

## Build, install, run

- `go build -o terraform-provider-vngcloud` — compile the plugin binary.
- `./build.sh` — full local install loop: builds the binary and copies it into `~/.terraform.d/plugins/registry.terraform.nhontt/vngcloud/vngcloud/1.0.0/<os>_<arch>/`, then runs `terraform init`. Edit the path/version in `build.sh` if you need a different local source. Hardcoded to `linux_amd64`.
- `./run_mode.sh <terraform-args>` — `cd ./examples && terraform <args>`. Use to apply/plan against the example HCL.
- `./debug_run_mode.sh <terraform-args>` — same as above with `TF_LOG=debug` and `TF_LOG_PROVIDERS=debug`.
- Go toolchain pinned to **1.16.3** (`.go-version`, `go.mod`). Stay on this Go version unless a bump is intentional — newer Go can break `terraform-plugin-sdk/v2 v2.6.1`.
- Releases are produced by GoReleaser via `.github/workflows/release.yml` on `v*` tags (cross-compiles freebsd/windows/linux/darwin × amd64/386/arm/arm64). Don't run goreleaser locally without GPG setup.

There are no unit tests in this repo and no lint config; the only "test" loop is build → install → `terraform plan/apply` against `examples/`.

## Architecture

### Plugin entrypoint
`main.go` → `provider.Provider()` in `provider/provider.go`. That single function is the registry — it wires every resource and data source into Terraform's `ResourcesMap` / `DataSourcesMap` and declares the provider-level schema (`client_id`, `client_secret`, `token_url`, `*_base_url`). To add a new resource/data source, you must register it here in addition to implementing it.

`providerConfigure` builds a `*client.Client` via `client.NewClientV2(...)` and returns it as the meta value passed to every CRUD function — every resource starts with `cli := m.(*client.Client)`.

### Client layer (`client/`)
- `client/client.go` — top-level `Client` struct aggregating per-service API clients (`VserverClient`, `VlbClient`, `VksClient`, `Vdbv2Client`, plus legacy `VdbClient`) and the shared `AuthenClient`.
- `NewClientV2` is the path used today. `NewClient` is the older entrypoint kept around but not invoked by `provider.go`.
- `client/authen/` — OAuth2 client_credentials flow. The `*http.Client` it returns is shared across all service clients so every request automatically carries a fresh access token.
- `client/vserver/`, `client/vks/`, `client/vloadbalancing/`, `client/vdb/`, `client/vdbv2/` — **swagger-codegen output**. Do not hand-edit these files casually; they're regenerated from `api/swagger.yaml` (present inside e.g. `client/vdbv2/api/`). If the upstream API changes, regenerate rather than patch.
- `client/vdb` (legacy) and `client/vdbv2` (current) coexist. New work should target `vdbv2`. The legacy `vdb` client is still imported by `client.go` but is not wired into `NewClientV2`.

### Resource layer (`resource/<service>/`)
One package per service mirroring the client layout: `vserver`, `vks`, `vloadbalancing`, `vdbv2`. Each package contains:
- `resource_*.go` — `func ResourceX() *schema.Resource` returning the Terraform schema plus `Create/Read/Update/Delete` callbacks. CRUD bodies marshal Terraform state into the swagger client request structs and back.
- `data_source_*.go` — same shape but with only `Read`.
- `config.go` — package-scoped state-machine constants used with `helper/resource.StateChangeConf`. Pattern is `<thing><Action>Pending` / `<thing><Action>Target` / `Timeout` / `Delay` / `MinTimeout` (e.g. `databaseCreatePending`, `serverResizing`). When adding a long-running operation, follow this convention rather than inlining the strings.
- `util.go` — small helpers; notably `CheckErrorResponse(httpResponse)` + `GetResponseBody(httpResponse)` are the standard pattern for surfacing non-2xx errors back to Terraform. Many handlers ignore the `err` returned from the swagger client and rely on `CheckErrorResponse` instead — match the surrounding style.
- `resource/resource.go` is an empty placeholder; ignore it.

### State machines for async ops
Provisioning, resize, start/stop, and delete are all asynchronous on the VNG Cloud side. The convention everywhere in this repo is: kick off the API call, then poll status with `resource.StateChangeConf{ Pending, Target, Refresh, Timeout, Delay, MinTimeout }`. The string sets in `config.go` are the source of truth for valid transitions — if the backend introduces a new transient status, add it there rather than per-call.

## Conventions worth knowing

- **Resource naming**: Terraform resource keys are `vngcloud_<service>_<thing>` (e.g. `vngcloud_vserver_server`, `vngcloud_vdb_relational_database`, `vngcloud_vks_cluster`). Match this when registering new resources.
- **`m.(*client.Client)`** is the universal entrypoint inside CRUD funcs. The per-service API client is reached via `cli.Vdbv2Client.RelationalDatabaseAPIApi.X(...)`, etc.
- **Two vDB clients**: when touching database code, confirm whether you're in `client/vdb` (legacy) or `client/vdbv2` (current). Resources under `resource/vdbv2/` should use `cli.Vdbv2Client`; some files (`resource_relational_database.go`) import both packages — usually for shared model types — so don't blindly remove an import.
- **Docs in `docs/`** are user-facing Terraform Registry docs (`docs/resources/*.md`, `docs/data-sources/*.md`). When you add or change a resource schema, update the matching markdown.
- **`examples/main.tf` and root `variable.tf`** contain real-looking client IDs / project IDs / resource IDs. Treat them as test fixtures pointing at a dev tenant — don't introduce new secrets here.
