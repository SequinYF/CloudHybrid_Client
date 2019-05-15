#!/usr/bin/expect

set UPLOAD_FILE  [lindex $argv 0]
set USER_NAME  [lindex $argv 1]
set USER_PASSWD [lindex $argv 2]

puts $USER_PASSWD

spawn rsync -a --progress $UPLOAD_FILE $USER_NAME@118.89.234.46:~/
expect "*password:"
send "$USER_PASSWD\n"
#interact
expect eof
exit
