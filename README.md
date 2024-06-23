# Gitlab CLI helper
This cli application is used to create merge requests in gitlab
It uses gitlab api, yandex tracker api to create merge request with Title and Description,
and git commands for getting current brunch and repository

## Installation

```bash
go install --mod=mod github.com/emildeev/gitlab_helper/gitlab
```

## Usage

### Show all commands
```bash
gilatb --help
```

### Start configure
```bash
gitlab configure
```
This command will request your gitlab url and token, yandex tracker organization id and token
for help enter empty value

### Create merge requests
```bash
gitlab create
```
This command get required information from git and yandex tracker and ask addition information in cli for creating
merge requests from current brunch to target main brunch (configured) and additional brunches (configured).
After creating merge requests it will show links to created merge requests and will set it to ticket in yandex tracker

### Logs
If you have any problems, you can turn on logging, run commands with flag -l {log_level}. For example:
```bash
gitlab create -l debug
```