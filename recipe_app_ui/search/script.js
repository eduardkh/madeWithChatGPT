// Function to toggle between single search and ingredient search
function toggleSearchType() {
    const singleSearchDiv = document.getElementById('singleSearch');
    const ingredientSearchDiv = document.getElementById('ingredientSearch');
    if (singleSearchDiv.classList.contains('hidden')) {
        singleSearchDiv.classList.remove('hidden');
        ingredientSearchDiv.classList.add('hidden');
    } else {
        singleSearchDiv.classList.add('hidden');
        ingredientSearchDiv.classList.remove('hidden');
    }
}

// Function to add new ingredient input fields
function addIngredientField() {
    const ingredientFieldsDiv = document.getElementById('ingredientFields');
    const newFieldDiv = document.createElement('div');
    newFieldDiv.classList.add('flex', 'justify-center', 'mb-2', 'w-full');
    newFieldDiv.innerHTML = `
        <input type="text" placeholder="Ingredient" class="w-1/2 p-2 text-lg leading-relaxed tracking-wide rounded-l border">
        <input type="text" placeholder="Quantity" class="w-1/12 p-2 text-lg leading-relaxed tracking-wide rounded-r border border-l-0">
    `;
    ingredientFieldsDiv.appendChild(newFieldDiv);
}

// Ensure the first ingredient fields are present
addIngredientField();
