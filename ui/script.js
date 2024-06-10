const baseUrl = 'http://localhost:8080';

document.getElementById('subscriptionForm').addEventListener('submit', function (e) {
    e.preventDefault();

    const formData = new FormData(this);
    const data = Object.fromEntries(formData.entries());

    data.subscriberId = parseInt(data.subscriberId, 10);

    fetch(baseUrl.concat('/subscription'), {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }).then(response => response.json())
        .then(result => {

            alert('Subscription added successfully!');
            loadData();
            showTab('data');

        });
});

function uploadJsonFile() {
    const fileInput = document.getElementById('jsonFileInput');
    const file = fileInput.files[0];

    if (file) {
        const reader = new FileReader();
        reader.onload = function (event) {
            const data = JSON.parse(event.target.result);

            if (Array.isArray(data)) {
                data.forEach(sub => {
                    sub.subscriberId = parseInt(sub.subscriberId, 10);
                });
            } else {
                data.subscriberId = parseInt(data.subscriberId, 10);
            }

            fetch(baseUrl.concat('/subscription'), {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            }).then(response => response.json())
                .then(result => {

                    alert('Subscriptions added successfully!');
                    loadData();
                    showTab('data');

                });
        };
        reader.readAsText(file);
    } else {
        alert('Please select a JSON file to upload.');
    }
}

function generateTableHead() {
    const table = document.getElementById('subscriptionTable');
    const thead = document.createElement('thead');
    thead.innerHTML = `
      <tr>
        <th style="background-color: #f2f2f2; padding: 8px; text-align: left; border: 1px solid #ddd;">Subscriber ID</th>
        <th style="background-color: #f2f2f2; padding: 8px; text-align: left; border: 1px solid #ddd;">Subscription ID</th>
        <th style="background-color: #f2f2f2; padding: 8px; text-align: left; border: 1px solid #ddd;">Subscriber Name</th>
        <th style="background-color: #f2f2f2; padding: 8px; text-align: left; border: 1px solid #ddd;">Subscriber Country</th>
        <th style="background-color: #f2f2f2; padding: 8px; text-align: left; border: 1px solid #ddd;">Subscription Date</th>
      </tr>
    `;
    table.insertBefore(thead, table.childNodes[0]); // Insert thead before tbody
}
function loadData() {
    fetch(baseUrl.concat('/subscriptions'))
        .then(response => response.json())
        .then(data => {
            const tableBody = document.getElementById('subscriptionTableBody');
            tableBody.innerHTML = '';

            data.forEach(sub => {
                const row = document.createElement('tr');
                row.style.backgroundColor = '#f5f5f5'; // Apply hover effect inline
                row.innerHTML = `
                <td style="padding: 8px; text-align: left; border: 1px solid #ddd;">${sub.subscriberId}</td>
                <td style="padding: 8px; text-align: left; border: 1px solid #ddd;">${sub.subscriptionId}</td>
                <td style="padding: 8px; text-align: left; border: 1px solid #ddd;">${sub.subscriberName}</td>
                <td style="padding: 8px; text-align: left; border: 1px solid #ddd;">${sub.subscriberCountry}</td>
                <td style="padding: 8px; text-align: left; border: 1px solid #ddd;">${new Date(sub.subscriptionDate).toLocaleString()}</td>
              `;
                tableBody.appendChild(row);
            });
        })
        .catch(error => {
            console.error('Error fetching data:', error);
        });


    // Call the function to load data when the page is loaded
    document.addEventListener("DOMContentLoaded", function () {
        loadData()
        generateTableHead();
    });


    fetch(baseUrl.concat('/subscriptions/total'))
        .then(response => response.json())
        .then(data => {
            document.getElementById('totalSubscriberCount').textContent = data.Total_Subscriber_Count;
        });

    fetch(baseUrl.concat('/subscriptions/longest'))
        .then(response => response.json())
        .then(data => {
            document.getElementById('longestDurationSubscriber').textContent = `${data.Subscriber_Longest_Duration} `;
        });

    fetch(baseUrl.concat('/subscriptions/most-subscribers'))
        .then(response => response.json())
        .then(data => {
            document.getElementById('countryWithMostSubscribers').textContent = data.Country_with_Most_Subscribers;
        });
}

function showTab(tab) {
    document.querySelectorAll('.tab-content').forEach(function (tabContent) {
        tabContent.style.display = 'none';
    });
    document.getElementById(tab).style.display = 'block';
}

// Load initial data
loadData();