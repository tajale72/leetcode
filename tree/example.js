const apiUrl = 'http://localhost:8080/btc';

// Fetch the API data and update the table with the response
fetch(apiUrl)
    .then(response => response.json())
    .then(data => {
        // Select the table element
        const table = document.getElementById('api-table');

        // Create a new row for the API data
        const row = table.insertRow();

        // Add each data field as a table cell
        const dateCell = row.insertCell();
        dateCell.innerHTML = data.date;

        const currencyCell = row.insertCell();
        currencyCell.innerHTML = data.currency;

        const nameCell = row.insertCell();
        nameCell.innerHTML = data.name;

        const priceCell = row.insertCell();
        priceCell.innerHTML = data.price;

        const lastUpdatedCell = row.insertCell();
        lastUpdatedCell.innerHTML = data.lastUpdated;
    })
    .catch(error => console.error(error));
