# yaksok
*Python Framework for job routine.*

yaksok은 반복시켜야하는 job을 반복시키고 반복되는 작업을 관리하는 프레임워크입니다.
프로그램 내부적으로 스케쥴표에 접근하여 등록/수정/삭제할 수 있습니다.
또한 flask를 활용하여 RESTapi와 webhook을 제공합니다.
이를 통해 외부 프로그램이 yaksok의 일정에 맞추어 함께 동작할 수 있습니다.

## Prerequirements
**Python Libraries**
- schedule(>=0.4.3)
- flask(1.11)

## Feature
- schedule by every min, hour, day, week, month,
- job start at HH:MM:ss
- job list
- RESTful HTTP API(Optional)
- Webhook(pre, post)

## Interface
> **Deprecated**  
> This interface will be changed to HTTP api.

### Schedule interface
|cmd|Req/Opt|description|example|
|---|---|---|---|
|`--at`, `-t`|Required|The time when the job must start. (Format: HH:mm)|`yaksok 'echo hello world' --at=8:45`, `yaksok 'echo hello world' -a 8:45`|
|`--days`, `-d`|Optional|schedule routine (interval unit: day)|`yaksok 'echo hello world' --days=1 --at=8:45`, `yaksok 'echo hello world' -d 1 -a 8:45`|
|`--weeks`, `-w`|Optional|schedule routine (interval unit: week)|`yaksok 'echo hello world' --weeks=1 --at=8:45`, `yaksok 'echo hello world' -w 1 -a 8:45`|


### Job interface
|cmd|Req/Opt|description|example|
|---|---|---|---|
|`--cmd`, `-c`|Rqeuired|Specify command. Omittable if the command is at front or rear.|`yaksok --cmd='echo hello world' -a 8:45`, `yaksok -c 'echo hello world' -a 8:45`, `yaksok 'echo hello world' -a 8:45`, `yaksok -a 8:45 'echo hello world'`|
|`--stdout`, `-o`|Optional|File path for stdout. If omitted, stdout filepath is same as the path command `yaksok` ran.|`yaksok -a 8:45 --stdout='here.txt' 'echo hello world'`, `yaksok -a 8:45 -o 'here.txt' 'echo hello world'`|
|`--parallel`, `-p`|Optional|Run jobs as parallel by process or thread. Default is `process`.|`yaksok 'echo hello world' -a 8:45 --parallel=thread`, `yaksok 'echo hello world' -a 8:45 -p process`|
|`--name`, `-n`|Optional|Name the job.| `yaksok 'echo hello world' -a 8:45 -n hello` |  
|`--tag`, `-t`|Optional|put tag on the job. You can tag multiple keys separated with comma.| `yaksok 'echo hello world' -a 8:45 -t this,is,connect,ring,between,you,and,me` |  

### Management interface
|cmd|Req/Opt|description|example|
|---|---|---|---|
|`--list`, `-l`|Optional|Load scheduled job list. The list is owned by user.|`yaksok -l`, `yaksok --list`|
|`--list-all`|Optional|Load schedule job list. The job list is readable by root. This command can run only sudoers.|`yaksok --list-all`|
|`--delete`, `-d`|Optional|Delete the job in the list.| `yaksok --delete jobname` |  

## Example
**Run echo everyday at 8:45**  
`yaksok 'echo hello world' --days=1 --at=8:45`

**Run echo at 8:45 (on next day if registed over 8:45)**  
`yaksok 'echo hello world' --at=8:45`
`yaksok 'echo hello world' -t=8:45`

**Run echo every week on friday at 8:45**  
`yaksok 'echo hello world' --weeks=1 --on=fri --at=8:45`
