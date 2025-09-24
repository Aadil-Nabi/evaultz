### `evaultz`

THIS PROJECT IS STILL IN THE INITIAL PHASE, STAY TUNED

### `Description`

# Welcome to evaultz app 👋

This is an file storage app, we can store any file like docs, videos, pdf etc.

### `Installation` ✔

Make sure you have latest GoLang installed on your workstation, and you can download the code and use it directly on any IDE

### `Usage` ▶

- Create a env file, example as below. Make sure you don't store this config file in a public repository or network share for security reasons
  You can store the passwords or secrets in an external secret management vault like AKeyless Vault or Hashicorp Vault

```bash
ENV=dev
PORT=8082
DB_URL="host=192.168.238.147 user=postgres password=Asdf@1234 dbname=evaultz port=5432 sslmode=disable"
AWS_REGION=ap-south-1
AWS_ACCESS_KEY_ID=XXXXXXXXXXXXXXXXX
AWS_SECRET_ACCESS_KEY=XXXXXXXXXXXXXXXXXXXXXX+XXXXXXXXXXXXXXXX/w
TOKEN_SECRET=XXXXXXXXXXXXXXXXXXXXXXXX
```

To run the program you need to execute the run command in the below format

```bash
go run cmd/evaultz/main.go
```

Build the program into an executable file for your specific platform, Example below for Mac/Linux

```bash
go build -o evaultz cmd/evaultz/main.go
```

### `Contributing` 🤝

Contributions, issues and feature requests are welcome, but it is paused for now.
Feel free to check issues page if you want to contribute in future
