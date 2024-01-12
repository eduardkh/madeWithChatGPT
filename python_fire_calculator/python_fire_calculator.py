import fire


class Calculator:
    """Simple calculator class with basic operations."""

    def add(self, x, y):
        """Adds two numbers."""
        return x + y

    def subtract(self, x, y):
        """Subtracts second number from first."""
        return x - y

    def multiply(self, x, y):
        """Multiplies two numbers."""
        return x * y

    def divide(self, x, y):
        """Divides first number by second. Raises error if division by zero."""
        if y == 0:
            raise ValueError("Cannot divide by zero.")
        return x / y


if __name__ == '__main__':
    fire.Fire(Calculator)
