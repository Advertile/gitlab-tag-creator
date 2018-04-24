# GitLab Tag Creator

This simple GO program pulls the latest tag (version) from GitLab, increases the version and creates a new tag.

## Install

```bash
# Find the available versions here: https://github.com/Advertile/gitlab-tag-creator/releases
export VERSION=x.x.x
wget -O gitlab-tag-creator https://github.com/Advertile/gitlab-tag-creator/releases/download/${VERSION}/gitlab-tag-creator_${VERSION}_linux_amd64
chmod u+x gitlab-tag-creator
```

## Required Environment Variables

- `CI_COMMIT_SHA` - The commit revision for which project is built.
- `CI_PROJECT_ID` - The project ID which is related to your repository. [Read here](https://stackoverflow.com/a/45500237/325852) how to find it.
- `GITLAB_TOKEN` - The token used to communicate with the GitLab API.

## Usage

Update the major version/tag (from `2.1.1` to `3.0.0`):

```bash
gitlab-tag-creator update major
```

Update the minor version/tag (from `2.1.1` to `2.2.0`):

```bash
gitlab-tag-creator update minor
```

Update the patch version/tag (from `2.1.1` to `2.1.2`):

```bash
gitlab-tag-creator update patch
```
