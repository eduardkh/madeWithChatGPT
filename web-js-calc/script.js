document.addEventListener('DOMContentLoaded', () => {
    let currentInput = '';
    let previousInput = '';
    let operation = null;
    const display = document.querySelector('.text-3xl');

    document.querySelectorAll('.calc-button').forEach(button => {
        button.addEventListener('click', () => {
            const value = button.textContent;

            if (value >= '0' && value <= '9') {
                currentInput += value;
                updateDisplay();
            } else if (value === 'CE') {
                clear();
            } else if (value === '=') {
                if (operation && currentInput) {
                    calculate();
                    updateDisplay();
                    operation = null;
                }
            } else {
                if (currentInput) {
                    if (!operation) {
                        previousInput = currentInput;
                        currentInput = '';
                    }
                    operation = value;
                }
            }
        });
    });

    function calculate() {
        let result;
        const prev = parseFloat(previousInput);
        const current = parseFloat(currentInput);

        switch (operation) {
            case '+':
                result = prev + current;
                break;
            case '-':
                result = prev - current;
                break;
            case '*':
                result = prev * current;
                break;
            case '/':
                if (current === 0) {
                    alert("Can't divide by 0.");
                    return;
                } else {
                    result = prev / current;
                }
                break;
            default:
                return;
        }
        currentInput = result.toString();
        operation = undefined;
        previousInput = '';
    }

    function updateDisplay() {
        // If currentInput is empty, display '0'. Otherwise, display currentInput.
        display.textContent = currentInput === '' ? '0' : currentInput;
    }

    function clear() {
        // Set currentInput to '0' instead of empty string to ensure display shows '0'
        currentInput = '';
        previousInput = '';
        operation = null;
        updateDisplay();
    }
});
