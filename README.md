# Subscription_Management_System

```markdown

This is a basic application for a Publishing Company to collect, store, and display subscription details. It includes two pages: one for data entry (either through a form or JSON file upload) and another for displaying the stored data in a tabular format, along with insights like total subscriber count, subscriber with the longest duration, and the country with the most subscribers.

## Technologies Used

- Backend: Golang, Gin framework, GORM
- Database: PostgreSQL
- Frontend: HTML, CSS, JavaScript

## Setup Instructions

1. Clone the repository:

   ```bash
   git clone https://github.com/MrSanketKumar/Subscription_Management_System.git
   ```

2. Navigate to the project directory:

   ```bash
   cd Subscription_Management_System
   ```

3. Install the required dependencies (assuming you have Go and Node.js installed).

4. Set up the PostgreSQL database and configure the connection details in the `config/config.go` file.

5. Run the application:

   ```bash
   go run main.go
   ```

   The application should now be running at `http://localhost:8080`.

## Endpoints

### Page 1: Data Entry

#### Form Submission

- **URL:** `/submit-form`
- **Method:** `POST`
- **Description:** Accepts form data containing the subscription details (subscriberId, subscriberName, subscriberCountry, subscriberDuration). Validates the data, inserts it into the database, and shows errors (if any) before navigating to the second page.

#### JSON Upload

- **URL:** `/upload-json`
- **Method:** `POST`
- **Description:** Accepts a JSON file containing multiple subscription details. Validates the data, inserts it into the database, and returns a success or failure response.

### Page 2: Data Display and Insights

#### Get Subscription Data

- **URL:** `/subscriptions`
- **Method:** `GET`
- **Description:** Retrieves all subscription data from the database and returns it in tabular format.

#### Get Total Subscriber Count

- **URL:** `/insights/total-subscribers`
- **Method:** `GET`
- **Description:** Retrieves the total number of subscribers from the database.

#### Get Subscriber with Longest Duration

- **URL:** `/insights/longest-duration`
- **Method:** `GET`
- **Description:** Retrieves the subscriber with the longest subscription duration from the database.

#### Get Country with Most Subscribers

- **URL:** `/insights/top-country`
- **Method:** `GET`
- **Description:** Retrieves the country with the highest number of subscribers from the database.

## Backend API Testing

The following constants define the paths for the backend API endpoints:

```go
    CreateSubscriptionPath             = "/subscription"
    GetAllSubscriptionsPath            = "/subscriptions"
    GetTotalSubscribersPath            = "/subscriptions/total"
    GetLongestSubscriptionDurationPath = "/subscriptions/longest"
    GetMostSubscribersPath             = "/subscriptions/most-subscribers"
```

These paths can be used for testing the backend API endpoints using tools like Postman or cURL. Here's how you can test each endpoint:

### Create Subscription

- **URL:** `http://localhost:8080/subscription`
- **Method:** `POST`
- **Request Body:** JSON object representing the subscription details.
- **Example cURL Request:**

  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"subscriberId": "1234", "subscriberName": "John Doe", "subscriberCountry": "USA", "subscriberDuration": 365}' http://localhost:8080/subscription
  ```

### Get All Subscriptions

- **URL:** `http://localhost:8080/subscriptions`
- **Method:** `GET`
- **Example cURL Request:**

  ```bash
  curl http://localhost:8080/subscriptions
  ```

### Get Total Subscribers

- **URL:** `http://localhost:8080/subscriptions/total`
- **Method:** `GET`
- **Example cURL Request:**

  ```bash
  curl http://localhost:8080/subscriptions/total
  ```

### Get Subscriber with Longest Duration

- **URL:** `http://localhost:8080/subscriptions/longest`
- **Method:** `GET`
- **Example cURL Request:**

  ```bash
  curl http://localhost:8080/subscriptions/longest
  ```

### Get Country with Most Subscribers

- **URL:** `http://localhost:8080/subscriptions/most-subscribers`
- **Method:** `GET`
- **Example cURL Request:**

  ```bash
  curl http://localhost:8080/subscriptions/most-subscribers
  ```

These examples demonstrate how to test the backend API endpoints using cURL. You can also use tools like Postman or other HTTP clients to send requests to the respective endpoints and verify the responses.

## UI

The application includes two pages:

1. **Data Entry Page:** This page allows users to enter subscription details either through a form or by uploading a JSON file. The page is accessible at `http://localhost:8080/`.

2. **Data Display and Insights Page:** This page displays the subscription data in a tabular format and provides insights like total subscriber count, subscriber with the longest duration, and the country with the most subscribers. The page is accessible at `http://localhost:8080/subscriptions`.

## Screenshots

### Data Entry Page
![Data Entry Page](https://github.com/MrSanketkumar/Bito_Assignment_Sanket_Chaudhary/blob/97e7c7e0bbecfc8e514d7b2b1a8261648ddf7f70/Bito_Submit.PNG)

### Data Display and Insights Page
![Data Display and Insights Page](https://github.com/MrSanketkumar/Bito_Assignment_Sanket_Chaudhary/blob/97e7c7e0bbecfc8e514d7b2b1a8261648ddf7f70/Bito_ViewData.PNG)
