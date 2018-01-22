# Yaksok interface

## once

`once` command uses 1 option, `at`.
If you does not specify the time, it works at `anytime` on next day.

| cmd| description |
| ---| --- |
| `--at` | Schedule a job. It works just once. |
| `--job` | command what you want to schedule. If your job command is next to keyword, option is not to be needed. |

``` sh
yaksok once --job "echo 'hell world'" --at 8:45
yaksok once "echo 'hell world'" --at 8:45
yaksok once "echo 'hell world'" --at 8:45 12:30 # works twice
```

## daily

`once` command uses 1 option, `at`.
If you does not specify the time, it works at `anytime` on everyday.

| cmd| description |
| ---| --- |
| `--at` | Schedule a job. It works just once. |
| `--job` | command what you want to schedule. If your job command is next to keyword, option is not to be needed. |

``` sh
yaksok daily --job "echo 'hell world'" --at 8:45
yaksok daily "echo 'hell world'" --at 8:45
yaksok daily "echo 'hell world'" --at 8:45 12:30  # works twice everyday
```

## weekly

`weekly` command use twe ontions, `on` and `at`.
If you does not specify the time, it works at `anytime`.
If you does not specify the day, it works on `firstday`.
You can set the `firstday` by useing `preference` command.

| cmd| description |
| ---| --- |
| `--at` | Schedule a job when it work on time. It works just once. |
|`--on`| Schedule a job when it works on day. It works everyweek. |
| `--job` | command what you want to schedule. If your job command is next to keyword, option is not to be needed. |

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

| cmd| description |
| ---| --- |
| `--at` | Schedule a job when it work on time. It works just once. |
|`--on`| Schedule a job when it works on day or week. It works everymonth. |
| `--job` | command what you want to schedule. If your job command is next to keyword, option is not to be needed. |

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

| cmd| description |
| ---| --- |
| `--at` | Schedule a job when it work on time. It works just once. |
|`--on`| Schedule a job when it works on date, day, week or month. It works every year. |
| `--job` | command what you want to schedule. If your job command is next to keyword, option is not to be needed. |

``` sh
yaksok yearly "echo \'hi\'" --month=10md
yaksok yearly "echo \'hi\'" --month=10md 8:45
yaksok yearly "echo \'hi\'" --month=10md --at 8:45
yaksok yearly "echo \'hi\'" --month=2wk --at 8:45  # run a job on sunday
```

| cmd| description |
| ---| --- |
| `--at` | Schedule a job. It works just once. |
| `--job` | command what you want to schedule. |

## hourly

| cmd| description |
| ---| --- |
| `--at` | Schedule a job when it works on minute. It works just once. |
| `--job` | command what you want to schedule. If your job command is next to keyword, option is not to be needed. |

``` sh
yaksok hourly "echo 'hell world'" # works every hour on time.
yaksok hourly "echo 'hell world'" --at 30  # works every hour on 30 minute.
yaksok hourly "echo 'hell world'" --at 30:30  # works every hour on 30 minute 30 second.
yaksok hourly "echo 'hell world'" --at 0 20 40   # works every hour on time, 20 minute, 40 minte.
```

## minutely

| cmd| description |
| ---| --- |
| `--at` | Schedule a job when it works on second. It works just once. |
| `--job` | command what you want to schedule. If your job command is next to keyword, option is not to be needed. |

``` sh
yaksok minutely "echo 'hell world'" # works every minute at time.
yaksok minutely "echo 'hell world'" --at 0 20 40 # works every hour at time, 20 second, 40 second.
```

## secondly

This command does not explicit when it work. It just run a job every second.

| cmd| description |
| ---| --- |
| `--job` | command what you want to schedule. If your job command is next to keyword, option is not to be needed. |

``` sh
yaksok secondly --job "echo 'hell world'" # works every second.
yaksok secondly "echo 'hell world'" # works every second.
```

## list

show scheduled jobs list. The list is composed by user.

``` sh
yaksok list  # show jobs which you scheduned or had auth.
yaksok list --all  # show jobs which you scheduled and others job whant you can read with auth.
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