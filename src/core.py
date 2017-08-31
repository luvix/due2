import os
import sqlite3


def initial_program_setup():
    pass


def do_main_program():
    pass


def program_cleanup():
    pass


def reload_program_config():
    pass


class Yaks:
    def __init__(self):
        self._db = Alchemist()
        pass

    def list(self, *args):
        """
        show job list.
        :param args: show jobs which has the tags in the parameter.
        :return:
        """
        self._db.list()
        print("list in yaks")
        pass

    def listall(self):
        """
        list all jobs.
        :return:
        """
        print("listall in yaks")
        pass
    pass


class Alchemist:
    def __init__(self):
        self._conn = sqlite3.connect('yaks.sqlite')
        self._csrr = self._conn.cursor()

    def list(self, *args):
        result = self._csrr.execute("SELECT * FROM jobs WHERE user = ?", "user")
        return result

    def __del__(self):
        self._conn.close()
