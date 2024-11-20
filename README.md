# Receipt Processor Project

## Intro
This is my challenge solution for 
[Fetch Receipt Processor](https://github.com/fetch-rewards/receipt-processor-challenge).


### Prerequisites
* [Go](https://go.dev/) if want to run locally
* [Docker](https://docs.docker.com/get-started/get-docker/)

### Run on Local Machine on Docker
1. Clone the repo and change directory
   ```sh
   git clone https://github.com/RichtXO/receipt-processor.git
   cd receipt-processor
   ```
2. Run the Receipt Processor Server!
   ```sh
   docker compose up --build -d
   ```

## Libraries Used
* [mux](https://github.com/gorilla/mux)
* [uuid](https://github.com/google/uuid)