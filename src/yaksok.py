import os
import schedule
from .alchemy import YakSokController
# from configparser import ConfigParser
# _config_ = ConfigParser().read('.yaks.cfg')


def initial_program_setup():
    pass


def do_main_program():
    pass


def program_cleanup():
    pass


def reload_program_config():
    pass


class YakSok(schedule.Job):
    def __init__(self, interval):
        super().__init__(interval)

    def list(self, *args, **kwargs):
        """
        Inquery registered jobs.
        Inquering criteria: tag, uid, name, group
        :param args: For only tag
        :param kwargs:
        :return:
        """
        if len(args) > 0:
            # TODO Inquering by tags
            pass
        elif len(kwargs) > 0:
            if 'uid' in kwargs:
                # TODO Inquering by uid
                pass
            elif 'name' in kwargs:
                # TODO Inquering by job name
                pass
            elif 'group' in kwargs:
                # TODO Inquering by job group
                pass
        pass

    def do(self, job_func, *args, **kwargs):
        # TODO Override Job.do
        # insert db
        return super().do(job_func, *args, **kwargs)

    def listall(self):
        pass

    def insert(self, *args, **kwargs):
        pass

    def delete(self, *args, **kwargs):
        pass

    def update_tags(self, uid: str, tags: list or str):
        pass


