import os
# import grp  # group id processor
import signal
import daemon
import lockfile

from .core import (
    initial_program_setup,
    do_main_program,
    program_cleanup,
    reload_program_config,
    )

context = daemon.DaemonContext(
    working_directory='/var/lib/yaks',
    umask=0o002,
    pidfile=lockfile.FileLock('/var/run/yaks.pid'),
    )

context.signal_map = {
    signal.SIGTERM: program_cleanup,        # signal to terminate daemon, like `systemctl yaks stop`
    signal.SIGHUP: 'terminate',             # signal to control by ctrl+D
    signal.SIGUSR1: reload_program_config,  # user defined signal 1...I don't know how to use it.
    }

# mail_gid = grp.getgrnam('mail').gr_gid
# context.gid = mail_gid

db_path = './.db/jobs.sq3'
context.files_preserve = [db_path]

initial_program_setup()

with context:
    do_main_program()