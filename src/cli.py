"""

"""
import os
from argparse import ArgumentParser
from core import Yaks


def is_if_hasattr(cls, attr):
    if hasattr(cls, attr):
        return getattr(cls, attr)
    return False


def parse_cli(yaks: Yaks):
    parser = ArgumentParser(description="Alternative job scheduler to cron.")

    ### Management interface
    parser.add_argument('--list', dest='list', nargs='?',
                        help="Load scheduled job list. The list is owned by user.")
    parser.add_argument('--list-all', dest='list-all', action='store_true',
                        help="Load schedule job list. The job list is readable by root. "
                             "This command can run only sudoers.")

    ### Schedule interface
    parser.add_argument('-a', '--at', dest='at', nargs=1,
                        help="Specify What time is the job start.")
    parser.add_argument('-d', '--day', dest='day', nargs=1,
                        help="Specify What day is the job start.")
    parser.add_argument('-w', '--week', dest='week',
                        help="Specify What week is the job start.")
    parser.add_argument('-m', '--month', dest='month',
                        help="Specify What month is the job start.")
    parser.add_argument('-y', '--year', dest='year',
                        help="Specify What year is the job start.")
    return parser.parse_args()


def runapp(parsed: ArgumentParser):
    """

    :type parsed: ArgumentParser
    """
    if is_if_hasattr(parsed, 'list-all'):
        print("is_if_hasattr")
        yaks.listall()
    if hasattr(parsed, 'list'):
        print("hasattr")
        # if args.list:
        #     print("I am the list.")
    else:
        print("Who am I?")
    pass


if __name__ == '__main__':
    yaks = Yaks()
    commands = parse_cli(yaks)
    runapp(commands)
