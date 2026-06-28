# Safety

AgentOS includes multiple safety mechanisms to prevent accidental damage.

## Command Denylist

Shell commands are checked against a denylist before execution. Denied by default:
- `rm -rf`, `rm -rf /`, `rm -rf /*`
- `sudo`
- `docker run --privileged`
- `curl`, `wget`
- `ssh`, `scp`

Additional deny patterns can be added in the profile YAML under `tools.deny_commands`.

## Secret File Detection

The following files are treated as secrets and not read:
- `.env`, `.env.*`
- `*.pem`
- `id_rsa`, `id_rsa.pub`, `id_ed25519`, `id_ed25519.pub`
- `*.key`
- `.credentials*`, `.aws/credentials`, `.gcp/credentials*`
- `.token*`

## Main Branch Protection

- Direct changes to the `main` branch are prohibited
- A new branch is created for each run

## Run Isolation

- Each run creates `.agentos/runs/{task_id}/` for all artifacts
- All file changes are tracked via git diff
- Before/after state is preserved

## Limits

- Maximum changed files: configurable (default 20)
- Maximum retries on test/lint failure: configurable (default 3)
- Maximum runtime: configurable (default 30 minutes)
