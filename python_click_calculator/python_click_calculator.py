import click


@click.group()
def cli():
    """Simple Calculator CLI built with Click."""
    pass


@cli.command(help="Adds two numbers. Usage: python calculator.py add [X] [Y]")
@click.argument('x', type=float)
@click.argument('y', type=float)
def add(x, y):
    """Addition of two numbers."""
    click.echo(f"{x} + {y} = {x + y}")


@cli.command(help="Subtracts two numbers. Usage: python calculator.py subtract [X] [Y]")
@click.argument('x', type=float)
@click.argument('y', type=float)
def subtract(x, y):
    """Subtraction of two numbers."""
    click.echo(f"{x} - {y} = {x - y}")


@cli.command(help="Multiplies two numbers. Usage: python calculator.py multiply [X] [Y]")
@click.argument('x', type=float)
@click.argument('y', type=float)
def multiply(x, y):
    """Multiplication of two numbers."""
    click.echo(f"{x} * {y} = {x * y}")


@cli.command(help="Divides two numbers. Usage: python calculator.py divide [X] [Y]")
@click.argument('x', type=float)
@click.argument('y', type=float)
def divide(x, y):
    """Division of two numbers."""
    try:
        result = x / y
        click.echo(f"{x} / {y} = {result}")
    except ZeroDivisionError:
        click.echo("Error: Division by zero is not allowed.")


if __name__ == '__main__':
    cli()
