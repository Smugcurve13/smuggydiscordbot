

cmd = "!run echo hello"

def command_extracter(cmd):
    cmd = cmd[5:]
    return cmd

print(command_extracter(cmd))