<<<<<<< HEAD
# Yaksok

Alternative job yaksokduler to cron.

## Feature

- Schedulable by every sec, min, hour, day, week, month, and year
- Support job list
- Job management on sqlite

## Format

Yaksok use two type of format to specifying when the job is run.

### Date format

*Date Format* is a format for specifying the date when the job is run. 
_TODO: 어떤 명령어에서 사용할 수 있는 지 설명해주세요_
The format examples is following:

- Second: 10s
- Minute: 10m(if you does not care any second), 10m10s
- Hour: 10h(if you does not care any minute), 10h10m, 10h10m10s
- Day: 10d, 10d10h, 10d10h10m, 10d10h10m10s
- month: 10M, 10M10d, 10M10d10h, 10M10d10h10m, 10M10d10h10m10s
- year: 2010y, 2010y10M, 2010y10M10d, 2010y10M10d10h, 2010y10M10d10h10m, 2010y10M10d10h10m10s
- shortyear: 10y, 10y10M, 10y10M10d, 10y10M10d10h, 10y10M10d10h10m, 10y10M10d10h10m10s

ISO format is also support, but you write on fully date, only date or only time.

``` sh
2010-10-10 10:10:10 # correct
10-10-10 10:10:10 # correct
2010-10-10 0:0:0 # correct
2010-10-10 0:0 # correct (run at midnight)
2010-10-10 0 # correct (run at midnight)
2010-10-10 # correct (run at anytime)
10:10:10 # correct
10:10 # run at 10:10:00
10 # if --at, run at 10:00:00. If --day, run everyday at 10:00:00. If --month, run on 10th date at anytime. Error if you apply on the option --year or --week.
10 10:10:10 # correct (works only --month. run following date. It works if today is before than the date but It works next month if today is after than.)
10-10 10:10:10 # correct (works only --year. run following date. It works if today is before than the date but It works next year if today is after than.)
```

## Day format

다음의 상황을 위해 날짜 형식이 아닌 요일 형식으로 입력할 수 있습니다.

- 무슨 요일에 작동
- 몇 째 주에 작동
- 몇 월에 작동(short format)
- 몇 번째 월에 작동

*Day format* 은 두가지 방식으로 입력할 수 있습니다. 해당 요일, 주, 월의 영어명을 사용하거나 숫자 뒤에 wd(무슨 요일), ws(몇 째 주), tm(몇 월), ms(몇 번째 월)을 붙입니다. 날짜는 일요일부터 시작하며, **1** 부터 시작합니다.
마지막주를 지원합니다. 마지막주는 **-1** 입니다.
현재를 기준으로 하지 않고 한 해를 기준으로 계산할 때도 있습니다. 분기, 반기 등이 그러합니다. 이때에는 `yly`접두어를 붙여서 지정합니다.

### Examples

- 두째주 수요일: 2ws4wd, 2wswed
- 10월 두째주 수요일: 10tm2ws4wd, oct2ws4wd, oct2wswed
- 매 분기 마지막주 금요일: yly3ms-1ws6wd, yly3ms-1wsfri
- 지금부터 5개월마다 마지막주 금요일: 5ms-1ws6wd, 5ms-1wsfri

## Example

### Run echo everyday at 8:45

### Run echo at 8:45 (on next day if registed over 8:45)

### Run echo every week on friday at 8:45
=======
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
>>>>>>> remotes/origin/dev
