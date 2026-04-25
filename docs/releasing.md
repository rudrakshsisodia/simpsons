# Releasing

Simpsons uses [GoReleaser](https://goreleaser.com/) via GitHub Actions. Pushing a version tag builds binaries, creates a GitHub release, and updates the Homebrew formula.

## Steps

1. Make sure `main` is clean and all tests pass:

   ```
   make build && make test && make lint
   ```

2. Tag the release:

   ```
   git tag v0.X.Y
   git push origin v0.X.Y
   ```

3. The `Release` workflow (`.github/workflows/release.yml`) runs automatically and:
   - Builds binaries for darwin/linux (amd64/arm64)
   - Creates a GitHub release with the binaries and checksums
   - Pushes the Homebrew formula to `rudrakshsisodia/homebrew-tap` → `Formula/simpsons.rb`

4. Monitor the workflow:

   ```
   gh run list --workflow=release.yml --limit 1
   gh run watch <run-id>
   ```

## Users install via

```
brew tap rudrakshsisodia/tap
brew install simpsons
```

or

```
go install github.com/rudrakshsisodia/simpsons@latest
```

## If a release fails

If the workflow fails partway (e.g. binaries uploaded but formula push failed):

1. Delete the release and re-tag:

   ```
   gh release delete vX.Y.Z --yes
   git push origin :refs/tags/vX.Y.Z
   git tag -d vX.Y.Z
   git tag vX.Y.Z
   git push origin vX.Y.Z
   ```

2. This triggers a fresh workflow run.

## Secrets

The release workflow needs `HOMEBREW_TAP_TOKEN` — a fine-grained GitHub PAT with **Contents: Read and write** on the `rudrakshsisodia/homebrew-tap` repo. Stored in the simpsons repo's Actions secrets.

Note: The `rudrakshsisodia` org requires fine-grained tokens; classic PATs are blocked.
