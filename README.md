# helloGoWeather
"hello world" with weather module - cmd line, uses http connection

## Instructions:
1. Register at `https://home.openweathermap.org/` to get your OpenWeather API Key. In the file `server.go`, replace `<YOUR_API_KEY>` with your actual API Key. Or contact me to get one.


2. In one terminal window run these commands:
```
git clone https://github.com/ellieasager/helloGoWeather
cd helloGoWeather
go mod init github.com/ellieasager/helloGoWeather
go install github.com/ellieasager/helloGoWeather
go build
go run server/server.go
```

3. In another terminal window type this:

`curl http://localhost:9000/weather/tokyo`
