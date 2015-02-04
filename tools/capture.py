#!/usr/bin/env python
#
# Copyright 2015 The Android Open Source Project
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""This script captures a graphics trace using libgfxspy.

Run with --help command-line option to see usage and parameter information.
"""

import logging
import optparse
import os
import socket
import subprocess
import sys
import time

# Path on the Android device where libgfxspy.so is installed.
GFXSPY_INSTALL_DIR = '/data/'
# Name of the file to install.
GFXSPY_FILE_NAME = 'libgfxspy.so'
# Path to the gfxspy library on the host machine.
DEFAULT_GFXSPY_HOST_PATH = './pkg/build/libs/armeabi-v7a/' + GFXSPY_FILE_NAME

# Port number for adb forward to use on local host machine.  This is the
# same port used by the eclipse gltrace plugin, but it can be any port.
LOCAL_TCP_PORT_NUMBER = 6039

# The socket to bind to on the target Android device.  This needs to be the
# same as the socket opened by libgfxspy.
REMOTE_SOCKET_NAME = 'localabstract:gltrace'

# Stop trying to connect after 2 minutes.
CONNECT_TIMEOUT_SECONDS = 120.0
# Don't try connecting more than 5 times per second.
CONNECT_RETRY_DELAY = 0.2

# Timeout if no data is received for 1/2 a second
SOCKET_RECEIVE_TIMEOUT = 0.5
# Standard network data packed size (recommended by Python docs).
SOCKET_RECEIVE_DATA_SIZE = 4096

# Extension used by the output trace files.
TRACE_FILE_EXTENSION = '.gltrace'

# The maximum length for Android system property names.
ANDROID_PROP_NAME_MAX = 32

# The maximum length for Android system property values.
ANDROID_PROP_VALUE_MAX = 92

# Options used to launch adb.
ADB_APP_NAME_PROPERTY = 'adb_app_name'
DEVICE_SERIAL_PROPERTY = 'device_serial'
_global_adb_options = {
    # The name of the adb executable.
    ADB_APP_NAME_PROPERTY: 'adb',

    # If non-empty will be passed to adb with -s option.
    DEVICE_SERIAL_PROPERTY: '',
}

# Used to temporarily change and then restore the logging level.
_logging_level_stack = []


def get_current_logging_level():
  """Returns an integer representing the current logging level.

  See logging Python docs for description of the levels.

  Returns:
    An integer representing the current logging level.
  """
  assert _logging_level_stack
  return _logging_level_stack[-1]


def push_logging_level(new_logging_level):
  """Changes the current logging level.

  Args:
    new_logging_level: An integer for the new logging level.
  """
  _logging_level_stack.append(new_logging_level)
  logging.getLogger().setLevel(new_logging_level)


def pop_logging_level():
  """Restores the previous logging level.
  """
  _logging_level_stack.pop()
  logging.getLogger().setLevel(get_current_logging_level())


def log_to_stdout(level, message):
  """Writes the given message to stdout.

  If logging is enabled for the given level, writes a message directly
  to stdout.  Otherwise, does nothing.

  The default Python logging inserts new-lines at the end of each message.
  log_to_stdout() does not insert new-lines or any other formatting.

  Args:
    level: The logging level for this message.
    message: The message text.
  """
  if get_current_logging_level() <= level:
    sys.stdout.write(message)
    sys.stdout.flush()


def get_adb_app_name():
  """Returns the name of the adb application.

  Returns:
    The name of the adb application.
  """
  return _global_adb_options[ADB_APP_NAME_PROPERTY]


def get_device_serial():
  """Returns the serial number of the target device.

  Returns:
    The serial number of the target device, or an empty string if no serial
    has been set.
  """
  return _global_adb_options[DEVICE_SERIAL_PROPERTY]


def set_device_serial(device_serial):
  """Sets the serial number of the target device.

  Args:
    device_serial: The serial number of the target device.
  """
  _global_adb_options[DEVICE_SERIAL_PROPERTY] = device_serial


def make_adb_command(adb_parameters):
  """Builds a full command-line needed to run adb.

  Adds the adb executable and device specifier to the given list of
  parameters.

  Args:
    adb_parameters: A list with parameters to pass to adb.

  Returns:
    A list representing the full command-line to start adb.
  """
  adb_command = [get_adb_app_name()]
  device_serial = get_device_serial()
  if device_serial:
    adb_command += ['-s', device_serial]
  return adb_command + adb_parameters


def run_adb(adb_parameters, quit_on_error=True):
  """Runs adb with the given list of parameters.

  Returns the output from adb (stdout + stderr) in a single string.
  This function will not return if adb returns an error-code (non-zero return
  value) and quit_on_error is True.

  Args:
    adb_parameters: The parameters to pass to adb.
    quit_on_error: If True, this script exit if adb exits with a non-zero
        return code.

  Returns:
    A string with the stdout and stderr from adb.
  """
  command = make_adb_command(adb_parameters)
  out = ''
  err = ''
  logging.info('> %s', ' '.join(command))
  try:
    process = subprocess.Popen(command, stdout=subprocess.PIPE,
                               stderr=subprocess.PIPE)
    out, err = process.communicate()
  except OSError as error:
    logging.error('The command "%s" failed with the following '
                  'error:', ' '.join(command))
    logging.error(error)
    sys.exit(1)

  return_code = process.returncode
  output_logging_level = logging.DEBUG
  if return_code:
    output_logging_level = logging.WARNING
  if out: log_to_stdout(output_logging_level, out)
  if err: log_to_stdout(output_logging_level, err)
  if return_code and quit_on_error:
    logging.error('The command "%s" exited with an error.',
                  ' '.join(command))
    sys.exit(1)
  return out + err


def do_install(options):
  """Based on the given options, installs libgfxspy.so on the target device.

  Args:
    options: A class with the command-line options used to start this script.
  """
  if not options.no_install:
    gfxspy_path = os.path.normpath(options.gfxspy_path)
    # For simplicity, the file will always be copied to the target device as
    # /data/libgfxspy.so, even if the library has a different name on the
    # host machine.
    run_adb(['push', gfxspy_path, GFXSPY_INSTALL_DIR + GFXSPY_FILE_NAME])


def build_trace_file_name(given_trace_file_name, app_name):
  """Constructs a file name to use for storing the captured trace.

  Args:
    given_trace_file_name: The name of the trace file specified on the
        command-line.
    app_name: The name of the Android application that is being traced.

  Returns:
    A file name.
  """
  trace_file_name = given_trace_file_name
  if not trace_file_name:
    (unused_str, unused_sep, app_name_suffix) = app_name.rpartition('.')
    trace_file_name = app_name_suffix + TRACE_FILE_EXTENSION
  elif '.' not in trace_file_name:
    trace_file_name += TRACE_FILE_EXTENSION
  return trace_file_name


def create_trace_file(trace_file_name):
  """Creates and opens a file object with the given name.

  Args:
    trace_file_name: The name of the file to create.

  Returns:
    A file object opened for writing.
  """
  logging.debug('Creating output file "%s".', trace_file_name)
  try:
    trace_file = open(trace_file_name, 'wb')
  except IOError as error:
    logging.error('Unable to create the file "%s":', trace_file_name)
    logging.error(error)
    sys.exit(1)
  return trace_file


def bytes_to_megabytes(num_bytes):
  """Converts bytes to Megabytes.

  Args:
    num_bytes: A number in bytes.

  Returns:
    The number of megabytes.
  """
  return num_bytes / (1024.0 * 1024.0)


def app_is_running(app_name):
  """Checks to see if the given application is running.

  Args:
    app_name: The Android application name.

  Returns:
    True if the application is running, and false otherwise.
  """
  push_logging_level(logging.WARNING)
  output = run_adb(['shell', 'ps', '|', 'grep', app_name])
  pop_logging_level()
  return len(output) > 0


def build_wrap_prop_string(app_name):
  """Creates a string to set properties for the given application.

  Args:
    app_name: The Android application name.

  Returns:
    A string that can be passed to adb shell set/getprop.
  """
  # Android does not allow property names longer than 32 characters, so we must
  # truncate the wrap property name.  Truncated wrap property names are still
  # matched correctly when the app is started.
  return ('wrap.%s' % app_name)[:(ANDROID_PROP_NAME_MAX - 1)]


def start_app(app_name, activity_name):
  """Launches the given app with the given activity.

  Args:
    app_name: The name of the Android application to run.
    activity_name: The name of the activity to run.
  """
  component_string = '%s/%s' % (app_name, activity_name)

  wrap_string = build_wrap_prop_string(app_name)

  ld_preload_string = 'LD_PRELOAD=' + GFXSPY_INSTALL_DIR + GFXSPY_FILE_NAME
  assert len(ld_preload_string) <= ANDROID_PROP_VALUE_MAX
  run_adb(['shell', 'setprop', wrap_string, ld_preload_string])

  # Verify that setprop worked.
  push_logging_level(logging.WARNING)
  adb_output = run_adb(['shell', 'getprop', wrap_string])
  pop_logging_level()
  if ld_preload_string not in adb_output:
    logging.error(get_adb_app_name() + ' setprop failed.  Aborting.')
    sys.exit(1)

  run_adb(['forward', 'tcp:' + str(LOCAL_TCP_PORT_NUMBER), REMOTE_SOCKET_NAME])
  app_start_output = run_adb(['shell', 'am', 'start', '-S', '-W', '-n',
                              component_string])

  if app_is_running(app_name):
    logging.info('Started.')
  else:
    # Print any adb output that may help diagnose the failure.  If debug logging
    # is on, the output has already been printed, so don't print it a
    # second time here.
    if app_start_output and get_current_logging_level() > logging.DEBUG:
      logging.info(app_start_output)
    logging.warning('Application "%s" does not appear to be running.',
                    app_name)
  logging.info('Press <ctrl> + c to stop.')


def update_capture_status(start_time, received_megabytes):
  """Prints a status message to stdout.

  Args:
    start_time: The time the capture was started.
    received_megabytes: The number of megabytes captured so far.
  """
  if received_megabytes == 0.0:
    log_to_stdout(logging.INFO, '.')
  else:
    log_to_stdout(logging.INFO, '\rCapture duration %.0fs.  Trace file '
                  'size: %.1f MB.' % (time.time() - start_time,
                                      received_megabytes))


def collect_trace(trace_file):
  """Captures a gfxspy trace.

  Connects to the gfxspy server on the target Android device and captures
  a trace.

  Args:
    trace_file: The file object where the trace will be stored.
  """
  received_megabytes = 0.0
  log_to_stdout(logging.INFO, 'Attempting to connect...')
  connect_start_time = time.time()
  capture_start_time = 0.0
  while (received_megabytes == 0.0) and (time.time() - connect_start_time <
                                         CONNECT_TIMEOUT_SECONDS):
    trace_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_address = ('localhost', LOCAL_TCP_PORT_NUMBER)
    trace_socket.connect(server_address)
    trace_socket.settimeout(SOCKET_RECEIVE_TIMEOUT)
    while True:
      update_capture_status(capture_start_time, received_megabytes)
      try:
        data = trace_socket.recv(SOCKET_RECEIVE_DATA_SIZE)
      except socket.timeout:
        pass
      else:
        if data:
          if received_megabytes == 0.0:
            capture_start_time = time.time()
            log_to_stdout(logging.INFO, '\nConnected.\n')
          received_megabytes += bytes_to_megabytes(len(data))
          trace_file.write(data)
        else:
          if received_megabytes == 0.0:
            time.sleep(CONNECT_RETRY_DELAY)
          else:
            log_to_stdout(logging.INFO, '\nConnection closed.\n')
          trace_socket.close()
          break


def set_permissions():
  """Puts adb in root mode and turns off SELinux enforcement.
  """
  adb_output = run_adb(['root'])
  # adb root does not return an error if it fails, so we must parse the
  # output string to see if we think it failed.  The most likely error is
  # attempting to run in root mode on a production build, so we'll look
  # for the word 'production' in the output.  With verbose logging on,
  # the adb output has already been printed, so don't log it again here.
  if 'production' in adb_output and get_current_logging_level() > logging.DEBUG:
    logging.info(adb_output)

  # Check to see if SELinux enforcing is on.
  push_logging_level(logging.WARNING)
  adb_output = run_adb(['shell', 'getenforce'])
  pop_logging_level()
  was_enforcing = 'enforcing' in adb_output.lower()

  # Turn off enforcement.  This will silently fail if the device is not rooted.
  run_adb(['shell', 'setenforce', '0'])

  # Check to see if enforcement was turned off.
  push_logging_level(logging.WARNING)
  adb_output = run_adb(['shell', 'getenforce'])
  pop_logging_level()
  is_enforcing = 'permissive' not in adb_output.lower()

  if is_enforcing:
    # SELinux enforcing is still on after we tried to turn it off.
    logging.warning('Unable to disable SELinux enforcing. '
                    'Capturing will likely fail.')
  elif was_enforcing:
    # Warn the user that we turned off enforcement.
    logging.warning('SELinux security has been disabled on the target device. '
                    'Enforcement will not be automatically enabled until the '
                    'next reboot.  To enable security before then, run the '
                    'command "' +
                    ' '.join(make_adb_command(['setenforce', '1'])) + '".')


def validate_options(options):
  """Checks for invalid command-line options.

  Args:
    options: The options structure to check.

  Returns:
    True if the options are valid, and False otherwise.
  """
  if not options.app_name:
    logging.error('APP_NAME must not be empty.')
    return False
  if not options.activity_name:
    logging.error('ACTIVITY_NAME must not be empty.')
    return False
  return True


def parse_options():
  """Parses command-line options.

  Returns:
    An options object initialized with settings for this invocation.
  """
  usage = 'Usage: %prog --app_name=<app> --activity-name=<activity> [options]'
  desc = ('Example: %prog --app_name=com.android.example '
          '--activity-name=.MainActivity --gfxspy-path=' +
          DEFAULT_GFXSPY_HOST_PATH)
  parser = optparse.OptionParser(usage=usage, description=desc)
  parser.add_option('--app-name', dest='app_name', default=None, type='string',
                    action='store',
                    help='The name of the application to trace '
                    '(e.g. com.android.example).')
  parser.add_option('--activity-name', dest='activity_name', default=None,
                    type='string', action='store',
                    help='The name of the activity to start '
                    '(e.g. .MainActivity).')
  parser.add_option('--trace-file', dest='trace_file_name',
                    default=None, type='string', action='store',
                    help='The name output trace file. Default is '
                    '<app_name>.gltrace.')
  parser.add_option('--gfxspy-path', dest='gfxspy_path',
                    default=DEFAULT_GFXSPY_HOST_PATH,
                    type='string', action='store',
                    help='Relative or absolute path to the gfxspy library to '
                    'install on the target device. Default is '
                    '"' + DEFAULT_GFXSPY_HOST_PATH + '".')
  parser.add_option('--no-install', dest='no_install',
                    action='store_true', default=False,
                    help='Do not install ' + GFXSPY_FILE_NAME + '. ' +
                    GFXSPY_FILE_NAME + ' must already be installed on the '
                    'target device in /data/.')
  parser.set_defaults(logging_level=logging.INFO)
  parser.add_option('-v', '--verbose', dest='logging_level',
                    action='store_const', const=logging.DEBUG,
                    help='Print extra debugging information.')
  parser.add_option('-q', '--quiet', dest='logging_level',
                    action='store_const', const=logging.WARNING,
                    help='Suppresses most console output.')
  parser.add_option('-s', dest='device_serial',
                    default='', type='string', action='store',
                    help='Directs adb commands to the device or emulator with '
                    'the given serial number or qualifier. Overrides '
                    'ANDROID_SERIAL environment variable.')

  options, unused_args = parser.parse_args()
  if not validate_options(options):
    parser.print_help()
    sys.exit(1)
  return options


def main():
  """Application entry point--initializes and runs the capture.
  """
  # Start the logging system.
  logging.basicConfig(format='%(levelname)s: %(message)s',
                      level=logging.INFO)

  options = parse_options()

  # Initialize our logging level stack.
  push_logging_level(options.logging_level)

  set_device_serial(options.device_serial)

  set_permissions()
  do_install(options)
  trace_file_name = build_trace_file_name(options.trace_file_name,
                                          options.app_name)
  trace_file = create_trace_file(trace_file_name)
  try:
    start_app(options.app_name, options.activity_name)
    collect_trace(trace_file)
  except KeyboardInterrupt:
    log_to_stdout(logging.INFO, '\nCollection terminated.\n')
  finally:
    logging.debug('Cleaning up...')
    trace_file.close()
    run_adb(['shell', 'setprop', build_wrap_prop_string(options.app_name), ''])

  try:
    if os.path.getsize(trace_file_name) == 0:
      logging.warning('No trace data was collected.  Removing "%s".',
                      trace_file_name)
      os.remove(trace_file_name)
  except OSError:
    logging.warning('Unable to access trace file "%s".', trace_file_name)

  logging.debug('Done.')


if __name__ == '__main__':
  main()
