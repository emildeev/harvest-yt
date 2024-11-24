# Gitlab CLI helper
This cli application is used to create merge requests in gitlab
It uses gitlab api, yandex tracker api to create merge request with Title and Description,
and git commands for getting current brunch and repository

## Installation

```bash
go install --mod=mod github.com/emildeev/harvest_yt/hytmigrator
```

## Usage

### Show all commands
```bash
gilatb --help
```

### Start configure
```bash
hytmigrator configure
```
This command will request your harvest token, yandex tracker organization id and token
for help enter empty value

### Migrates timers
```bash
hytmigrator migrate
```
This command will migrate all timers from harvest to yandex tracker
### Logs
If you have any problems, you can turn on logging, run commands with flag -l {log_level}. For example:
```bash
hytmigrator migrate -l debug
```