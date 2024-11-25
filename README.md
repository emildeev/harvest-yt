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
hytmigrator --help
```

### Start configure
```bash
hytmigrator configure
```
This command will request your harvest account id and token, yandex tracker organization id and token
for help enter empty value

Generate harvest token and get harvest id you can [here](https://id.getharvest.com/oauth2/access_tokens/new)
Get yandex tracker organization ID you can [here](https://tracker.yandex.ru/settings) :
![img.png](docs/org_id.png)
Generate yandex tracker token you can [here](https://oauth.yandex.ru/authorize?response_type=token&client_id=711865fe0ef3478ea09e895878cd275b)


### configure product tasks 
In product tasks migrator expect yandex tracker task key in harvest timer description, like:
![img.png](docs/product_task.png)

this tasks spend without comment

for get all configured tasks run
```bash
hytmigrator configure_tasks developer
```

for update tasks configuration run
```bash
hytmigrator configure_tasks developer_update
```

### configure communication tasks
In communication tasks migrator get yandex tracker task key from config and harvest timer description add to spend 
comment

for get all configured tasks run
```bash
hytmigrator configure_tasks communication
```

for update tasks configuration run
```bash
hytmigrator configure_tasks communication_update
```

### configure skipped tasks
Migrator will ignore all tasks from this list

for get all configured tasks run
```bash
hytmigrator configure_tasks skipped
```

for update tasks configuration run
```bash
hytmigrator configure_tasks skipped_update
```

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