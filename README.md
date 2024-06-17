# examsPrepGo
How to set in local
1. Make sure you have go lang version 1.21 and above
2. Configure API_KEY, refer [Python GenAI Reference](https://github.com/gitish/GenAI-Gemini)
3. Configure environment variable for API_KEY
```export GOOGLE_API_KEY=<your-api-key>```
4. Run below code for install dependencies and run server
```
    go version
    go mod init shl/genai
    go get
    go mod tidy
    go run .
```
