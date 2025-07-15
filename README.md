# Worldline – Find Perfect Numbers

Challenge for the position of **Senior Go Developer** at **Worldline**.  
The task consisted of developing a REST API that identifies **perfect numbers** within a given range.

---

## 📌 What is a Perfect Number?

A **perfect number** is a positive integer that is equal to the sum of its proper divisors (excluding itself).  
Example: 6 is perfect because 1 + 2 + 3 = 6.  
For more details, see [Wikipedia – Perfect number](https://en.wikipedia.org/wiki/Perfect_number).

---

## 🛠️ How the API Works

The API receives a JSON payload with two integers: `start` and `end`.  
It returns a list of all perfect numbers within that inclusive range.

---

## ✅ Assumptions

- The `start` and `end` values must be positive integers.
- If `start > end`, the API returns an empty list.
- The algorithm handles large ranges efficiently using concurrent processing (via goroutines).
- Input validation is in place to ensure the payload is correct.

---

## 🧪 Example Request

```
POST /perfect-numbers
Content-Type: application/json

{
  "start": 1,
  "end": 10000
}
```

---

## 🔁 Example Response

```json
{
  "perfect_numbers": [6, 28, 496, 8128]
}
```

---

## 🚀 Running the API

In the terminal, navigate to the project folder and run:

```
go run main.go
```

Once running, test it with:

```
curl -X POST http://localhost:8080/perfect-numbers \
     -H "Content-Type: application/json" \
     -d '{"start": 1, "end": 10000}' | jq
```

---

## 🧪 Tests

Unit tests were written for the core logic and the HTTP handler.  
Run them using:

```
go test -v ./...
```

---

## 📚 Tech Stack

- Language: Go
- Framework: Echo (for HTTP routing)
- Standard library for JSON handling and concurrency

Code and comments are in English, as per the developer's usual practice.