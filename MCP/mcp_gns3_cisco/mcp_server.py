from flask import Flask, request, jsonify
import paramiko
import time

app = Flask(__name__)

# Cisco device credentials
CISCO_HOST = "172.28.112.98"
CISCO_USER = "cisco"
CISCO_PASS = "cisco"


def run_interactive_commands(commands):
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh.connect(
        hostname=CISCO_HOST,
        username=CISCO_USER,
        password=CISCO_PASS,
        allow_agent=False,
        look_for_keys=False,
        disabled_algorithms={
            'kex': ['ecdh-sha2-nistp256', 'curve25519-sha256'],
            'mac': ['hmac-sha2-256', 'hmac-sha2-512']
        }
    )

    shell = ssh.invoke_shell()
    time.sleep(1)
    shell.recv(1000)  # Clear banner

    output = ""
    for cmd in commands:
        shell.send(cmd + "\n")
        time.sleep(1.2)
        output += shell.recv(5000).decode(errors='ignore')

    ssh.close()
    return output


@app.route('/execute', methods=['POST'])
def execute():
    data = request.get_json()

    single_cmd = data.get("command")
    command_list = data.get("commands")

    if single_cmd:
        return exec_single_command(single_cmd)

    elif isinstance(command_list, list):
        try:
            out = run_interactive_commands(command_list)
            return jsonify({"stdout": out, "stderr": ""})
        except Exception as e:
            return jsonify({"error": str(e)}), 500
    else:
        return jsonify({"error": "Missing 'command' or 'commands'"}), 400


def exec_single_command(command):
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    try:
        ssh.connect(
            hostname=CISCO_HOST,
            username=CISCO_USER,
            password=CISCO_PASS,
            allow_agent=False,
            look_for_keys=False,
            disabled_algorithms={
                'kex': ['ecdh-sha2-nistp256', 'curve25519-sha256'],
                'mac': ['hmac-sha2-256', 'hmac-sha2-512']
            }
        )

        stdin, stdout, stderr = ssh.exec_command(command)
        stdout_text = stdout.read().decode(errors='ignore')
        stderr_text = stderr.read().decode(errors='ignore')

        return jsonify({
            "stdout": stdout_text,
            "stderr": stderr_text
        })

    except Exception as e:
        return jsonify({"error": str(e)}), 500

    finally:
        ssh.close()


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
