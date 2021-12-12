git rebase -i looppointer
go mod edit -require=github.com/kyoh86/looppointer@latest
go mod tidy
go mod edit -replace=golang.org/x/tools=github.com/kamilsk/go-tools@v0.7.0
go mod tidy
git add .
eval GIT_COMMITTER_DATE=2023-04-09T10:54:51+03:00 git commit --amend --date=2023-04-08T22:54:11+03:00 --no-edit
git rebase --continue
git fetch origin --force --prune --tags
git fetch upstream --force --prune --tags
git rebase --onto v1.52.2 upstream/master
git tag -d v1.52.2 || true
git tag -m 'sync with upstream' v1.52.2
git tag -d looppointer || true
git push origin :refs/tags/looppointer || true
git tag -m 'extend by looppointer' looppointer da04413a8a1eefb8c10161c9f2b558138d01815c
git push --force-with-lease origin extended
git push origin refs/tags/looppointer
git push origin refs/tags/v1.52.2
gh release -R kamilsk/golangci-lint delete v1.52.2 -y || true
gh release -R kamilsk/golangci-lint edit looppointer --draft=false --prerelease
goreleaser --clean --release-notes CHANGES.md
rm -rf dist
gh release -R kamilsk/golangci-lint view v1.52.2 -w
