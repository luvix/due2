# Yaksok

Alternative job yaksokduler to cron.

[![Build Status](https://travis-ci.org/theodore-kim/yaksok.svg?branch=dev)](https://travis-ci.org/theodore-kim/yaksok)
[![Go Report Card](https://goreportcard.com/badge/github.com/theodore-kim/yaksok)](https://goreportcard.com/report/github.com/theodore-kim/yaksok)

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
