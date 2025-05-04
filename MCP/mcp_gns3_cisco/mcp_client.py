import requests
import ollama
import json
import re

# === Config ===
MCP_SERVER_URL = "http://192.168.1.250:5000/execute"
MODEL_NAME = "llama3:8b"


# === LLM Query ===
def query_model(messages):
    raw = ollama.chat(model=MODEL_NAME, messages=messages)[
        'message']['content']

    # Extract JSON block if present
    json_match = re.search(r'\{.*\}', raw, re.DOTALL)
    if json_match:
        try:
            return json.loads(json_match.group())
        except json.JSONDecodeError:
            pass

    # Try fallback cleanup
    if "```json" in raw:
        raw = raw.split("```json")[1].split("```", 1)[0].strip()
    elif "```" in raw:
        raw = raw.split("```", 1)[1].split("```", 1)[0].strip()

    try:
        return json.loads(raw)
    except json.JSONDecodeError:
        return {
            "action": "respond",
            "message": f"⚠️ Invalid JSON from model. Raw output:\n{raw}"
        }


# === MCP Runner ===
def run_command_on_switch(payload: dict):
    try:
        response = requests.post(MCP_SERVER_URL, json=payload)
        return response.json()
    except Exception as e:
        return {"stdout": "", "stderr": f"❌ Failed to run command: {str(e)}"}


# === Main Loop ===
def main():
    print("\U0001F4AC Talk to the Cisco agent (via llama3:8b). Type 'exit' to quit.\n")

    messages = [
        {"role": "system", "content": (
            "You are a Cisco CLI assistant. Always respond in JSON only. Never use markdown or code blocks.\n"
            "Respond using one of:\n"
            "{ \"action\": \"run_command\", \"command\": \"show version\" }\n"
            "{ \"action\": \"run_commands\", \"commands\": [\"conf t\", \"vlan 20\", \"name foo\"] }\n"
            "{ \"action\": \"respond\", \"message\": \"<reply>\" }"
        )}
    ]

    last_command = None

    while True:
        user_input = input("\U0001F9D1‍\U0001F4BB: ").strip()
        if user_input.lower() in ("exit", "quit"):
            break

        messages.append({"role": "user", "content": user_input})
        result = query_model(messages)

        if result["action"] == "run_command":
            command = result.get("command", "").strip()
            print(f"\n▶ Executing: {command}")
            mcp_output = run_command_on_switch({"command": command})
            stdout = mcp_output.get("stdout", "").strip()
            stderr = mcp_output.get("stderr", "").strip()
            if stdout:
                print(f"\n\U0001F4E1 Output:\n{stdout}")
            if stderr:
                print(f"\n⚠️ Error:\n{stderr}")

            # Evaluate response
            last_command = command
            if any(x in user_input.lower() for x in ["ip address", "vlans", "interfaces", "list"]):
                messages.append(
                    {"role": "user", "content": f"Please analyze the following output of `{command}` and answer the original request:```\n{stdout}\n```"})
                result = query_model(messages)
                if result["action"] == "respond":
                    print(f"\n\U0001F916: {result['message']}\n")
                    messages.append(
                        {"role": "assistant", "content": result["message"]})

        elif result["action"] == "run_commands":
            commands = result.get("commands", [])
            print(f"\n▶ Executing interactive command sequence:")
            for c in commands:
                print(f" → {c}")
            mcp_output = run_command_on_switch({"commands": commands})
            stdout = mcp_output.get("stdout", "").strip()
            stderr = mcp_output.get("stderr", "").strip()
            if stdout:
                print(f"\n\U0001F4E1 Output:\n{stdout}")
            if stderr:
                print(f"\n⚠️ Error:\n{stderr}")

            messages.append(
                {"role": "user", "content": f"Command result:\n{stdout or stderr}"})

        elif result["action"] == "respond":
            print(f"\n\U0001F916: {result['message']}\n")
            messages.append(
                {"role": "assistant", "content": result["message"]})

        else:
            print(f"\n⚠️ Unknown action or invalid format:\n{result}")


if __name__ == "__main__":
    main()
