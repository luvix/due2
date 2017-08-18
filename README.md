# sche
Alternative job scheduler to cron.

## Feature
- schedule by every min, hour, day, week, month, years
- job start at HH:mm
- job list
- support multithreading

## Interface
### Schedule interface
|cmd|Req/Opt|description|example|
|---|---|---|---|
|`--at`, `-t`|Required|The time when the job must start. (Format: HH:mm)|`sche 'echo hello world' --at=8:45`, `sche 'echo hello world' -a 8:45`|
|`--days`, `-d`|Optional|schedule routine (interval unit: day)|`sche 'echo hello world' --days=1 --at=8:45`, `sche 'echo hello world' -d 1 -a 8:45`|
|`--weeks`, `-w`|Optional|schedule routine (interval unit: week)|`sche 'echo hello world' --weeks=1 --at=8:45`, `sche 'echo hello world' -w 1 -a 8:45`|
|`--months`, `-m`|Optional|schedule routine (interval unit: month)|`sche 'echo hello world' --months=1 --at=8:45`, `sche 'echo hello world' -m 1 -a 8:45`|
|`--years`, `-y`|Optional|schedule routine (interval unit: year)|`sche 'echo hello world' --years=1 --at=8:45`, `sche 'echo hello world' -y 1 -a 8:45`|

### Job interface
|cmd|Req/Opt|description|example|
|---|---|---|---|
|`--cmd`, `-c`|Rqeuired|Specify command. Omittable if the command is at front or rear.|`sche --cmd='echo hello world' -a 8:45`, `sche -c 'echo hello world' -a 8:45`, `sche 'echo hello world' -a 8:45`, `sche -a 8:45 'echo hello world'`|
|`--stdout`, `-o`|Optional|File path for stdout. If omitted, stdout filepath is same as the path command `sche` ran.|`sche -a 8:45 --stdout='here.txt' 'echo hello world'`, `sche -a 8:45 -o 'here.txt' 'echo hello world'`|
|`--parallel`, `-p`|Optional|Run jobs as parallel by process or thread. Default is `process`.|`sche 'echo hello world' -a 8:45 --parallel=thread`, `sche 'echo hello world' -a 8:45 -p process`|
|`--name`, `-n`|Optional|Name the job.| `sche 'echo hello world' -a 8:45 -n hello` |  
|`--tag`, `-t`|Optional|put tag on the job. You can tag multiple keys separated with comma.| `sche 'echo hello world' -a 8:45 -t this,is,connect,ring,between,you,and,me` |  

### Management interface
|cmd|Req/Opt|description|example|
|---|---|---|---|
|`--list`, `-l`|Optional|Load scheduled job list. The list is owned by user.|`sche -l`, `sche --list`|
|`--list-all`|Optional|Load schedule job list. The job list is readable by root. This command can run only sudoers.|`sche --list-all`|
|`--delete`, `-d`|Optional|Delete the job in the list.|  |  
## Example
**Run echo everyday at 8:45**  
`sche 'echo hello world' --days=1 --at=8:45`

**Run echo at 8:45 (on next day if registed over 8:45)**  
`sche 'echo hello world' --at=8:45`
`sche 'echo hello world' -t=8:45`

**Run echo every week on friday at 8:45**  
`sche 'echo hello world' --weeks=1 --on=fri --at=8:45`
