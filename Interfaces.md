# Yaksok interface

## Common options

| cmd | Req/Opt | description | example | 
| --- | --- | --- | --- | 
| `--cmd`, `-c` | Rqeuired | Specify command. Omittable if the command is at front or rear. | `yaksok --cmd='echo hello world' -a 8:45`, `yaksok -c 'echo hello world' -a 8:45`, `yaksok 'echo hello world' -a 8:45`, `yaksok -a 8:45 'echo hello world'` | 
| `--name`, `-n` | Optional | Name the job. It should be unique. |  `yaksok 'echo hello world' -a 8:45 -n hello`  | 
| `--tag`, `-t` | Optional | put tag on the job. You can tag multiple keys separated with comma. |  `yaksok 'echo hello world' -a 8:45 -t this,is,connect,ring,between,you,and,me`  |

## once

`once` command uses 1 option, `at`.
If you does not specify the time, it works at `anytime` on next day.

``` sh
yaksok once "echo 'hell world'" --at 8:45
yaksok once "echo 'hell world'" --at 8:45 12:30 # works twice
```

## daily

`once` command uses 1 option, `at`.
If you does not specify the time, it works at `anytime` on everyday.

``` sh
yaksok daily "echo 'hell world'" --at 8:45
yaksok daily "echo 'hell world'" --at 8:45 12:30  # works twice everyday
```

## weekly

`weekly` command use twe ontions, `on` and `at`.
If you does not specify the time, it works at `anytime`.
If you does not specify the day, it works on `firstday`.
You can set the `firstday` by useing `preference` command.

``` sh
yaksok weekly "echo 'hell world'" --on fri --at 8:45
yaksok weekly "echo 'hell world'" --at fri 8:45  # same as above.
yaksok daily "echo 'hell world'" --at 8:45 12:30  # works twice on the firstday
yaksok weekly "echo 'hell world'" --at 8:45  # work on firstday
yaksok weekly "echo 'hell world'" --on mon wed fri --at 8:45  # works on monday, wednesday, friday

yaksok weekly "echo 'hell world'" --on sat --at fri 8:45 # ERROR
```

## monthly

`weekly` command use twe ontions, `on` and `at`.
If you does not specify the time, it works at `anytime`.
If you does not specify the day, it works on *1st day in month*.

``` sh
yaksok monthly "echo 'hell world'" --on fri --at 8:45
yaksok monthly "echo 'hell world'" --at fri 8:45
yaksok monthly "echo 'hell world'" --on sat --at fri 8:45 # ERROR
yaksok add "echo \'hi\'" --on 2wk2wd --at 8:45  # run a job every second week on monday (second weekday) at anytime. Supports only dateformat.
yaksok add "echo \'hi\'" --on 2wk2wd 08:45  # run a job every second week on monday (second weekday) at 08:45:00. Supports only dateformat.
yaksok add "echo \'hi\'" --on 2wk2wd --at 8:45  # run a job every second week on monday (second weekday) at 08:45:00. Supports only dateformat.
```

## yearly

`weekly` command use twe ontions, `on` and `at`.
If you does not specify the time, it works at `anytime`.
If you does not specify the day, it works on *january 1st*.

``` sh
yaksok yearly "echo \'hi\'" --month=10md
yaksok yearly "echo \'hi\'" --month=10md 8:45
yaksok yearly "echo \'hi\'" --month=10md --at 8:45
yaksok yearly "echo \'hi\'" --month=2wk --at 8:45  # run a job on sunday
```

| cmd | Req/Opt | description | example |
| --- | --- | --- | --- |
| `--at`, `-t`  |   Optional  |  Schedule a job. It works just once.  |  `yaksok 'echo hello world' --at=8:45`, `yaksok 'echo hello world' -t 8:45`|
| `--day`, `-d` | Optional | Schedule a job. It works everyday. You can schedule multiple times. critaria is comma(`,`) | `yaksok 'echo hello world' --day=8:45`, `yaksok 'echo hello world' -d 8:45,12:30` |
| `--week`, `-w` | Optional | Schedule a job. It works everyweek. You can schedule multiple days. critaria is comma(`,`) | `yaksok 'echo hello world' --weeks=1 --at=8:45`, `yaksok 'echo hello world' -w 1 -a 8:45` |
| `--month`, `-m` | Optional | yaksokdule routine (interval unit: month) | `yaksok 'echo hello world' --months=1 --at=8:45`, `yaksok 'echo hello world' -m 1 -a 8:45` |
| `--year`, `-y` | Optional | yaksokdule routine (interval unit: year) | `yaksok 'echo hello world' --years=1 --at=8:45`, `yaksok 'echo hello world' -y 1 -a 8:45` |

## hourly

``` sh
yaksok hourly "echo 'hell world'" # works every hour on time.
yaksok hourly "echo 'hell world'" --at 30  # works every hour on 30 minute.
yaksok hourly "echo 'hell world'" --at 30:30  # works every hour on 30 minute 30 second.
yaksok hourly "echo 'hell world'" --at 0 20 40   # works every hour on time, 20 minute, 40 minte.
```

## minutely

``` sh
yaksok minutely "echo 'hell world'" # works every minute at time.
yaksok minutely "echo 'hell world'" --at 0 20 40 # works every hour at time, 20 second, 40 second.
```

## secondly

no more optional option without command.

``` sh
yaksok secondly "echo 'hell world'" # works every second.
```

## list

show scheduled jobs list. The list is composed by user.

``` sh
yaksok list  # show jobs which you have auth.
yaksok list --all # show jobs which you have auth.
yaksok list --who {user}  # show whose jobs which you have auth.
yaksok list --name {name}  # show jobs whose name has queried name as substring
yaksok list --tag {tag}  # show jobs whose name has queried tag (not substring)
yaksok list --id {jobID}  # show job whose jobID is same as queried jobID
```

## delete

Delete the job in the list.

``` sh
yaksok delete {jobID}  # delete by jobID
yaksok delete --name {name}  # delete by name
yaksok delete --tag {tag}  # delete by tag
```

## prefrerence

``` sh
yaksok preference --anytime 04:00:00  # run a job which is specified to anytime
```

| cmd | description |
| --- | --- |
| `--anytime` | 아무 시간에나 실행할 프로그램이 실행되도 괜찮은 시간을 지정합니다. |
| `--firstday` | 매 주의 첫째날을 지정합니다. 기본은 일요일입니다. |