import paramiko


def run_command(host, username, password, command):
    client = paramiko.SSHClient()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    try:
        client.connect(
            hostname=host,
            username=username,
            password=password,
            allow_agent=False,
            look_for_keys=False,
            disabled_algorithms={
                'kex': ['ecdh-sha2-nistp256', 'curve25519-sha256'],
                'mac': ['hmac-sha2-256', 'hmac-sha2-512']
            }
        )

        stdin, stdout, stderr = client.exec_command(command)
        output = stdout.read().decode()
        print(f"[Output]\n{output}")

    except Exception as e:
        print(f"[ERROR] {e}")
    finally:
        client.close()


if __name__ == "__main__":
    run_command(
        host="172.28.112.98",
        username="cisco",
        password="cisco",
        command="show ip interface brief"
    )
