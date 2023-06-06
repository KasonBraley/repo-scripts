# Scripts

Leverage https://github.com/lindell/multi-gitter to apply a change to multiple repositories.

Scripts are located in the [cmd](./cmd) directory.

- update-action
    - Uses `sed` to perform a search & replace. Primarily used for GitHub yaml workflow files, but can
    be used for more than that.
- update-go-version
    - Update the Go version in a go.mod file.
- update-go-package
    - Update a single Go package or all packages in a go.mod file.

- Use `--dry-run` flag to validate before running for real.

Example:
```
multi-gitter run "$PWD/cmd/update-checkout-action/update-checkout-action" --dry-run --ssh-auth --concurrent 1 --log-level=debug --git-type=cmd -m "ci: Bump actions/checkout from v2 to v3" -B bump-checkout --assignees KasonBraley --reviewers foo -R KasonBraley/foo
```
