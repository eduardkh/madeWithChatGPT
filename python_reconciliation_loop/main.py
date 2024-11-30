import time

DESIRED_STATE_FILE = "desired_state.txt"
CURRENT_STATE_FILE = "current_state.txt"


def get_desired_state():
    """Read the desired state from the file."""
    try:
        with open(DESIRED_STATE_FILE, 'r') as file:
            return [line.strip() for line in file.readlines()]
    except FileNotFoundError:
        print(f"File {DESIRED_STATE_FILE} not found. Creating an empty file.")
        with open(DESIRED_STATE_FILE, 'w') as file:
            pass  # Create an empty file
        return []


def get_current_state():
    """Read the current state from the file."""
    try:
        with open(CURRENT_STATE_FILE, 'r') as file:
            return [line.strip() for line in file.readlines()]
    except FileNotFoundError:
        print(f"File {CURRENT_STATE_FILE} not found. Creating an empty file.")
        with open(CURRENT_STATE_FILE, 'w') as file:
            pass  # Create an empty file
        return []


def reconcile_state(desired, current):
    """Reconcile the current state with the desired state."""
    if desired != current:
        with open(CURRENT_STATE_FILE, 'w') as file:
            for item in desired:
                file.write(item + '\n')
        print("Current state updated!")
    else:
        print("Current state is already in sync.")


def main():
    try:
        while True:
            print("\n=== Reconciliation Loop Started ===")
            desired_state = get_desired_state()
            current_state = get_current_state()

            print(f"Desired state: {desired_state}")
            print(f"Current state before reconciliation: {current_state}")

            reconcile_state(desired_state, current_state)

            print("=== Reconciliation Loop Finished ===\n")

            # Wait before running the loop again
            time.sleep(5)
    except KeyboardInterrupt:
        print("\nReconciliation process interrupted. Exiting gracefully.")
        exit(0)


if __name__ == "__main__":
    main()
