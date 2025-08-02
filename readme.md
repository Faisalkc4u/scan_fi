✅ Step 1: Build your Go binary locally
In your project root, run:

bash
Copy
Edit
GOOS=linux GOARCH=amd64 go build -o main .
This compiles the Go binary for Linux (needed since the Docker image is based on Alpine Linux).

✅ Step 2: Create the Dockerfile
Make sure your Dockerfile looks like this (already done):

dockerfile
Copy
Edit
FROM alpine:latest
WORKDIR /app
COPY main .
EXPOSE 8080
CMD ["./main"]
✅ Step 3: Build the Docker image
In the same directory as the Dockerfile, run:

bash
Copy
Edit
docker build -t my-go-app .
✅ Step 4: Run the Docker container
bash
Copy
Edit
docker run -p 8080:8080 my-go-app
If your Go app uses a different port (e.g. 8000 from os.Getenv("PORT")), run:

bash
docker run -e PORT=8000 -p 8000:8000 scan-fi
🔁 Optional: Clean up old builds
To remove the local binary after building the image:

bash
Copy
Edit
rm main
